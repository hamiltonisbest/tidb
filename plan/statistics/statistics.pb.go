// Code generated by protoc-gen-go.
// source: statistics.proto
// DO NOT EDIT!

/*
Package statistics is a generated protocol buffer package.

It is generated from these files:
	statistics.proto

It has these top-level messages:
	ColumnPB
	TablePB
*/
package statistics

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type ColumnPB struct {
	Id               *int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Ndv              *int64  `protobuf:"varint,2,opt,name=ndv" json:"ndv,omitempty"`
	Numbers          []int64 `protobuf:"varint,3,rep,name=numbers" json:"numbers,omitempty"`
	Value            []byte  `protobuf:"bytes,4,opt,name=value" json:"value,omitempty"`
	Repeats          []int64 `protobuf:"varint,5,rep,name=repeats" json:"repeats,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ColumnPB) Reset()                    { *m = ColumnPB{} }
func (m *ColumnPB) String() string            { return proto.CompactTextString(m) }
func (*ColumnPB) ProtoMessage()               {}
func (*ColumnPB) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ColumnPB) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *ColumnPB) GetNdv() int64 {
	if m != nil && m.Ndv != nil {
		return *m.Ndv
	}
	return 0
}

func (m *ColumnPB) GetNumbers() []int64 {
	if m != nil {
		return m.Numbers
	}
	return nil
}

func (m *ColumnPB) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *ColumnPB) GetRepeats() []int64 {
	if m != nil {
		return m.Repeats
	}
	return nil
}

type TablePB struct {
	Id               *int64      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Ts               *int64      `protobuf:"varint,2,opt,name=ts" json:"ts,omitempty"`
	Count            *int64      `protobuf:"varint,3,opt,name=count" json:"count,omitempty"`
	Columns          []*ColumnPB `protobuf:"bytes,4,rep,name=columns" json:"columns,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *TablePB) Reset()                    { *m = TablePB{} }
func (m *TablePB) String() string            { return proto.CompactTextString(m) }
func (*TablePB) ProtoMessage()               {}
func (*TablePB) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TablePB) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *TablePB) GetTs() int64 {
	if m != nil && m.Ts != nil {
		return *m.Ts
	}
	return 0
}

func (m *TablePB) GetCount() int64 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

func (m *TablePB) GetColumns() []*ColumnPB {
	if m != nil {
		return m.Columns
	}
	return nil
}

func init() {
	proto.RegisterType((*ColumnPB)(nil), "statistics.ColumnPB")
	proto.RegisterType((*TablePB)(nil), "statistics.TablePB")
}

var fileDescriptor0 = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x2e, 0x49, 0x2c,
	0xc9, 0x2c, 0x2e, 0xc9, 0x4c, 0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x28, 0x85, 0x70, 0x71, 0x38, 0xe7, 0xe7, 0x94, 0xe6, 0xe6, 0x05, 0x38, 0x09, 0x71, 0x71, 0x31,
	0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x0b, 0x71, 0x73, 0x31, 0xe7, 0xa5, 0x94, 0x49,
	0x30, 0x81, 0x39, 0xfc, 0x5c, 0xec, 0x79, 0xa5, 0xb9, 0x49, 0xa9, 0x45, 0xc5, 0x12, 0xcc, 0x0a,
	0xcc, 0x40, 0x01, 0x5e, 0x2e, 0xd6, 0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x16, 0xa0, 0x3c, 0x0f,
	0x48, 0xbe, 0x28, 0xb5, 0x20, 0x35, 0xb1, 0xa4, 0x58, 0x82, 0x15, 0x24, 0xaf, 0x14, 0xce, 0xc5,
	0x1e, 0x92, 0x98, 0x94, 0x93, 0x8a, 0x66, 0x28, 0x90, 0x0d, 0x54, 0x02, 0x31, 0x13, 0x68, 0x44,
	0x72, 0x7e, 0x69, 0x5e, 0x09, 0xd0, 0x44, 0x10, 0x57, 0x95, 0x8b, 0x3d, 0x19, 0xec, 0x8e, 0x62,
	0xa0, 0x99, 0xcc, 0x1a, 0xdc, 0x46, 0x22, 0x7a, 0x48, 0xee, 0x86, 0x39, 0x11, 0x10, 0x00, 0x00,
	0xff, 0xff, 0xe6, 0x97, 0x42, 0x05, 0xcd, 0x00, 0x00, 0x00,
}
