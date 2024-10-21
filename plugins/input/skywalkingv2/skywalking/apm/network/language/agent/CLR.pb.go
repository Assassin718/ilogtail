//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: common/CLR.proto

package agent

import (
	common "github.com/alibaba/ilogtail/plugins/input/skywalkingv2/skywalking/apm/network/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CLRMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time   int64       `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Cpu    *common.CPU `protobuf:"bytes,2,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Gc     *ClrGC      `protobuf:"bytes,3,opt,name=gc,proto3" json:"gc,omitempty"`
	Thread *ClrThread  `protobuf:"bytes,4,opt,name=thread,proto3" json:"thread,omitempty"`
}

func (x *CLRMetric) Reset() {
	*x = CLRMetric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_CLR_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CLRMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CLRMetric) ProtoMessage() {}

func (x *CLRMetric) ProtoReflect() protoreflect.Message {
	mi := &file_common_CLR_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CLRMetric.ProtoReflect.Descriptor instead.
func (*CLRMetric) Descriptor() ([]byte, []int) {
	return file_common_CLR_proto_rawDescGZIP(), []int{0}
}

func (x *CLRMetric) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *CLRMetric) GetCpu() *common.CPU {
	if x != nil {
		return x.Cpu
	}
	return nil
}

func (x *CLRMetric) GetGc() *ClrGC {
	if x != nil {
		return x.Gc
	}
	return nil
}

func (x *CLRMetric) GetThread() *ClrThread {
	if x != nil {
		return x.Thread
	}
	return nil
}

type ClrGC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gen0CollectCount int64 `protobuf:"varint,1,opt,name=Gen0CollectCount,proto3" json:"Gen0CollectCount,omitempty"`
	Gen1CollectCount int64 `protobuf:"varint,2,opt,name=Gen1CollectCount,proto3" json:"Gen1CollectCount,omitempty"`
	Gen2CollectCount int64 `protobuf:"varint,3,opt,name=Gen2CollectCount,proto3" json:"Gen2CollectCount,omitempty"`
	HeapMemory       int64 `protobuf:"varint,4,opt,name=HeapMemory,proto3" json:"HeapMemory,omitempty"`
}

func (x *ClrGC) Reset() {
	*x = ClrGC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_CLR_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClrGC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClrGC) ProtoMessage() {}

func (x *ClrGC) ProtoReflect() protoreflect.Message {
	mi := &file_common_CLR_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClrGC.ProtoReflect.Descriptor instead.
func (*ClrGC) Descriptor() ([]byte, []int) {
	return file_common_CLR_proto_rawDescGZIP(), []int{1}
}

func (x *ClrGC) GetGen0CollectCount() int64 {
	if x != nil {
		return x.Gen0CollectCount
	}
	return 0
}

func (x *ClrGC) GetGen1CollectCount() int64 {
	if x != nil {
		return x.Gen1CollectCount
	}
	return 0
}

func (x *ClrGC) GetGen2CollectCount() int64 {
	if x != nil {
		return x.Gen2CollectCount
	}
	return 0
}

func (x *ClrGC) GetHeapMemory() int64 {
	if x != nil {
		return x.HeapMemory
	}
	return 0
}

type ClrThread struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvailableCompletionPortThreads int32 `protobuf:"varint,1,opt,name=AvailableCompletionPortThreads,proto3" json:"AvailableCompletionPortThreads,omitempty"`
	AvailableWorkerThreads         int32 `protobuf:"varint,2,opt,name=AvailableWorkerThreads,proto3" json:"AvailableWorkerThreads,omitempty"`
	MaxCompletionPortThreads       int32 `protobuf:"varint,3,opt,name=MaxCompletionPortThreads,proto3" json:"MaxCompletionPortThreads,omitempty"`
	MaxWorkerThreads               int32 `protobuf:"varint,4,opt,name=MaxWorkerThreads,proto3" json:"MaxWorkerThreads,omitempty"`
}

func (x *ClrThread) Reset() {
	*x = ClrThread{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_CLR_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClrThread) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClrThread) ProtoMessage() {}

func (x *ClrThread) ProtoReflect() protoreflect.Message {
	mi := &file_common_CLR_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClrThread.ProtoReflect.Descriptor instead.
func (*ClrThread) Descriptor() ([]byte, []int) {
	return file_common_CLR_proto_rawDescGZIP(), []int{2}
}

func (x *ClrThread) GetAvailableCompletionPortThreads() int32 {
	if x != nil {
		return x.AvailableCompletionPortThreads
	}
	return 0
}

func (x *ClrThread) GetAvailableWorkerThreads() int32 {
	if x != nil {
		return x.AvailableWorkerThreads
	}
	return 0
}

func (x *ClrThread) GetMaxCompletionPortThreads() int32 {
	if x != nil {
		return x.MaxCompletionPortThreads
	}
	return 0
}

func (x *ClrThread) GetMaxWorkerThreads() int32 {
	if x != nil {
		return x.MaxWorkerThreads
	}
	return 0
}

var File_common_CLR_proto protoreflect.FileDescriptor

var file_common_CLR_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x43, 0x4c, 0x52, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x09, 0x43, 0x4c, 0x52, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x03, 0x63, 0x70, 0x75, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x04, 0x2e, 0x43, 0x50, 0x55, 0x52, 0x03, 0x63, 0x70, 0x75,
	0x12, 0x16, 0x0a, 0x02, 0x67, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x43,
	0x6c, 0x72, 0x47, 0x43, 0x52, 0x02, 0x67, 0x63, 0x12, 0x22, 0x0a, 0x06, 0x74, 0x68, 0x72, 0x65,
	0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x43, 0x6c, 0x72, 0x54, 0x68,
	0x72, 0x65, 0x61, 0x64, 0x52, 0x06, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x22, 0xab, 0x01, 0x0a,
	0x05, 0x43, 0x6c, 0x72, 0x47, 0x43, 0x12, 0x2a, 0x0a, 0x10, 0x47, 0x65, 0x6e, 0x30, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x10, 0x47, 0x65, 0x6e, 0x30, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x47, 0x65, 0x6e, 0x31, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x47, 0x65,
	0x6e, 0x31, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a,
	0x0a, 0x10, 0x47, 0x65, 0x6e, 0x32, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x47, 0x65, 0x6e, 0x32, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x48, 0x65,
	0x61, 0x70, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x48, 0x65, 0x61, 0x70, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x22, 0xf3, 0x01, 0x0a, 0x09, 0x43,
	0x6c, 0x72, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x12, 0x46, 0x0a, 0x1e, 0x41, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x6f, 0x72, 0x74, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x1e, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73,
	0x12, 0x36, 0x0a, 0x16, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x57, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x16, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x12, 0x3a, 0x0a, 0x18, 0x4d, 0x61, 0x78, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x54, 0x68, 0x72,
	0x65, 0x61, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x18, 0x4d, 0x61, 0x78, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x54, 0x68, 0x72,
	0x65, 0x61, 0x64, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x4d, 0x61, 0x78, 0x57, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10,
	0x4d, 0x61, 0x78, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73,
	0x42, 0xa1, 0x01, 0x0a, 0x30, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e,
	0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x6d, 0x2e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x01, 0x5a, 0x4e, 0x6c, 0x6f, 0x67, 0x74, 0x61, 0x69, 0x6c,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x2f, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x76,
	0x32, 0x2f, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x6d,
	0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0xaa, 0x02, 0x1a, 0x53, 0x6b, 0x79, 0x57, 0x61, 0x6c,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_CLR_proto_rawDescOnce sync.Once
	file_common_CLR_proto_rawDescData = file_common_CLR_proto_rawDesc
)

func file_common_CLR_proto_rawDescGZIP() []byte {
	file_common_CLR_proto_rawDescOnce.Do(func() {
		file_common_CLR_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_CLR_proto_rawDescData)
	})
	return file_common_CLR_proto_rawDescData
}

var file_common_CLR_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_common_CLR_proto_goTypes = []interface{}{
	(*CLRMetric)(nil),  // 0: CLRMetric
	(*ClrGC)(nil),      // 1: ClrGC
	(*ClrThread)(nil),  // 2: ClrThread
	(*common.CPU)(nil), // 3: CPU
}
var file_common_CLR_proto_depIdxs = []int32{
	3, // 0: CLRMetric.cpu:type_name -> CPU
	1, // 1: CLRMetric.gc:type_name -> ClrGC
	2, // 2: CLRMetric.thread:type_name -> ClrThread
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_common_CLR_proto_init() }
func file_common_CLR_proto_init() {
	if File_common_CLR_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_CLR_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CLRMetric); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_CLR_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClrGC); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_CLR_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClrThread); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_CLR_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_CLR_proto_goTypes,
		DependencyIndexes: file_common_CLR_proto_depIdxs,
		MessageInfos:      file_common_CLR_proto_msgTypes,
	}.Build()
	File_common_CLR_proto = out.File
	file_common_CLR_proto_rawDesc = nil
	file_common_CLR_proto_goTypes = nil
	file_common_CLR_proto_depIdxs = nil
}