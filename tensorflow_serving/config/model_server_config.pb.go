// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow_serving/config/model_server_config.proto

package tensorflow_serving

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"
import storage_path "github.com/maxhawkins/tensorflow-serving-go/tensorflow_serving/sources/storage_path"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The type of model.
// TODO(b/31336131): DEPRECATED.
type ModelType int32

const (
	ModelType_MODEL_TYPE_UNSPECIFIED ModelType = 0 // Deprecated: Do not use.
	ModelType_TENSORFLOW             ModelType = 1 // Deprecated: Do not use.
	ModelType_OTHER                  ModelType = 2 // Deprecated: Do not use.
)

var ModelType_name = map[int32]string{
	0: "MODEL_TYPE_UNSPECIFIED",
	1: "TENSORFLOW",
	2: "OTHER",
}
var ModelType_value = map[string]int32{
	"MODEL_TYPE_UNSPECIFIED": 0,
	"TENSORFLOW":             1,
	"OTHER":                  2,
}

func (x ModelType) String() string {
	return proto.EnumName(ModelType_name, int32(x))
}
func (ModelType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_model_server_config_edad896613326891, []int{0}
}

// Common configuration for loading a model being served.
type ModelConfig struct {
	// Name of the model.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Base path to the model, excluding the version directory.
	// E.g> for a model at /foo/bar/my_model/123, where 123 is the version, the
	// base path is /foo/bar/my_model.
	//
	// (This can be changed once a model is in serving, *if* the underlying data
	// remains the same. Otherwise there are no guarantees about whether the old
	// or new data will be used for model versions currently loaded.)
	BasePath string `protobuf:"bytes,2,opt,name=base_path,json=basePath,proto3" json:"base_path,omitempty"`
	// Type of model.
	// TODO(b/31336131): DEPRECATED. Please use 'model_platform' instead.
	ModelType ModelType `protobuf:"varint,3,opt,name=model_type,json=modelType,proto3,enum=tensorflow.serving.ModelType" json:"model_type,omitempty"` // Deprecated: Do not use.
	// Type of model (e.g. "tensorflow").
	//
	// (This cannot be changed once a model is in serving.)
	ModelPlatform string `protobuf:"bytes,4,opt,name=model_platform,json=modelPlatform,proto3" json:"model_platform,omitempty"`
	// Version policy for the model indicating which version(s) of the model to
	// load and make available for serving simultaneously.
	// The default option is to serve only the latest version of the model.
	//
	// (This can be changed once a model is in serving.)
	ModelVersionPolicy *storage_path.FileSystemStoragePathSourceConfig_ServableVersionPolicy `protobuf:"bytes,7,opt,name=model_version_policy,json=modelVersionPolicy,proto3" json:"model_version_policy,omitempty"`
	// String labels to associate with versions of the model, allowing inference
	// queries to refer to versions by label instead of number. Multiple labels
	// can map to the same version, but not vice-versa.
	//
	// An envisioned use-case for these labels is canarying tentative versions.
	// For example, one can assign labels "stable" and "canary" to two specific
	// versions. Perhaps initially "stable" is assigned to version 0 and "canary"
	// to version 1. Once version 1 passes canary, one can shift the "stable"
	// label to refer to version 1 (at that point both labels map to the same
	// version -- version 1 -- which is fine). Later once version 2 is ready to
	// canary one can move the "canary" label to version 2. And so on.
	VersionLabels map[string]int64 `protobuf:"bytes,8,rep,name=version_labels,json=versionLabels,proto3" json:"version_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// Configures logging requests and responses, to the model.
	//
	// (This can be changed once a model is in serving.)
	LoggingConfig        *LoggingConfig `protobuf:"bytes,6,opt,name=logging_config,json=loggingConfig,proto3" json:"logging_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ModelConfig) Reset()         { *m = ModelConfig{} }
func (m *ModelConfig) String() string { return proto.CompactTextString(m) }
func (*ModelConfig) ProtoMessage()    {}
func (*ModelConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_server_config_edad896613326891, []int{0}
}
func (m *ModelConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelConfig.Unmarshal(m, b)
}
func (m *ModelConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelConfig.Marshal(b, m, deterministic)
}
func (dst *ModelConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelConfig.Merge(dst, src)
}
func (m *ModelConfig) XXX_Size() int {
	return xxx_messageInfo_ModelConfig.Size(m)
}
func (m *ModelConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ModelConfig proto.InternalMessageInfo

func (m *ModelConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ModelConfig) GetBasePath() string {
	if m != nil {
		return m.BasePath
	}
	return ""
}

// Deprecated: Do not use.
func (m *ModelConfig) GetModelType() ModelType {
	if m != nil {
		return m.ModelType
	}
	return ModelType_MODEL_TYPE_UNSPECIFIED
}

func (m *ModelConfig) GetModelPlatform() string {
	if m != nil {
		return m.ModelPlatform
	}
	return ""
}

func (m *ModelConfig) GetModelVersionPolicy() *storage_path.FileSystemStoragePathSourceConfig_ServableVersionPolicy {
	if m != nil {
		return m.ModelVersionPolicy
	}
	return nil
}

func (m *ModelConfig) GetVersionLabels() map[string]int64 {
	if m != nil {
		return m.VersionLabels
	}
	return nil
}

func (m *ModelConfig) GetLoggingConfig() *LoggingConfig {
	if m != nil {
		return m.LoggingConfig
	}
	return nil
}

// Static list of models to be loaded for serving.
type ModelConfigList struct {
	Config               []*ModelConfig `protobuf:"bytes,1,rep,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ModelConfigList) Reset()         { *m = ModelConfigList{} }
func (m *ModelConfigList) String() string { return proto.CompactTextString(m) }
func (*ModelConfigList) ProtoMessage()    {}
func (*ModelConfigList) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_server_config_edad896613326891, []int{1}
}
func (m *ModelConfigList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelConfigList.Unmarshal(m, b)
}
func (m *ModelConfigList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelConfigList.Marshal(b, m, deterministic)
}
func (dst *ModelConfigList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelConfigList.Merge(dst, src)
}
func (m *ModelConfigList) XXX_Size() int {
	return xxx_messageInfo_ModelConfigList.Size(m)
}
func (m *ModelConfigList) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelConfigList.DiscardUnknown(m)
}

var xxx_messageInfo_ModelConfigList proto.InternalMessageInfo

func (m *ModelConfigList) GetConfig() []*ModelConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

// ModelServer config.
type ModelServerConfig struct {
	// ModelServer takes either a static file-based model config list or an Any
	// proto representing custom model config that is fetched dynamically at
	// runtime (through network RPC, custom service, etc.).
	//
	// Types that are valid to be assigned to Config:
	//	*ModelServerConfig_ModelConfigList
	//	*ModelServerConfig_CustomModelConfig
	Config               isModelServerConfig_Config `protobuf_oneof:"config"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ModelServerConfig) Reset()         { *m = ModelServerConfig{} }
func (m *ModelServerConfig) String() string { return proto.CompactTextString(m) }
func (*ModelServerConfig) ProtoMessage()    {}
func (*ModelServerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_server_config_edad896613326891, []int{2}
}
func (m *ModelServerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelServerConfig.Unmarshal(m, b)
}
func (m *ModelServerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelServerConfig.Marshal(b, m, deterministic)
}
func (dst *ModelServerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelServerConfig.Merge(dst, src)
}
func (m *ModelServerConfig) XXX_Size() int {
	return xxx_messageInfo_ModelServerConfig.Size(m)
}
func (m *ModelServerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelServerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ModelServerConfig proto.InternalMessageInfo

type isModelServerConfig_Config interface {
	isModelServerConfig_Config()
}

type ModelServerConfig_ModelConfigList struct {
	ModelConfigList *ModelConfigList `protobuf:"bytes,1,opt,name=model_config_list,json=modelConfigList,proto3,oneof"`
}

type ModelServerConfig_CustomModelConfig struct {
	CustomModelConfig *any.Any `protobuf:"bytes,2,opt,name=custom_model_config,json=customModelConfig,proto3,oneof"`
}

func (*ModelServerConfig_ModelConfigList) isModelServerConfig_Config() {}

func (*ModelServerConfig_CustomModelConfig) isModelServerConfig_Config() {}

func (m *ModelServerConfig) GetConfig() isModelServerConfig_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *ModelServerConfig) GetModelConfigList() *ModelConfigList {
	if x, ok := m.GetConfig().(*ModelServerConfig_ModelConfigList); ok {
		return x.ModelConfigList
	}
	return nil
}

func (m *ModelServerConfig) GetCustomModelConfig() *any.Any {
	if x, ok := m.GetConfig().(*ModelServerConfig_CustomModelConfig); ok {
		return x.CustomModelConfig
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ModelServerConfig) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ModelServerConfig_OneofMarshaler, _ModelServerConfig_OneofUnmarshaler, _ModelServerConfig_OneofSizer, []interface{}{
		(*ModelServerConfig_ModelConfigList)(nil),
		(*ModelServerConfig_CustomModelConfig)(nil),
	}
}

func _ModelServerConfig_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ModelServerConfig)
	// config
	switch x := m.Config.(type) {
	case *ModelServerConfig_ModelConfigList:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ModelConfigList); err != nil {
			return err
		}
	case *ModelServerConfig_CustomModelConfig:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CustomModelConfig); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ModelServerConfig.Config has unexpected type %T", x)
	}
	return nil
}

func _ModelServerConfig_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ModelServerConfig)
	switch tag {
	case 1: // config.model_config_list
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ModelConfigList)
		err := b.DecodeMessage(msg)
		m.Config = &ModelServerConfig_ModelConfigList{msg}
		return true, err
	case 2: // config.custom_model_config
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(any.Any)
		err := b.DecodeMessage(msg)
		m.Config = &ModelServerConfig_CustomModelConfig{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ModelServerConfig_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ModelServerConfig)
	// config
	switch x := m.Config.(type) {
	case *ModelServerConfig_ModelConfigList:
		s := proto.Size(x.ModelConfigList)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ModelServerConfig_CustomModelConfig:
		s := proto.Size(x.CustomModelConfig)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*ModelConfig)(nil), "tensorflow.serving.ModelConfig")
	proto.RegisterMapType((map[string]int64)(nil), "tensorflow.serving.ModelConfig.VersionLabelsEntry")
	proto.RegisterType((*ModelConfigList)(nil), "tensorflow.serving.ModelConfigList")
	proto.RegisterType((*ModelServerConfig)(nil), "tensorflow.serving.ModelServerConfig")
	proto.RegisterEnum("tensorflow.serving.ModelType", ModelType_name, ModelType_value)
}

func init() {
	proto.RegisterFile("tensorflow_serving/config/model_server_config.proto", fileDescriptor_model_server_config_edad896613326891)
}

var fileDescriptor_model_server_config_edad896613326891 = []byte{
	// 581 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x9d, 0xdb, 0xae, 0xb4, 0xb7, 0x6a, 0xd7, 0x99, 0x09, 0x85, 0x22, 0xa0, 0x0c, 0x21, 0x55,
	0x3c, 0x24, 0x52, 0xf6, 0x00, 0xe2, 0x69, 0x6c, 0x4b, 0xd5, 0x8d, 0x7e, 0x91, 0x14, 0xd0, 0x9e,
	0xa2, 0xb4, 0xb8, 0x59, 0x84, 0x13, 0x47, 0xb1, 0x5b, 0x94, 0x07, 0x7e, 0x05, 0x7f, 0x86, 0x9f,
	0xc6, 0x23, 0x8a, 0x9d, 0x42, 0xca, 0x3a, 0xed, 0x2d, 0x3e, 0x39, 0xf7, 0x9c, 0x7b, 0xee, 0xb5,
	0xe1, 0x44, 0x90, 0x88, 0xb3, 0x64, 0x49, 0xd9, 0x77, 0x97, 0x93, 0x64, 0x1d, 0x44, 0xbe, 0xb1,
	0x60, 0xd1, 0x32, 0xf0, 0x8d, 0x90, 0x7d, 0x25, 0x54, 0x82, 0x24, 0x71, 0x15, 0xa6, 0xc7, 0x09,
	0x13, 0x0c, 0xe3, 0x7f, 0x45, 0x7a, 0x5e, 0xd4, 0x79, 0xec, 0x33, 0xe6, 0x53, 0x62, 0x48, 0xc6,
	0x7c, 0xb5, 0x34, 0xbc, 0x28, 0x55, 0xf4, 0x8e, 0x7e, 0xb7, 0x07, 0x65, 0xbe, 0x1f, 0x44, 0xfe,
	0x96, 0x7c, 0x67, 0xb4, 0x83, 0xcf, 0xd9, 0x2a, 0x59, 0x10, 0x6e, 0x70, 0xc1, 0x12, 0xcf, 0x27,
	0x6e, 0xec, 0x89, 0x1b, 0x63, 0x19, 0x50, 0xe2, 0xf2, 0x94, 0x0b, 0x12, 0xba, 0xc5, 0x1f, 0xae,
	0x62, 0x2b, 0xb9, 0xe3, 0x9f, 0x15, 0x68, 0x8c, 0xb2, 0x2c, 0xe7, 0xd2, 0x04, 0x63, 0xa8, 0x44,
	0x5e, 0x48, 0x34, 0xd4, 0x45, 0xbd, 0xba, 0x2d, 0xbf, 0xf1, 0x13, 0xa8, 0xcf, 0x3d, 0xae, 0xaa,
	0xb5, 0x92, 0xfc, 0x51, 0xcb, 0x80, 0xa9, 0x27, 0x6e, 0xf0, 0x29, 0x80, 0x9a, 0x85, 0x48, 0x63,
	0xa2, 0x95, 0xbb, 0xa8, 0xd7, 0x32, 0x9f, 0xea, 0xb7, 0x67, 0xa0, 0x4b, 0x97, 0x59, 0x1a, 0x93,
	0xb3, 0x92, 0x86, 0xec, 0x7a, 0xb8, 0x39, 0xe2, 0x57, 0xd0, 0x52, 0x0a, 0x31, 0xf5, 0xc4, 0x92,
	0x25, 0xa1, 0x56, 0x91, 0x1e, 0x4d, 0x89, 0x4e, 0x73, 0x10, 0xff, 0x80, 0x23, 0x45, 0x5b, 0x93,
	0x84, 0x07, 0x2c, 0x72, 0x63, 0x46, 0x83, 0x45, 0xaa, 0x3d, 0xe8, 0xa2, 0x5e, 0xc3, 0xfc, 0xb0,
	0xcb, 0xb2, 0x1f, 0x50, 0xe2, 0xc8, 0x09, 0x38, 0x6a, 0x00, 0x59, 0xc7, 0x8e, 0x8c, 0xaf, 0xe2,
	0xea, 0x0e, 0x49, 0xd6, 0xde, 0x9c, 0x92, 0xcf, 0x4a, 0x73, 0x2a, 0x25, 0x6d, 0x2c, 0x8d, 0xb6,
	0x30, 0x7c, 0x0d, 0xad, 0x8d, 0x31, 0xf5, 0xe6, 0x84, 0x72, 0xad, 0xd6, 0x2d, 0xf7, 0x1a, 0xa6,
	0x79, 0x67, 0xd6, 0xdc, 0x22, 0x97, 0x19, 0xca, 0x22, 0x2b, 0x12, 0x49, 0x6a, 0x37, 0xd7, 0x45,
	0x0c, 0x0f, 0xa0, 0xb5, 0xbd, 0x6a, 0xad, 0x2a, 0x33, 0xbd, 0xd8, 0x25, 0x3d, 0x54, 0x4c, 0x25,
	0x6e, 0x37, 0x69, 0xf1, 0xd8, 0x39, 0x05, 0x7c, 0xdb, 0x0e, 0xb7, 0xa1, 0xfc, 0x8d, 0xa4, 0xf9,
	0x4a, 0xb3, 0x4f, 0x7c, 0x04, 0xfb, 0x6b, 0x8f, 0xae, 0x88, 0xdc, 0x66, 0xd9, 0x56, 0x87, 0x77,
	0xa5, 0xb7, 0xe8, 0xaa, 0x52, 0xdb, 0x6f, 0x57, 0x8f, 0xaf, 0xe0, 0xa0, 0x10, 0x61, 0x18, 0x70,
	0x81, 0xdf, 0x40, 0x35, 0x6f, 0x0e, 0xc9, 0xdc, 0xcf, 0xef, 0xc9, 0x6d, 0xe7, 0xf4, 0xe3, 0x5f,
	0x08, 0x0e, 0x25, 0xee, 0xc8, 0xc7, 0x92, 0xdf, 0xb3, 0x8f, 0x70, 0xa8, 0xb6, 0xa9, 0x58, 0x2e,
	0x0d, 0xb8, 0x90, 0x1d, 0x36, 0xcc, 0x97, 0xf7, 0x28, 0x67, 0xed, 0x0c, 0xf6, 0xec, 0x83, 0xf0,
	0xbf, 0x0e, 0xfb, 0xf0, 0x70, 0xb1, 0xe2, 0x82, 0x85, 0x6e, 0x51, 0x59, 0x46, 0x6c, 0x98, 0x47,
	0xba, 0x7a, 0x82, 0xfa, 0xe6, 0x09, 0xea, 0xef, 0xa3, 0x74, 0xb0, 0x67, 0x1f, 0xaa, 0x92, 0x82,
	0xfc, 0x59, 0x6d, 0x93, 0xf4, 0xf5, 0x18, 0xea, 0x7f, 0x6f, 0x2d, 0x7e, 0x06, 0x8f, 0x46, 0x93,
	0x0b, 0x6b, 0xe8, 0xce, 0xae, 0xa7, 0x96, 0xfb, 0x69, 0xec, 0x4c, 0xad, 0xf3, 0xcb, 0xfe, 0xa5,
	0x75, 0xd1, 0xde, 0xeb, 0x94, 0x6a, 0x08, 0x63, 0x80, 0x99, 0x35, 0x76, 0x26, 0x76, 0x7f, 0x38,
	0xf9, 0xd2, 0x46, 0x12, 0x6b, 0xc2, 0xfe, 0x64, 0x36, 0xb0, 0xec, 0x76, 0x29, 0x3b, 0x9e, 0x95,
	0x7f, 0x23, 0x34, 0xaf, 0xca, 0x0e, 0x4e, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x26, 0x49, 0xef,
	0x06, 0x5d, 0x04, 0x00, 0x00,
}