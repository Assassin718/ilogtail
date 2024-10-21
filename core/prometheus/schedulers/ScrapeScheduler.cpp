/*
 * Copyright 2024 iLogtail Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#include "prometheus/schedulers/ScrapeScheduler.h"

#include <cstddef>
#include <memory>
#include <string>
#include <utility>

#include "common/StringTools.h"
#include "common/TimeUtil.h"
#include "common/timer/HttpRequestTimerEvent.h"
#include "common/timer/Timer.h"
#include "logger/Logger.h"
#include "pipeline/queue/ProcessQueueItem.h"
#include "pipeline/queue/ProcessQueueManager.h"
#include "pipeline/queue/QueueKey.h"
#include "prometheus/Constants.h"
#include "prometheus/async/PromFuture.h"
#include "prometheus/async/PromHttpRequest.h"
#include "sdk/Common.h"

using namespace std;

namespace logtail {

ScrapeScheduler::ScrapeScheduler(std::shared_ptr<ScrapeConfig> scrapeConfigPtr,
                                 std::string host,
                                 int32_t port,
                                 Labels labels,
                                 QueueKey queueKey,
                                 size_t inputIndex)
    : mScrapeConfigPtr(std::move(scrapeConfigPtr)),
      mHost(std::move(host)),
      mPort(port),
      mTargetLabels(std::move(labels)),
      mQueueKey(queueKey),
      mInputIndex(inputIndex) {
    string tmpTargetURL = mScrapeConfigPtr->mScheme + "://" + mHost + ":" + ToString(mPort)
        + mScrapeConfigPtr->mMetricsPath
        + (mScrapeConfigPtr->mQueryString.empty() ? "" : "?" + mScrapeConfigPtr->mQueryString);
    mHash = mScrapeConfigPtr->mJobName + tmpTargetURL + ToString(mTargetLabels.Hash());
    mInstance = mHost + ":" + ToString(mPort);
    mInterval = mScrapeConfigPtr->mScrapeIntervalSeconds;

    mParser = make_unique<TextParser>();
}

void ScrapeScheduler::OnMetricResult(const HttpResponse& response, uint64_t timestampMilliSec) {
    mSelfMonitor->AddCounter(METRIC_PLUGIN_OUT_EVENTS_TOTAL, response.mStatusCode);
    mSelfMonitor->AddCounter(METRIC_PLUGIN_OUT_SIZE_BYTES, response.mStatusCode, response.mBody.size());
    mSelfMonitor->AddCounter(
        METRIC_PLUGIN_PROM_SCRAPE_TIME_MS, response.mStatusCode, GetCurrentTimeInMilliSeconds() - timestampMilliSec);

    mScrapeTimestampMilliSec = timestampMilliSec;
    mScrapeDurationSeconds = 1.0 * (GetCurrentTimeInMilliSeconds() - timestampMilliSec) / 1000;
    mScrapeResponseSizeBytes = response.mBody.size();
    mUpState = response.mStatusCode == 200;
    if (response.mStatusCode != 200) {
        mScrapeResponseSizeBytes = 0;
        string headerStr;
        for (const auto& [k, v] : mScrapeConfigPtr->mRequestHeaders) {
            headerStr.append(k).append(":").append(v).append(";");
        }
        LOG_WARNING(sLogger,
                    ("scrape failed, status code", response.mStatusCode)("target", mHash)("http header", headerStr));
    }
    auto eventGroup = BuildPipelineEventGroup(response.mBody);

    SetAutoMetricMeta(eventGroup);
    SetTargetLabels(eventGroup);
    PushEventGroup(std::move(eventGroup));
    mPluginTotalDelayMs->Add(GetCurrentTimeInMilliSeconds() - timestampMilliSec);
}

void ScrapeScheduler::SetAutoMetricMeta(PipelineEventGroup& eGroup) {
    eGroup.SetMetadata(EventGroupMetaKey::PROMETHEUS_SCRAPE_TIMESTAMP_MILLISEC, ToString(mScrapeTimestampMilliSec));
    eGroup.SetMetadata(EventGroupMetaKey::PROMETHEUS_SCRAPE_DURATION, ToString(mScrapeDurationSeconds));
    eGroup.SetMetadata(EventGroupMetaKey::PROMETHEUS_SCRAPE_RESPONSE_SIZE, ToString(mScrapeResponseSizeBytes));
    eGroup.SetMetadata(EventGroupMetaKey::PROMETHEUS_UP_STATE, ToString(mUpState));
}

void ScrapeScheduler::SetTargetLabels(PipelineEventGroup& eGroup) {
    mTargetLabels.Range([&eGroup](const std::string& key, const std::string& value) { eGroup.SetTag(key, value); });
}

PipelineEventGroup ScrapeScheduler::BuildPipelineEventGroup(const std::string& content) {
    return mParser->BuildLogGroup(content);
}

void ScrapeScheduler::PushEventGroup(PipelineEventGroup&& eGroup) {
    auto item = make_unique<ProcessQueueItem>(std::move(eGroup), mInputIndex);
#ifdef APSARA_UNIT_TEST_MAIN
    mItem.push_back(std::move(item));
    return;
#endif
    while (true) {
        if (ProcessQueueManager::GetInstance()->PushQueue(mQueueKey, std::move(item)) == 0) {
            break;
        }
        usleep(10 * 1000);
    }
}

string ScrapeScheduler::GetId() const {
    return mHash;
}

void ScrapeScheduler::ScheduleNext() {
    auto future = std::make_shared<PromFuture<const HttpResponse&, uint64_t>>();
    auto isContextValidFuture = std::make_shared<PromFuture<>>();
    future->AddDoneCallback([this](const HttpResponse& response, uint64_t timestampMilliSec) {
        this->OnMetricResult(response, timestampMilliSec);
        this->ExecDone();
        this->ScheduleNext();
        return true;
    });
    isContextValidFuture->AddDoneCallback([this]() -> bool {
        if (ProcessQueueManager::GetInstance()->IsValidToPush(mQueueKey)) {
            return true;
        } else {
            this->DelayExecTime(1);
            this->mPromDelayTotal->Add(1);
            this->ScheduleNext();
            return false;
        }
    });

    if (IsCancelled()) {
        mFuture->Cancel();
        mIsContextValidFuture->Cancel();
        return;
    }

    {
        WriteLock lock(mLock);
        mFuture = future;
        mIsContextValidFuture = isContextValidFuture;
    }

    auto event = BuildScrapeTimerEvent(GetNextExecTime());
    mTimer->PushEvent(std::move(event));
}

void ScrapeScheduler::ScrapeOnce(std::chrono::steady_clock::time_point execTime) {
    auto future = std::make_shared<PromFuture<const HttpResponse&, uint64_t>>();
    future->AddDoneCallback([this](const HttpResponse& response, uint64_t timestampMilliSec) {
        this->OnMetricResult(response, timestampMilliSec);
        return true;
    });
    mFuture = future;
    auto event = BuildScrapeTimerEvent(execTime);
    if (mTimer) {
        mTimer->PushEvent(std::move(event));
    }
}

std::unique_ptr<TimerEvent> ScrapeScheduler::BuildScrapeTimerEvent(std::chrono::steady_clock::time_point execTime) {
    auto request = std::make_unique<PromHttpRequest>(sdk::HTTP_GET,
                                                     mScrapeConfigPtr->mScheme == prometheus::HTTPS,
                                                     mHost,
                                                     mPort,
                                                     mScrapeConfigPtr->mMetricsPath,
                                                     mScrapeConfigPtr->mQueryString,
                                                     mScrapeConfigPtr->mRequestHeaders,
                                                     "",
                                                     mScrapeConfigPtr->mScrapeTimeoutSeconds,
                                                     mScrapeConfigPtr->mScrapeIntervalSeconds
                                                         / mScrapeConfigPtr->mScrapeTimeoutSeconds,
                                                     this->mFuture,
                                                     this->mIsContextValidFuture);
    auto timerEvent = std::make_unique<HttpRequestTimerEvent>(execTime, std::move(request));
    return timerEvent;
}

void ScrapeScheduler::Cancel() {
    if (mFuture != nullptr) {
        mFuture->Cancel();
    }
    if (mIsContextValidFuture != nullptr) {
        mIsContextValidFuture->Cancel();
    }
    {
        WriteLock lock(mLock);
        mValidState = false;
    }
}

void ScrapeScheduler::SetTimer(std::shared_ptr<Timer> timer) {
    mTimer = std::move(timer);
}

void ScrapeScheduler::InitSelfMonitor(const MetricLabels& defaultLabels) {
    mSelfMonitor = std::make_shared<PromSelfMonitorUnsafe>();
    MetricLabels labels = defaultLabels;
    labels.emplace_back(METRIC_LABEL_KEY_INSTANCE, mInstance);

    static const std::unordered_map<std::string, MetricType> sScrapeMetricKeys = {
        {METRIC_PLUGIN_OUT_EVENTS_TOTAL, MetricType::METRIC_TYPE_COUNTER},
        {METRIC_PLUGIN_OUT_SIZE_BYTES, MetricType::METRIC_TYPE_COUNTER},
        {METRIC_PLUGIN_PROM_SCRAPE_TIME_MS, MetricType::METRIC_TYPE_COUNTER},
    };

    mSelfMonitor->InitMetricManager(sScrapeMetricKeys, labels);

    WriteMetrics::GetInstance()->PrepareMetricsRecordRef(mMetricsRecordRef, std::move(labels));
    mPromDelayTotal = mMetricsRecordRef.CreateCounter(METRIC_PLUGIN_PROM_SCRAPE_DELAY_TOTAL);
    mPluginTotalDelayMs = mMetricsRecordRef.CreateCounter(METRIC_PLUGIN_TOTAL_DELAY_MS);
}

} // namespace logtail