// Code generated by protoc-gen-go.
// source: myprobe/myprobe.proto
// DO NOT EDIT!

/*
Package myprobe is a generated protocol buffer package.

It is generated from these files:
	myprobe/myprobe.proto

It has these top-level messages:
	ProbeConf
*/
package myprobe

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import cloudprober_probes "github.com/google/cloudprober/probes/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Redis operation
type ProbeConf_Op int32

const (
	ProbeConf_GET    ProbeConf_Op = 0
	ProbeConf_SET    ProbeConf_Op = 1
	ProbeConf_DELETE ProbeConf_Op = 2
)

var ProbeConf_Op_name = map[int32]string{
	0: "GET",
	1: "SET",
	2: "DELETE",
}
var ProbeConf_Op_value = map[string]int32{
	"GET":    0,
	"SET":    1,
	"DELETE": 2,
}

func (x ProbeConf_Op) Enum() *ProbeConf_Op {
	p := new(ProbeConf_Op)
	*p = x
	return p
}
func (x ProbeConf_Op) String() string {
	return proto.EnumName(ProbeConf_Op_name, int32(x))
}
func (x *ProbeConf_Op) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ProbeConf_Op_value, data, "ProbeConf_Op")
	if err != nil {
		return err
	}
	*x = ProbeConf_Op(value)
	return nil
}
func (ProbeConf_Op) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type ProbeConf struct {
	Op *ProbeConf_Op `protobuf:"varint,1,req,name=op,enum=myprober.ProbeConf_Op" json:"op,omitempty"`
	// Key and value for the redis operation
	Key              *string `protobuf:"bytes,2,req,name=key" json:"key,omitempty"`
	Value            *string `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ProbeConf) Reset()                    { *m = ProbeConf{} }
func (m *ProbeConf) String() string            { return proto.CompactTextString(m) }
func (*ProbeConf) ProtoMessage()               {}
func (*ProbeConf) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ProbeConf) GetOp() ProbeConf_Op {
	if m != nil && m.Op != nil {
		return *m.Op
	}
	return ProbeConf_GET
}

func (m *ProbeConf) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *ProbeConf) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

var E_RedisProbe = &proto.ExtensionDesc{
	ExtendedType:  (*cloudprober_probes.ProbeDef)(nil),
	ExtensionType: (*ProbeConf)(nil),
	Field:         200,
	Name:          "myprober.redis_probe",
	Tag:           "bytes,200,opt,name=redis_probe,json=redisProbe",
	Filename:      "myprobe/myprobe.proto",
}

func init() {
	proto.RegisterType((*ProbeConf)(nil), "myprober.ProbeConf")
	proto.RegisterEnum("myprober.ProbeConf_Op", ProbeConf_Op_name, ProbeConf_Op_value)
	proto.RegisterExtension(E_RedisProbe)
}

func init() { proto.RegisterFile("myprobe/myprobe.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcd, 0xad, 0x2c, 0x28,
	0xca, 0x4f, 0x4a, 0xd5, 0x87, 0xd2, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x1c, 0x50, 0x6e,
	0x91, 0x94, 0x79, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x7a, 0x7e,
	0x7e, 0x7a, 0x4e, 0xaa, 0x7e, 0x72, 0x4e, 0x7e, 0x69, 0x0a, 0x44, 0x5a, 0x1f, 0x4c, 0x15, 0xeb,
	0x83, 0x75, 0xe9, 0x27, 0xe7, 0xe7, 0xa5, 0x65, 0xa6, 0x43, 0x8c, 0x50, 0xaa, 0xe7, 0xe2, 0x0c,
	0x00, 0x49, 0x3a, 0xe7, 0xe7, 0xa5, 0x09, 0xa9, 0x71, 0x31, 0xe5, 0x17, 0x48, 0x30, 0x2a, 0x30,
	0x69, 0xf0, 0x19, 0x89, 0xe9, 0xc1, 0x0c, 0xd7, 0x83, 0x2b, 0xd0, 0xf3, 0x2f, 0x08, 0x62, 0xca,
	0x2f, 0x10, 0x12, 0xe0, 0x62, 0xce, 0x4e, 0xad, 0x94, 0x60, 0x52, 0x60, 0xd2, 0xe0, 0x0c, 0x02,
	0x31, 0x85, 0x44, 0xb8, 0x58, 0xcb, 0x12, 0x73, 0x4a, 0x53, 0x25, 0x98, 0x15, 0x18, 0x35, 0x38,
	0x83, 0x20, 0x1c, 0x25, 0x25, 0x2e, 0x26, 0xff, 0x02, 0x21, 0x76, 0x2e, 0x66, 0x77, 0xd7, 0x10,
	0x01, 0x06, 0x10, 0x23, 0xd8, 0x35, 0x44, 0x80, 0x51, 0x88, 0x8b, 0x8b, 0xcd, 0xc5, 0xd5, 0xc7,
	0x35, 0xc4, 0x55, 0x80, 0xc9, 0x2a, 0x98, 0x8b, 0xbb, 0x28, 0x35, 0x25, 0xb3, 0x38, 0x1e, 0x6c,
	0x99, 0x90, 0x8c, 0x1e, 0x92, 0xbb, 0xf5, 0x20, 0xee, 0x86, 0x38, 0xc0, 0x25, 0x35, 0x4d, 0xe2,
	0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0x30, 0x16, 0xb7, 0x05, 0x71, 0x81, 0x8d, 0x01, 0xf3,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x12, 0xc5, 0x1f, 0xa5, 0x30, 0x01, 0x00, 0x00,
}