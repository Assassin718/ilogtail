# Copyright 2023 iLogtail Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

cmake_minimum_required(VERSION 3.22)
project(pipeline_unittest)

add_executable(global_config_unittest GlobalConfigUnittest.cpp)
target_link_libraries(global_config_unittest ${UT_BASE_TARGET})

add_executable(pipeline_unittest PipelineUnittest.cpp)
target_link_libraries(pipeline_unittest ${UT_BASE_TARGET})

add_executable(pipeline_manager_unittest PipelineManagerUnittest.cpp)
target_link_libraries(pipeline_manager_unittest ${UT_BASE_TARGET})

add_executable(concurrency_limiter_unittest ConcurrencyLimiterUnittest.cpp)
target_link_libraries(concurrency_limiter_unittest ${UT_BASE_TARGET})

add_executable(pipeline_update_unittest PipelineUpdateUnittest.cpp)
target_link_libraries(pipeline_update_unittest ${UT_BASE_TARGET})

include(GoogleTest)
gtest_discover_tests(global_config_unittest)
gtest_discover_tests(pipeline_unittest)
gtest_discover_tests(pipeline_manager_unittest)
gtest_discover_tests(concurrency_limiter_unittest)
gtest_discover_tests(pipeline_update_unittest)

