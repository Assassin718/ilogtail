/*
 * Copyright 2022 iLogtail Authors
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

#pragma once
#include <string>
#include <unordered_map>

#include "file_server/event/Event.h"
#include "common/FeedbackInterface.h"
#include "common/Flags.h"
#include "common/Lock.h"
#include "pipeline/queue/QueueKey.h"

DECLARE_FLAG_INT32(max_block_event_timeout);

namespace logtail {

class BlockedEventManager : public FeedbackInterface {
public:
    static BlockedEventManager* GetInstance() {
        static BlockedEventManager* s_Instance = new BlockedEventManager;
        return s_Instance;
    }

    void UpdateBlockEvent(QueueKey logstoreKey,
                          const std::string& configName,
                          const Event& event,
                          const DevInode& devInode,
                          int32_t curTime);
    void GetTimeoutEvent(std::vector<Event*>& eventVec, int32_t curTime);
    void Feedback(int64_t key) override;

protected:
    struct BlockedEvent {
        BlockedEvent() : mInvalidTime(time(NULL)) {}
        void Update(QueueKey key, Event* pEvent, int32_t curTime) {
            if (mEvent != NULL) {
                // There are only two situations where event coverage is possible
                // 1. the new event is not timeout event
                // 2. old event is timeout event
                if (!pEvent->IsReaderFlushTimeout() || mEvent->IsReaderFlushTimeout()) {
                    delete mEvent;
                } else {
                    return;
                }
            }
            mEvent = pEvent;
            mQueueKey = key;
            // will become traditional block event if processor queue is not ready
            if (mEvent->IsReaderFlushTimeout()) {
                mTimeout = curTime - mInvalidTime;
            } else {
                mTimeout = (curTime - mInvalidTime) * 2 + 1;
                if (mTimeout > INT32_FLAG(max_block_event_timeout)) {
                    mTimeout = INT32_FLAG(max_block_event_timeout);
                }
            }
        }
        void SetInvalidAgain(int32_t curTime) {
            mTimeout *= 2;
            if (mTimeout > INT32_FLAG(max_block_event_timeout)) {
                mTimeout = INT32_FLAG(max_block_event_timeout);
            }
        }

        QueueKey mQueueKey = -1;
        Event* mEvent = nullptr;
        int32_t mInvalidTime;
        int32_t mTimeout = 1;
    };

    BlockedEventManager();
    virtual ~BlockedEventManager();

    std::unordered_map<int64_t, BlockedEvent> mBlockEventMap;
    SpinLock mLock;

private:
#ifdef APSARA_UNIT_TEST_MAIN
    friend class ForceReadUnittest;
#endif
};

} // namespace logtail
