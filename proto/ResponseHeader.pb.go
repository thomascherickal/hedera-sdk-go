// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/ResponseHeader.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Every query receives a response containing the QueryResponseHeader. Either or both of the cost and stateProof fields may be blank, if the responseType didn't ask for the cost or stateProof.
type ResponseHeader struct {
	NodeTransactionPrecheckCode ResponseCodeEnum `protobuf:"varint,1,opt,name=nodeTransactionPrecheckCode,proto3,enum=proto.ResponseCodeEnum" json:"nodeTransactionPrecheckCode,omitempty"`
	ResponseType                ResponseType     `protobuf:"varint,2,opt,name=responseType,proto3,enum=proto.ResponseType" json:"responseType,omitempty"`
	Cost                        uint64           `protobuf:"varint,3,opt,name=cost,proto3" json:"cost,omitempty"`
	StateProof                  []byte           `protobuf:"bytes,4,opt,name=stateProof,proto3" json:"stateProof,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}         `json:"-"`
	XXX_unrecognized            []byte           `json:"-"`
	XXX_sizecache               int32            `json:"-"`
}

func (m *ResponseHeader) Reset()         { *m = ResponseHeader{} }
func (m *ResponseHeader) String() string { return proto.CompactTextString(m) }
func (*ResponseHeader) ProtoMessage()    {}
func (*ResponseHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_52a2ae021379db8d, []int{0}
}

func (m *ResponseHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseHeader.Unmarshal(m, b)
}
func (m *ResponseHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseHeader.Marshal(b, m, deterministic)
}
func (m *ResponseHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseHeader.Merge(m, src)
}
func (m *ResponseHeader) XXX_Size() int {
	return xxx_messageInfo_ResponseHeader.Size(m)
}
func (m *ResponseHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseHeader proto.InternalMessageInfo

func (m *ResponseHeader) GetNodeTransactionPrecheckCode() ResponseCodeEnum {
	if m != nil {
		return m.NodeTransactionPrecheckCode
	}
	return ResponseCodeEnum_OK
}

func (m *ResponseHeader) GetResponseType() ResponseType {
	if m != nil {
		return m.ResponseType
	}
	return ResponseType_ANSWER_ONLY
}

func (m *ResponseHeader) GetCost() uint64 {
	if m != nil {
		return m.Cost
	}
	return 0
}

func (m *ResponseHeader) GetStateProof() []byte {
	if m != nil {
		return m.StateProof
	}
	return nil
}

func init() {
	proto.RegisterType((*ResponseHeader)(nil), "proto.ResponseHeader")
}

func init() { proto.RegisterFile("proto/ResponseHeader.proto", fileDescriptor_52a2ae021379db8d) }

var fileDescriptor_52a2ae021379db8d = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8d, 0x31, 0x4b, 0xc4, 0x30,
	0x18, 0x86, 0x89, 0x56, 0x87, 0x70, 0xdc, 0x10, 0x87, 0x0b, 0x15, 0xa4, 0x38, 0x75, 0xb9, 0x14,
	0x74, 0x70, 0x57, 0x84, 0x1b, 0x6b, 0xb8, 0x45, 0xb7, 0x5c, 0xf2, 0xd9, 0x1c, 0x47, 0xf3, 0x95,
	0x24, 0x1d, 0xee, 0x9f, 0xfa, 0x73, 0xa4, 0x89, 0x48, 0xeb, 0xe0, 0xf4, 0x85, 0xf7, 0x7d, 0x9f,
	0x27, 0xb4, 0x1c, 0x3c, 0x46, 0x6c, 0x24, 0x84, 0x01, 0x5d, 0x80, 0x1d, 0x28, 0x03, 0x5e, 0xa4,
	0x90, 0x5d, 0xa5, 0x53, 0x6e, 0xf2, 0xe4, 0x6d, 0x04, 0x7f, 0x9e, 0xf7, 0x25, 0x5f, 0xb2, 0x2f,
	0x68, 0x20, 0x37, 0xf7, 0x5f, 0x84, 0xae, 0x97, 0x4a, 0xf6, 0x4e, 0x6f, 0x1d, 0x1a, 0xd8, 0x7b,
	0xe5, 0x82, 0xd2, 0xf1, 0x88, 0xae, 0xf5, 0xa0, 0x2d, 0xe8, 0xd3, 0xc4, 0x71, 0x52, 0x91, 0x7a,
	0xfd, 0xb0, 0xc9, 0xbc, 0x98, 0x2b, 0x5f, 0xdd, 0xd8, 0xcb, 0xff, 0x58, 0xf6, 0x44, 0x57, 0xfe,
	0x07, 0xd8, 0x9f, 0x07, 0xe0, 0x17, 0xc9, 0x75, 0xf3, 0xc7, 0x35, 0x55, 0x72, 0x31, 0x64, 0x8c,
	0x16, 0x1a, 0x43, 0xe4, 0x97, 0x15, 0xa9, 0x0b, 0x99, 0xde, 0xec, 0x8e, 0xd2, 0x10, 0x55, 0x84,
	0xd6, 0x23, 0x7e, 0xf2, 0xa2, 0x22, 0xf5, 0x4a, 0xce, 0x92, 0xe7, 0x1d, 0x2d, 0x35, 0xf6, 0xc2,
	0x82, 0x01, 0xaf, 0x84, 0x55, 0xc1, 0x76, 0x5e, 0x0d, 0x36, 0x7f, 0xd6, 0x92, 0x8f, 0xba, 0x3b,
	0x46, 0x3b, 0x1e, 0x84, 0xc6, 0xbe, 0xf9, 0x6d, 0x9b, 0x3c, 0xdf, 0x06, 0x73, 0xda, 0x76, 0xd8,
	0xa4, 0xed, 0xe1, 0x3a, 0x9d, 0xc7, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3b, 0xaf, 0x43, 0xf8,
	0x83, 0x01, 0x00, 0x00,
}
