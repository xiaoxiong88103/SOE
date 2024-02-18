// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: config.proto

package config

import (
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

// 空消息，用于不带任何参数的RPC调用。
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{0}
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string  `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"` // 返回的字符串
	Time     float32 `protobuf:"fixed32,2,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *Response) GetTime() float32 {
	if x != nil {
		return x.Time
	}
	return 0
}

// CPU 信息消息，包括核心数、型号、频率等具体信息。
type HardwareInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumCores        int32    `protobuf:"varint,1,opt,name=num_cores,json=numCores,proto3" json:"num_cores,omitempty"`                         // CPU 核心数
	ModelName       string   `protobuf:"bytes,2,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"`                       // CPU 型号名称
	MaxFrequencyGhz float32  `protobuf:"fixed32,3,opt,name=max_frequency_ghz,json=maxFrequencyGhz,proto3" json:"max_frequency_ghz,omitempty"` // 最大频率（GHz）
	TotalGb         float32  `protobuf:"fixed32,4,opt,name=total_gb,json=totalGb,proto3" json:"total_gb,omitempty"`                           // 内存总量（以GB为单位）
	DiskSizeGb      []string `protobuf:"bytes,5,rep,name=disk_size_gb,json=diskSizeGb,proto3" json:"disk_size_gb,omitempty"`                  // 磁盘大小（以GB为单位）
	NumBlocks       int32    `protobuf:"varint,6,opt,name=num_blocks,json=numBlocks,proto3" json:"num_blocks,omitempty"`                      // 磁盘块数
}

func (x *HardwareInfo) Reset() {
	*x = HardwareInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HardwareInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HardwareInfo) ProtoMessage() {}

func (x *HardwareInfo) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HardwareInfo.ProtoReflect.Descriptor instead.
func (*HardwareInfo) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{2}
}

func (x *HardwareInfo) GetNumCores() int32 {
	if x != nil {
		return x.NumCores
	}
	return 0
}

func (x *HardwareInfo) GetModelName() string {
	if x != nil {
		return x.ModelName
	}
	return ""
}

func (x *HardwareInfo) GetMaxFrequencyGhz() float32 {
	if x != nil {
		return x.MaxFrequencyGhz
	}
	return 0
}

func (x *HardwareInfo) GetTotalGb() float32 {
	if x != nil {
		return x.TotalGb
	}
	return 0
}

func (x *HardwareInfo) GetDiskSizeGb() []string {
	if x != nil {
		return x.DiskSizeGb
	}
	return nil
}

func (x *HardwareInfo) GetNumBlocks() int32 {
	if x != nil {
		return x.NumBlocks
	}
	return 0
}

// 系统信息消息，包括CPU使用率、内存使用率、VPU使用百分比、NPU使用百分比、GPU使用百分比和IO使用率（读取和写入）。
type SystemInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BandwidthUsagePerSec        float32   `protobuf:"fixed32,1,opt,name=bandwidth_usage_per_sec,json=bandwidthUsagePerSec,proto3" json:"bandwidth_usage_per_sec,omitempty"`                       // 每秒带宽使用率
	CpuUsagePercent             float32   `protobuf:"fixed32,2,opt,name=cpu_usage_percent,json=cpuUsagePercent,proto3" json:"cpu_usage_percent,omitempty"`                                        // CPU使用率
	MemoryUsagePercent          float32   `protobuf:"fixed32,3,opt,name=memory_usage_percent,json=memoryUsagePercent,proto3" json:"memory_usage_percent,omitempty"`                               // 内存使用率
	VpuUsagePercent             float32   `protobuf:"fixed32,4,opt,name=vpu_usage_percent,json=vpuUsagePercent,proto3" json:"vpu_usage_percent,omitempty"`                                        // VPU使用百分比
	NpuUsagePercent             []float32 `protobuf:"fixed32,5,rep,packed,name=npu_usage_percent,json=npuUsagePercent,proto3" json:"npu_usage_percent,omitempty"`                                 // NPU使用百分比
	GpuUsagePercent             float32   `protobuf:"fixed32,6,opt,name=gpu_usage_percent,json=gpuUsagePercent,proto3" json:"gpu_usage_percent,omitempty"`                                        // GPU使用百分比
	IoReadUsagePercent          float32   `protobuf:"fixed32,7,opt,name=io_read_usage_percent,json=ioReadUsagePercent,proto3" json:"io_read_usage_percent,omitempty"`                             // IO读取使用率
	IoWriteUsagePercent         float32   `protobuf:"fixed32,8,opt,name=io_write_usage_percent,json=ioWriteUsagePercent,proto3" json:"io_write_usage_percent,omitempty"`                          // IO写入使用率
	NetworkUploadUsagePercent   float32   `protobuf:"fixed32,9,opt,name=network_upload_usage_percent,json=networkUploadUsagePercent,proto3" json:"network_upload_usage_percent,omitempty"`        // 网络上传使用率
	NetworkDownloadUsagePercent float32   `protobuf:"fixed32,10,opt,name=network_download_usage_percent,json=networkDownloadUsagePercent,proto3" json:"network_download_usage_percent,omitempty"` // 网络下载使用率
	NetworkConnections          int64     `protobuf:"varint,11,opt,name=network_connections,json=networkConnections,proto3" json:"network_connections,omitempty"`                                 // 当前网络连接数
	SystemLoadAvg               float32   `protobuf:"fixed32,12,opt,name=system_load_avg,json=systemLoadAvg,proto3" json:"system_load_avg,omitempty"`                                             // 系统平均负载
	DiskSizeGbShengyu           string    `protobuf:"bytes,13,opt,name=disk_size_gb_shengyu,json=diskSizeGbShengyu,proto3" json:"disk_size_gb_shengyu,omitempty"`                                 // 磁盘剩余大小（以GB为单位）
	Time                        string    `protobuf:"bytes,14,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *SystemInfo) Reset() {
	*x = SystemInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemInfo) ProtoMessage() {}

func (x *SystemInfo) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemInfo.ProtoReflect.Descriptor instead.
func (*SystemInfo) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{3}
}

func (x *SystemInfo) GetBandwidthUsagePerSec() float32 {
	if x != nil {
		return x.BandwidthUsagePerSec
	}
	return 0
}

func (x *SystemInfo) GetCpuUsagePercent() float32 {
	if x != nil {
		return x.CpuUsagePercent
	}
	return 0
}

func (x *SystemInfo) GetMemoryUsagePercent() float32 {
	if x != nil {
		return x.MemoryUsagePercent
	}
	return 0
}

func (x *SystemInfo) GetVpuUsagePercent() float32 {
	if x != nil {
		return x.VpuUsagePercent
	}
	return 0
}

func (x *SystemInfo) GetNpuUsagePercent() []float32 {
	if x != nil {
		return x.NpuUsagePercent
	}
	return nil
}

func (x *SystemInfo) GetGpuUsagePercent() float32 {
	if x != nil {
		return x.GpuUsagePercent
	}
	return 0
}

func (x *SystemInfo) GetIoReadUsagePercent() float32 {
	if x != nil {
		return x.IoReadUsagePercent
	}
	return 0
}

func (x *SystemInfo) GetIoWriteUsagePercent() float32 {
	if x != nil {
		return x.IoWriteUsagePercent
	}
	return 0
}

func (x *SystemInfo) GetNetworkUploadUsagePercent() float32 {
	if x != nil {
		return x.NetworkUploadUsagePercent
	}
	return 0
}

func (x *SystemInfo) GetNetworkDownloadUsagePercent() float32 {
	if x != nil {
		return x.NetworkDownloadUsagePercent
	}
	return 0
}

func (x *SystemInfo) GetNetworkConnections() int64 {
	if x != nil {
		return x.NetworkConnections
	}
	return 0
}

func (x *SystemInfo) GetSystemLoadAvg() float32 {
	if x != nil {
		return x.SystemLoadAvg
	}
	return 0
}

func (x *SystemInfo) GetDiskSizeGbShengyu() string {
	if x != nil {
		return x.DiskSizeGbShengyu
	}
	return ""
}

func (x *SystemInfo) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

var File_config_proto protoreflect.FileDescriptor

var file_config_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x3a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0xd2, 0x01, 0x0a,
	0x0c, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a,
	0x09, 0x6e, 0x75, 0x6d, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x6e, 0x75, 0x6d, 0x43, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x6d, 0x61, 0x78,
	0x5f, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x67, 0x68, 0x7a, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x0f, 0x6d, 0x61, 0x78, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x79, 0x47, 0x68, 0x7a, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x67,
	0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x47, 0x62,
	0x12, 0x20, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x67, 0x62,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x69, 0x73, 0x6b, 0x53, 0x69, 0x7a, 0x65,
	0x47, 0x62, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x75, 0x6d, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6e, 0x75, 0x6d, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x73, 0x22, 0xb1, 0x05, 0x0a, 0x0a, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x35, 0x0a, 0x17, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x5f, 0x75, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x14, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x70, 0x75, 0x5f, 0x75,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0f, 0x63, 0x70, 0x75, 0x55, 0x73, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x75, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x12, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x61, 0x67, 0x65, 0x50, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x76, 0x70, 0x75, 0x5f, 0x75, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0f, 0x76, 0x70, 0x75, 0x55, 0x73, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x12, 0x2a, 0x0a, 0x11, 0x6e, 0x70, 0x75, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x70,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x02, 0x52, 0x0f, 0x6e, 0x70,
	0x75, 0x55, 0x73, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x2a, 0x0a,
	0x11, 0x67, 0x70, 0x75, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0f, 0x67, 0x70, 0x75, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x31, 0x0a, 0x15, 0x69, 0x6f, 0x5f,
	0x72, 0x65, 0x61, 0x64, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x12, 0x69, 0x6f, 0x52, 0x65, 0x61, 0x64,
	0x55, 0x73, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x16,
	0x69, 0x6f, 0x5f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x70,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x13, 0x69, 0x6f,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x12, 0x3f, 0x0a, 0x1c, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x75, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x19, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x73, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x12, 0x43, 0x0a, 0x1e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x64, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x1b, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x13, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x5f, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x61, 0x76, 0x67, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x67,
	0x12, 0x2f, 0x0a, 0x14, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x67, 0x62,
	0x5f, 0x73, 0x68, 0x65, 0x6e, 0x67, 0x79, 0x75, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x64, 0x69, 0x73, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x47, 0x62, 0x53, 0x68, 0x65, 0x6e, 0x67, 0x79,
	0x75, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x32, 0x89, 0x01, 0x0a, 0x0d, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x3d, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x68, 0x61,
	0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x15, 0x2e, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x2e, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x1a, 0x11, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x13, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x11, 0x2e, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_proto_rawDescOnce sync.Once
	file_config_proto_rawDescData = file_config_proto_rawDesc
)

func file_config_proto_rawDescGZIP() []byte {
	file_config_proto_rawDescOnce.Do(func() {
		file_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_proto_rawDescData)
	})
	return file_config_proto_rawDescData
}

var file_config_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_config_proto_goTypes = []interface{}{
	(*Empty)(nil),        // 0: monitor.Empty
	(*Response)(nil),     // 1: monitor.Response
	(*HardwareInfo)(nil), // 2: monitor.hardwareInfo
	(*SystemInfo)(nil),   // 3: monitor.SystemInfo
}
var file_config_proto_depIdxs = []int32{
	2, // 0: monitor.SystemMetrics.GethardwareInfo:input_type -> monitor.hardwareInfo
	3, // 1: monitor.SystemMetrics.GetSystemInfo:input_type -> monitor.SystemInfo
	1, // 2: monitor.SystemMetrics.GethardwareInfo:output_type -> monitor.Response
	1, // 3: monitor.SystemMetrics.GetSystemInfo:output_type -> monitor.Response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_config_proto_init() }
func file_config_proto_init() {
	if File_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HardwareInfo); i {
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
		file_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemInfo); i {
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
			RawDescriptor: file_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_config_proto_goTypes,
		DependencyIndexes: file_config_proto_depIdxs,
		MessageInfos:      file_config_proto_msgTypes,
	}.Build()
	File_config_proto = out.File
	file_config_proto_rawDesc = nil
	file_config_proto_goTypes = nil
	file_config_proto_depIdxs = nil
}