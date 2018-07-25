// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dataRequest.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DbRequest struct {
	// Token will be used to identify the server connection info
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// Or the server connection info was given in the DbRequest
	ServerInfo *ServerInfo `protobuf:"bytes,2,opt,name=serverInfo,proto3" json:"serverInfo,omitempty"`
	// The script name to locate a DB script for execution
	Script string `protobuf:"bytes,3,opt,name=script,proto3" json:"script,omitempty"`
	// The plain text sql statement
	SqlStatement         string   `protobuf:"bytes,4,opt,name=sqlStatement,proto3" json:"sqlStatement,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DbRequest) Reset()         { *m = DbRequest{} }
func (m *DbRequest) String() string { return proto.CompactTextString(m) }
func (*DbRequest) ProtoMessage()    {}
func (*DbRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dataRequest_a07dfe2c990dc3c8, []int{0}
}
func (m *DbRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DbRequest.Unmarshal(m, b)
}
func (m *DbRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DbRequest.Marshal(b, m, deterministic)
}
func (dst *DbRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DbRequest.Merge(dst, src)
}
func (m *DbRequest) XXX_Size() int {
	return xxx_messageInfo_DbRequest.Size(m)
}
func (m *DbRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DbRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DbRequest proto.InternalMessageInfo

func (m *DbRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *DbRequest) GetServerInfo() *ServerInfo {
	if m != nil {
		return m.ServerInfo
	}
	return nil
}

func (m *DbRequest) GetScript() string {
	if m != nil {
		return m.Script
	}
	return ""
}

func (m *DbRequest) GetSqlStatement() string {
	if m != nil {
		return m.SqlStatement
	}
	return ""
}

type ServerInfo struct {
	Server               string   `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Port                 int32    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Database             string   `protobuf:"bytes,3,opt,name=database,proto3" json:"database,omitempty"`
	UserId               string   `protobuf:"bytes,4,opt,name=userId,proto3" json:"userId,omitempty"`
	Password             string   `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerInfo) Reset()         { *m = ServerInfo{} }
func (m *ServerInfo) String() string { return proto.CompactTextString(m) }
func (*ServerInfo) ProtoMessage()    {}
func (*ServerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_dataRequest_a07dfe2c990dc3c8, []int{1}
}
func (m *ServerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerInfo.Unmarshal(m, b)
}
func (m *ServerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerInfo.Marshal(b, m, deterministic)
}
func (dst *ServerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerInfo.Merge(dst, src)
}
func (m *ServerInfo) XXX_Size() int {
	return xxx_messageInfo_ServerInfo.Size(m)
}
func (m *ServerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ServerInfo proto.InternalMessageInfo

func (m *ServerInfo) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

func (m *ServerInfo) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *ServerInfo) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

func (m *ServerInfo) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *ServerInfo) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type DbResponse struct {
	Succeeded            bool     `protobuf:"varint,1,opt,name=succeeded,proto3" json:"succeeded,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	ScalarValueType      string   `protobuf:"bytes,3,opt,name=scalarValueType,proto3" json:"scalarValueType,omitempty"`
	ScalarValue          *DBValue `protobuf:"bytes,4,opt,name=scalarValue,proto3" json:"scalarValue,omitempty"`
	RowEffected          int32    `protobuf:"varint,5,opt,name=rowEffected,proto3" json:"rowEffected,omitempty"`
	Dataset              *DataSet `protobuf:"bytes,6,opt,name=dataset,proto3" json:"dataset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DbResponse) Reset()         { *m = DbResponse{} }
func (m *DbResponse) String() string { return proto.CompactTextString(m) }
func (*DbResponse) ProtoMessage()    {}
func (*DbResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dataRequest_a07dfe2c990dc3c8, []int{2}
}
func (m *DbResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DbResponse.Unmarshal(m, b)
}
func (m *DbResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DbResponse.Marshal(b, m, deterministic)
}
func (dst *DbResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DbResponse.Merge(dst, src)
}
func (m *DbResponse) XXX_Size() int {
	return xxx_messageInfo_DbResponse.Size(m)
}
func (m *DbResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DbResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DbResponse proto.InternalMessageInfo

func (m *DbResponse) GetSucceeded() bool {
	if m != nil {
		return m.Succeeded
	}
	return false
}

func (m *DbResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DbResponse) GetScalarValueType() string {
	if m != nil {
		return m.ScalarValueType
	}
	return ""
}

func (m *DbResponse) GetScalarValue() *DBValue {
	if m != nil {
		return m.ScalarValue
	}
	return nil
}

func (m *DbResponse) GetRowEffected() int32 {
	if m != nil {
		return m.RowEffected
	}
	return 0
}

func (m *DbResponse) GetDataset() *DataSet {
	if m != nil {
		return m.Dataset
	}
	return nil
}

type DataSet struct {
	Tables               []*DataTable `protobuf:"bytes,1,rep,name=tables,proto3" json:"tables,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DataSet) Reset()         { *m = DataSet{} }
func (m *DataSet) String() string { return proto.CompactTextString(m) }
func (*DataSet) ProtoMessage()    {}
func (*DataSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_dataRequest_a07dfe2c990dc3c8, []int{3}
}
func (m *DataSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataSet.Unmarshal(m, b)
}
func (m *DataSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataSet.Marshal(b, m, deterministic)
}
func (dst *DataSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataSet.Merge(dst, src)
}
func (m *DataSet) XXX_Size() int {
	return xxx_messageInfo_DataSet.Size(m)
}
func (m *DataSet) XXX_DiscardUnknown() {
	xxx_messageInfo_DataSet.DiscardUnknown(m)
}

var xxx_messageInfo_DataSet proto.InternalMessageInfo

func (m *DataSet) GetTables() []*DataTable {
	if m != nil {
		return m.Tables
	}
	return nil
}

type DataTable struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Column definitions of this table
	Columns []*DataColumn `protobuf:"bytes,2,rep,name=columns,proto3" json:"columns,omitempty"`
	// Rows in this table
	Rows                 []*DataRow `protobuf:"bytes,3,rep,name=rows,proto3" json:"rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *DataTable) Reset()         { *m = DataTable{} }
func (m *DataTable) String() string { return proto.CompactTextString(m) }
func (*DataTable) ProtoMessage()    {}
func (*DataTable) Descriptor() ([]byte, []int) {
	return fileDescriptor_dataRequest_a07dfe2c990dc3c8, []int{4}
}
func (m *DataTable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataTable.Unmarshal(m, b)
}
func (m *DataTable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataTable.Marshal(b, m, deterministic)
}
func (dst *DataTable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataTable.Merge(dst, src)
}
func (m *DataTable) XXX_Size() int {
	return xxx_messageInfo_DataTable.Size(m)
}
func (m *DataTable) XXX_DiscardUnknown() {
	xxx_messageInfo_DataTable.DiscardUnknown(m)
}

var xxx_messageInfo_DataTable proto.InternalMessageInfo

func (m *DataTable) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DataTable) GetColumns() []*DataColumn {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *DataTable) GetRows() []*DataRow {
	if m != nil {
		return m.Rows
	}
	return nil
}

type DataColumn struct {
	// Zero based column index in the DB record set
	Index int32 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	// The name of the column
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The db type of the column
	DbType string `protobuf:"bytes,3,opt,name=dbType,proto3" json:"dbType,omitempty"`
	// The db size
	DbSize int32 `protobuf:"varint,4,opt,name=dbSize,proto3" json:"dbSize,omitempty"`
	// golang type
	Type                 string   `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	Precision            int32    `protobuf:"varint,6,opt,name=precision,proto3" json:"precision,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataColumn) Reset()         { *m = DataColumn{} }
func (m *DataColumn) String() string { return proto.CompactTextString(m) }
func (*DataColumn) ProtoMessage()    {}
func (*DataColumn) Descriptor() ([]byte, []int) {
	return fileDescriptor_dataRequest_a07dfe2c990dc3c8, []int{5}
}
func (m *DataColumn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataColumn.Unmarshal(m, b)
}
func (m *DataColumn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataColumn.Marshal(b, m, deterministic)
}
func (dst *DataColumn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataColumn.Merge(dst, src)
}
func (m *DataColumn) XXX_Size() int {
	return xxx_messageInfo_DataColumn.Size(m)
}
func (m *DataColumn) XXX_DiscardUnknown() {
	xxx_messageInfo_DataColumn.DiscardUnknown(m)
}

var xxx_messageInfo_DataColumn proto.InternalMessageInfo

func (m *DataColumn) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *DataColumn) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DataColumn) GetDbType() string {
	if m != nil {
		return m.DbType
	}
	return ""
}

func (m *DataColumn) GetDbSize() int32 {
	if m != nil {
		return m.DbSize
	}
	return 0
}

func (m *DataColumn) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *DataColumn) GetPrecision() int32 {
	if m != nil {
		return m.Precision
	}
	return 0
}

type DataRow struct {
	Values               []*DBValue `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *DataRow) Reset()         { *m = DataRow{} }
func (m *DataRow) String() string { return proto.CompactTextString(m) }
func (*DataRow) ProtoMessage()    {}
func (*DataRow) Descriptor() ([]byte, []int) {
	return fileDescriptor_dataRequest_a07dfe2c990dc3c8, []int{6}
}
func (m *DataRow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataRow.Unmarshal(m, b)
}
func (m *DataRow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataRow.Marshal(b, m, deterministic)
}
func (dst *DataRow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataRow.Merge(dst, src)
}
func (m *DataRow) XXX_Size() int {
	return xxx_messageInfo_DataRow.Size(m)
}
func (m *DataRow) XXX_DiscardUnknown() {
	xxx_messageInfo_DataRow.DiscardUnknown(m)
}

var xxx_messageInfo_DataRow proto.InternalMessageInfo

func (m *DataRow) GetValues() []*DBValue {
	if m != nil {
		return m.Values
	}
	return nil
}

type DBValue struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DBValue) Reset()         { *m = DBValue{} }
func (m *DBValue) String() string { return proto.CompactTextString(m) }
func (*DBValue) ProtoMessage()    {}
func (*DBValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_dataRequest_a07dfe2c990dc3c8, []int{7}
}
func (m *DBValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DBValue.Unmarshal(m, b)
}
func (m *DBValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DBValue.Marshal(b, m, deterministic)
}
func (dst *DBValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DBValue.Merge(dst, src)
}
func (m *DBValue) XXX_Size() int {
	return xxx_messageInfo_DBValue.Size(m)
}
func (m *DBValue) XXX_DiscardUnknown() {
	xxx_messageInfo_DBValue.DiscardUnknown(m)
}

var xxx_messageInfo_DBValue proto.InternalMessageInfo

func (m *DBValue) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*DbRequest)(nil), "proto.DbRequest")
	proto.RegisterType((*ServerInfo)(nil), "proto.ServerInfo")
	proto.RegisterType((*DbResponse)(nil), "proto.DbResponse")
	proto.RegisterType((*DataSet)(nil), "proto.DataSet")
	proto.RegisterType((*DataTable)(nil), "proto.DataTable")
	proto.RegisterType((*DataColumn)(nil), "proto.DataColumn")
	proto.RegisterType((*DataRow)(nil), "proto.DataRow")
	proto.RegisterType((*DBValue)(nil), "proto.DBValue")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RemoteDBServiceClient is the client API for RemoteDBService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RemoteDBServiceClient interface {
	ExecuteNoneQuery(ctx context.Context, in *DbRequest, opts ...grpc.CallOption) (*DbResponse, error)
	ExecuteScalar(ctx context.Context, in *DbRequest, opts ...grpc.CallOption) (*DbResponse, error)
	ExecuteDataSet(ctx context.Context, in *DbRequest, opts ...grpc.CallOption) (*DbResponse, error)
}

type remoteDBServiceClient struct {
	cc *grpc.ClientConn
}

func NewRemoteDBServiceClient(cc *grpc.ClientConn) RemoteDBServiceClient {
	return &remoteDBServiceClient{cc}
}

func (c *remoteDBServiceClient) ExecuteNoneQuery(ctx context.Context, in *DbRequest, opts ...grpc.CallOption) (*DbResponse, error) {
	out := new(DbResponse)
	err := c.cc.Invoke(ctx, "/proto.RemoteDBService/ExecuteNoneQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteDBServiceClient) ExecuteScalar(ctx context.Context, in *DbRequest, opts ...grpc.CallOption) (*DbResponse, error) {
	out := new(DbResponse)
	err := c.cc.Invoke(ctx, "/proto.RemoteDBService/ExecuteScalar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteDBServiceClient) ExecuteDataSet(ctx context.Context, in *DbRequest, opts ...grpc.CallOption) (*DbResponse, error) {
	out := new(DbResponse)
	err := c.cc.Invoke(ctx, "/proto.RemoteDBService/ExecuteDataSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RemoteDBServiceServer is the server API for RemoteDBService service.
type RemoteDBServiceServer interface {
	ExecuteNoneQuery(context.Context, *DbRequest) (*DbResponse, error)
	ExecuteScalar(context.Context, *DbRequest) (*DbResponse, error)
	ExecuteDataSet(context.Context, *DbRequest) (*DbResponse, error)
}

func RegisterRemoteDBServiceServer(s *grpc.Server, srv RemoteDBServiceServer) {
	s.RegisterService(&_RemoteDBService_serviceDesc, srv)
}

func _RemoteDBService_ExecuteNoneQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DbRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteDBServiceServer).ExecuteNoneQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RemoteDBService/ExecuteNoneQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteDBServiceServer).ExecuteNoneQuery(ctx, req.(*DbRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteDBService_ExecuteScalar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DbRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteDBServiceServer).ExecuteScalar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RemoteDBService/ExecuteScalar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteDBServiceServer).ExecuteScalar(ctx, req.(*DbRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteDBService_ExecuteDataSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DbRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteDBServiceServer).ExecuteDataSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RemoteDBService/ExecuteDataSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteDBServiceServer).ExecuteDataSet(ctx, req.(*DbRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RemoteDBService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.RemoteDBService",
	HandlerType: (*RemoteDBServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecuteNoneQuery",
			Handler:    _RemoteDBService_ExecuteNoneQuery_Handler,
		},
		{
			MethodName: "ExecuteScalar",
			Handler:    _RemoteDBService_ExecuteScalar_Handler,
		},
		{
			MethodName: "ExecuteDataSet",
			Handler:    _RemoteDBService_ExecuteDataSet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dataRequest.proto",
}

func init() { proto.RegisterFile("dataRequest.proto", fileDescriptor_dataRequest_a07dfe2c990dc3c8) }

var fileDescriptor_dataRequest_a07dfe2c990dc3c8 = []byte{
	// 581 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0xc5, 0x49, 0xec, 0x34, 0x37, 0x7d, 0x65, 0x54, 0x55, 0x56, 0x85, 0x44, 0xe4, 0x05, 0x58,
	0x42, 0x8a, 0x68, 0x2a, 0x81, 0xd8, 0x86, 0x74, 0xd1, 0x0d, 0x2a, 0xe3, 0x8a, 0xfd, 0xc4, 0xbe,
	0x29, 0x16, 0x89, 0xc7, 0x9d, 0x19, 0x37, 0x2d, 0x5f, 0xc0, 0x92, 0x25, 0x3f, 0xc3, 0x1f, 0xf1,
	0x11, 0xd5, 0x3c, 0xec, 0x3a, 0x59, 0x75, 0xe5, 0x39, 0xe7, 0xce, 0xb9, 0x73, 0xee, 0xc3, 0x30,
	0xca, 0x98, 0x62, 0x14, 0xef, 0x2a, 0x94, 0x6a, 0x52, 0x0a, 0xae, 0x38, 0xf1, 0xcd, 0x27, 0xfa,
	0xe3, 0xc1, 0x60, 0xbe, 0x70, 0x21, 0x72, 0x02, 0xbe, 0xe2, 0x3f, 0xb1, 0x08, 0xbd, 0xb1, 0x17,
	0x0f, 0xa8, 0x05, 0xe4, 0x1c, 0x40, 0xa2, 0xb8, 0x47, 0x71, 0x55, 0x2c, 0x79, 0xd8, 0x19, 0x7b,
	0xf1, 0x70, 0x3a, 0xb2, 0x69, 0x26, 0x49, 0x13, 0xa0, 0xad, 0x4b, 0xe4, 0x14, 0x02, 0x99, 0x8a,
	0xbc, 0x54, 0x61, 0xd7, 0x64, 0x72, 0x88, 0x44, 0xb0, 0x2f, 0xef, 0x56, 0x89, 0x62, 0x0a, 0xd7,
	0x58, 0xa8, 0xb0, 0x67, 0xa2, 0x5b, 0x5c, 0xf4, 0xdb, 0x03, 0x48, 0xb6, 0x53, 0x19, 0xe4, 0x4c,
	0x39, 0x44, 0x08, 0xf4, 0x4a, 0x2e, 0x94, 0xf1, 0xe3, 0x53, 0x73, 0x26, 0x67, 0xb0, 0xa7, 0x2b,
	0x5d, 0x30, 0x89, 0xee, 0xe1, 0x06, 0xeb, 0x3c, 0x95, 0x44, 0x71, 0x95, 0xb9, 0x47, 0x1d, 0xd2,
	0x9a, 0x92, 0x49, 0xb9, 0xe1, 0x22, 0x0b, 0x7d, 0xab, 0xa9, 0x71, 0xf4, 0xdf, 0x03, 0xd0, 0xdd,
	0x91, 0x25, 0x2f, 0x24, 0x92, 0xd7, 0x30, 0x90, 0x55, 0x9a, 0x22, 0x66, 0x98, 0x19, 0x37, 0x7b,
	0xf4, 0x99, 0x20, 0x21, 0xf4, 0xd7, 0x28, 0x25, 0xbb, 0x45, 0xe3, 0x69, 0x40, 0x6b, 0x48, 0x62,
	0x38, 0x92, 0x29, 0x5b, 0x31, 0xf1, 0x9d, 0xad, 0x2a, 0xbc, 0x79, 0x2c, 0x6b, 0x77, 0xbb, 0x34,
	0xf9, 0x00, 0xc3, 0x16, 0x65, 0x9c, 0x0e, 0xa7, 0x87, 0xae, 0xd7, 0xf3, 0x99, 0x61, 0x69, 0xfb,
	0x0a, 0x19, 0xc3, 0x50, 0xf0, 0xcd, 0xe5, 0x72, 0x89, 0xa9, 0x42, 0x5b, 0x81, 0x4f, 0xdb, 0x14,
	0x89, 0xa1, 0xaf, 0x9b, 0x20, 0x51, 0x85, 0xc1, 0x76, 0x3e, 0xa6, 0x58, 0x82, 0x8a, 0xd6, 0xe1,
	0xe8, 0x02, 0xfa, 0x8e, 0x23, 0x31, 0x04, 0x8a, 0x2d, 0x56, 0x28, 0x43, 0x6f, 0xdc, 0x8d, 0x87,
	0xd3, 0xe3, 0x96, 0xe6, 0x46, 0x07, 0xa8, 0x8b, 0x47, 0x25, 0x0c, 0x1a, 0x52, 0x0f, 0xa5, 0x60,
	0x6b, 0x74, 0xa3, 0x32, 0x67, 0xf2, 0x1e, 0xfa, 0x29, 0x5f, 0x55, 0xeb, 0x42, 0x86, 0x1d, 0x93,
	0x6b, 0xd4, 0xca, 0xf5, 0xc5, 0x44, 0x68, 0x7d, 0x83, 0x44, 0xd0, 0x13, 0x7c, 0x23, 0xc3, 0xae,
	0xb9, 0xd9, 0x76, 0x4a, 0xf9, 0x86, 0x9a, 0x58, 0xf4, 0x57, 0x4f, 0xa5, 0xd1, 0xea, 0xa5, 0xcd,
	0x8b, 0x0c, 0x1f, 0xcc, 0xa3, 0x3e, 0xb5, 0xa0, 0x71, 0xd2, 0x69, 0x39, 0x39, 0x85, 0x20, 0x5b,
	0xb4, 0xda, 0xef, 0x90, 0xe5, 0x93, 0xfc, 0x97, 0x6d, 0xb8, 0x4f, 0x1d, 0xd2, 0x39, 0x94, 0xbe,
	0x6d, 0xd7, 0xc2, 0x9c, 0xf5, 0x0e, 0x94, 0x02, 0xd3, 0x5c, 0xe6, 0xbc, 0x30, 0xfd, 0xf4, 0xe9,
	0x33, 0x11, 0x9d, 0xdb, 0x0e, 0x52, 0xbe, 0x21, 0x6f, 0x21, 0xb8, 0xd7, 0x13, 0xaa, 0x3b, 0xb8,
	0x3b, 0x45, 0x17, 0x8d, 0xde, 0x40, 0xdf, 0x51, 0xba, 0x12, 0x43, 0x9a, 0x4a, 0xf6, 0xa9, 0x05,
	0xd3, 0x7f, 0x1e, 0x1c, 0x51, 0x5c, 0x73, 0x85, 0xf3, 0x99, 0xfe, 0x2f, 0xf2, 0x14, 0xc9, 0x67,
	0x38, 0xbe, 0x7c, 0xc0, 0xb4, 0x52, 0xf8, 0x95, 0x17, 0xf8, 0xad, 0x42, 0xf1, 0x48, 0x9a, 0x11,
	0xd5, 0xbf, 0xf3, 0xd9, 0xa8, 0xc5, 0xd8, 0x15, 0x8e, 0x5e, 0x91, 0x8f, 0x70, 0xe0, 0xa4, 0x89,
	0x59, 0xa3, 0x97, 0xea, 0x3e, 0xc1, 0xa1, 0xd3, 0xd5, 0x3b, 0xf2, 0x32, 0xe1, 0xec, 0x1d, 0x1c,
	0x2c, 0x73, 0xf9, 0x63, 0x72, 0x2b, 0xca, 0x74, 0x22, 0x32, 0x36, 0x3b, 0xd9, 0xa9, 0xe6, 0x5a,
	0x6b, 0xae, 0xbd, 0x45, 0x60, 0xc4, 0x17, 0x4f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8d, 0xb3, 0x07,
	0xf6, 0xae, 0x04, 0x00, 0x00,
}
