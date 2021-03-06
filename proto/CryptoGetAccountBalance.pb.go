// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/CryptoGetAccountBalance.proto

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

// Get the balance of a cryptocurrency account. This returns only the balance, so it is a smaller and faster reply than CryptoGetInfo, which returns the balance plus additional information.
type CryptoGetAccountBalanceQuery struct {
	Header               *QueryHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	AccountID            *AccountID   `protobuf:"bytes,2,opt,name=accountID,proto3" json:"accountID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CryptoGetAccountBalanceQuery) Reset()         { *m = CryptoGetAccountBalanceQuery{} }
func (m *CryptoGetAccountBalanceQuery) String() string { return proto.CompactTextString(m) }
func (*CryptoGetAccountBalanceQuery) ProtoMessage()    {}
func (*CryptoGetAccountBalanceQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_195622b96c4005ba, []int{0}
}

func (m *CryptoGetAccountBalanceQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CryptoGetAccountBalanceQuery.Unmarshal(m, b)
}
func (m *CryptoGetAccountBalanceQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CryptoGetAccountBalanceQuery.Marshal(b, m, deterministic)
}
func (m *CryptoGetAccountBalanceQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CryptoGetAccountBalanceQuery.Merge(m, src)
}
func (m *CryptoGetAccountBalanceQuery) XXX_Size() int {
	return xxx_messageInfo_CryptoGetAccountBalanceQuery.Size(m)
}
func (m *CryptoGetAccountBalanceQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_CryptoGetAccountBalanceQuery.DiscardUnknown(m)
}

var xxx_messageInfo_CryptoGetAccountBalanceQuery proto.InternalMessageInfo

func (m *CryptoGetAccountBalanceQuery) GetHeader() *QueryHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CryptoGetAccountBalanceQuery) GetAccountID() *AccountID {
	if m != nil {
		return m.AccountID
	}
	return nil
}

// Response when the client sends the node CryptoGetAccountBalanceQuery
type CryptoGetAccountBalanceResponse struct {
	Header               *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	AccountID            *AccountID      `protobuf:"bytes,2,opt,name=accountID,proto3" json:"accountID,omitempty"`
	Balance              uint64          `protobuf:"varint,3,opt,name=balance,proto3" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CryptoGetAccountBalanceResponse) Reset()         { *m = CryptoGetAccountBalanceResponse{} }
func (m *CryptoGetAccountBalanceResponse) String() string { return proto.CompactTextString(m) }
func (*CryptoGetAccountBalanceResponse) ProtoMessage()    {}
func (*CryptoGetAccountBalanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_195622b96c4005ba, []int{1}
}

func (m *CryptoGetAccountBalanceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CryptoGetAccountBalanceResponse.Unmarshal(m, b)
}
func (m *CryptoGetAccountBalanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CryptoGetAccountBalanceResponse.Marshal(b, m, deterministic)
}
func (m *CryptoGetAccountBalanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CryptoGetAccountBalanceResponse.Merge(m, src)
}
func (m *CryptoGetAccountBalanceResponse) XXX_Size() int {
	return xxx_messageInfo_CryptoGetAccountBalanceResponse.Size(m)
}
func (m *CryptoGetAccountBalanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CryptoGetAccountBalanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CryptoGetAccountBalanceResponse proto.InternalMessageInfo

func (m *CryptoGetAccountBalanceResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CryptoGetAccountBalanceResponse) GetAccountID() *AccountID {
	if m != nil {
		return m.AccountID
	}
	return nil
}

func (m *CryptoGetAccountBalanceResponse) GetBalance() uint64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func init() {
	proto.RegisterType((*CryptoGetAccountBalanceQuery)(nil), "proto.CryptoGetAccountBalanceQuery")
	proto.RegisterType((*CryptoGetAccountBalanceResponse)(nil), "proto.CryptoGetAccountBalanceResponse")
}

func init() {
	proto.RegisterFile("proto/CryptoGetAccountBalance.proto", fileDescriptor_195622b96c4005ba)
}

var fileDescriptor_195622b96c4005ba = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x8f, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0xc6, 0x89, 0x7f, 0x56, 0x8c, 0x17, 0x09, 0xa8, 0xa5, 0x08, 0x2e, 0xeb, 0xa5, 0x08, 0x4d,
	0x41, 0x9f, 0x60, 0xab, 0xe0, 0x7a, 0xd3, 0xe2, 0xc9, 0x5b, 0x9a, 0x0e, 0xcd, 0xa2, 0xdb, 0x84,
	0x24, 0x3d, 0xd4, 0x47, 0xf1, 0x69, 0x85, 0x4c, 0xaa, 0xab, 0xd0, 0xcb, 0x9e, 0x86, 0xce, 0xef,
	0xd7, 0xef, 0x9b, 0xd0, 0x6b, 0x63, 0xb5, 0xd7, 0xc5, 0xbd, 0x1d, 0x8c, 0xd7, 0x8f, 0xe0, 0x97,
	0x52, 0xea, 0xbe, 0xf3, 0xa5, 0xf8, 0x10, 0x9d, 0x04, 0x1e, 0x28, 0x3b, 0x0c, 0x23, 0x3d, 0x47,
	0xb7, 0x14, 0x6e, 0x2d, 0x5f, 0x07, 0x03, 0x0e, 0x71, 0x7a, 0x81, 0xfb, 0x97, 0x1e, 0xec, 0xb0,
	0x02, 0xd1, 0x80, 0x8d, 0x20, 0x45, 0x50, 0x81, 0x33, 0xba, 0x73, 0xb0, 0xcd, 0x16, 0x9f, 0xf4,
	0x72, 0xa2, 0x34, 0xe4, 0xb0, 0x1b, 0x3a, 0x53, 0xc1, 0x4f, 0xc8, 0x9c, 0x64, 0x27, 0xb7, 0x0c,
	0xff, 0xe3, 0x5b, 0x2d, 0x55, 0x34, 0x18, 0xa7, 0xc7, 0x02, 0x23, 0x9e, 0x1e, 0x92, 0xbd, 0xa0,
	0x9f, 0x46, 0x7d, 0x39, 0xee, 0xab, 0x5f, 0x65, 0xf1, 0x45, 0xe8, 0xd5, 0x44, 0xf9, 0x78, 0x2b,
	0xcb, 0xff, 0xf5, 0x9f, 0xc5, 0xc0, 0xbf, 0x8f, 0xd9, 0xf5, 0x04, 0x96, 0xd0, 0xa3, 0x1a, 0x1b,
	0x93, 0xfd, 0x39, 0xc9, 0x0e, 0xaa, 0xf1, 0xb3, 0x5c, 0xd1, 0x54, 0xea, 0x0d, 0x57, 0xd0, 0x80,
	0x15, 0x5c, 0x09, 0xa7, 0x5a, 0x2b, 0x8c, 0xc2, 0xb0, 0x67, 0xf2, 0x96, 0xb5, 0x6b, 0xaf, 0xfa,
	0x9a, 0x4b, 0xbd, 0x29, 0x7e, 0x68, 0x81, 0x7a, 0xee, 0x9a, 0xf7, 0xbc, 0xd5, 0x45, 0x70, 0xeb,
	0x59, 0x18, 0x77, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xff, 0xdc, 0x3f, 0xb0, 0xe4, 0x01, 0x00,
	0x00,
}
