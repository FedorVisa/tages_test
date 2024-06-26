// Code generated by protoc-gen-go. DO NOT EDIT.
// source: structure.proto

package protobuf

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type FileInfo struct {
	Filename             string   `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Fileformat           string   `protobuf:"bytes,2,opt,name=fileformat,proto3" json:"fileformat,omitempty"`
	CreatedAt            string   `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileInfo) Reset()         { *m = FileInfo{} }
func (m *FileInfo) String() string { return proto.CompactTextString(m) }
func (*FileInfo) ProtoMessage()    {}
func (*FileInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f25f331c538a6a0, []int{0}
}

func (m *FileInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileInfo.Unmarshal(m, b)
}
func (m *FileInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileInfo.Marshal(b, m, deterministic)
}
func (m *FileInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileInfo.Merge(m, src)
}
func (m *FileInfo) XXX_Size() int {
	return xxx_messageInfo_FileInfo.Size(m)
}
func (m *FileInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FileInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FileInfo proto.InternalMessageInfo

func (m *FileInfo) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *FileInfo) GetFileformat() string {
	if m != nil {
		return m.Fileformat
	}
	return ""
}

func (m *FileInfo) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *FileInfo) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type Image struct {
	Info                 *FileInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Data                 []byte    `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f25f331c538a6a0, []int{1}
}

func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (m *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(m, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetInfo() *FileInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *Image) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type UploadRequest struct {
	Image                *Image   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadRequest) Reset()         { *m = UploadRequest{} }
func (m *UploadRequest) String() string { return proto.CompactTextString(m) }
func (*UploadRequest) ProtoMessage()    {}
func (*UploadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f25f331c538a6a0, []int{2}
}

func (m *UploadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadRequest.Unmarshal(m, b)
}
func (m *UploadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadRequest.Marshal(b, m, deterministic)
}
func (m *UploadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadRequest.Merge(m, src)
}
func (m *UploadRequest) XXX_Size() int {
	return xxx_messageInfo_UploadRequest.Size(m)
}
func (m *UploadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadRequest proto.InternalMessageInfo

func (m *UploadRequest) GetImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

type UploadResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadResponse) Reset()         { *m = UploadResponse{} }
func (m *UploadResponse) String() string { return proto.CompactTextString(m) }
func (*UploadResponse) ProtoMessage()    {}
func (*UploadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f25f331c538a6a0, []int{3}
}

func (m *UploadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadResponse.Unmarshal(m, b)
}
func (m *UploadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadResponse.Marshal(b, m, deterministic)
}
func (m *UploadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadResponse.Merge(m, src)
}
func (m *UploadResponse) XXX_Size() int {
	return xxx_messageInfo_UploadResponse.Size(m)
}
func (m *UploadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadResponse proto.InternalMessageInfo

func (m *UploadResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *UploadResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ListFilesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListFilesRequest) Reset()         { *m = ListFilesRequest{} }
func (m *ListFilesRequest) String() string { return proto.CompactTextString(m) }
func (*ListFilesRequest) ProtoMessage()    {}
func (*ListFilesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f25f331c538a6a0, []int{4}
}

func (m *ListFilesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListFilesRequest.Unmarshal(m, b)
}
func (m *ListFilesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListFilesRequest.Marshal(b, m, deterministic)
}
func (m *ListFilesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFilesRequest.Merge(m, src)
}
func (m *ListFilesRequest) XXX_Size() int {
	return xxx_messageInfo_ListFilesRequest.Size(m)
}
func (m *ListFilesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFilesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListFilesRequest proto.InternalMessageInfo

type ListFilesResponse struct {
	Files                []string `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListFilesResponse) Reset()         { *m = ListFilesResponse{} }
func (m *ListFilesResponse) String() string { return proto.CompactTextString(m) }
func (*ListFilesResponse) ProtoMessage()    {}
func (*ListFilesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f25f331c538a6a0, []int{5}
}

func (m *ListFilesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListFilesResponse.Unmarshal(m, b)
}
func (m *ListFilesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListFilesResponse.Marshal(b, m, deterministic)
}
func (m *ListFilesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFilesResponse.Merge(m, src)
}
func (m *ListFilesResponse) XXX_Size() int {
	return xxx_messageInfo_ListFilesResponse.Size(m)
}
func (m *ListFilesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFilesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListFilesResponse proto.InternalMessageInfo

func (m *ListFilesResponse) GetFiles() []string {
	if m != nil {
		return m.Files
	}
	return nil
}

type DownloadRequest struct {
	Filename             string   `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	FileFormat           string   `protobuf:"bytes,2,opt,name=file_format,json=fileFormat,proto3" json:"file_format,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadRequest) Reset()         { *m = DownloadRequest{} }
func (m *DownloadRequest) String() string { return proto.CompactTextString(m) }
func (*DownloadRequest) ProtoMessage()    {}
func (*DownloadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f25f331c538a6a0, []int{6}
}

func (m *DownloadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadRequest.Unmarshal(m, b)
}
func (m *DownloadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadRequest.Marshal(b, m, deterministic)
}
func (m *DownloadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadRequest.Merge(m, src)
}
func (m *DownloadRequest) XXX_Size() int {
	return xxx_messageInfo_DownloadRequest.Size(m)
}
func (m *DownloadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadRequest proto.InternalMessageInfo

func (m *DownloadRequest) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *DownloadRequest) GetFileFormat() string {
	if m != nil {
		return m.FileFormat
	}
	return ""
}

type DownloadResponse struct {
	Image                *Image   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadResponse) Reset()         { *m = DownloadResponse{} }
func (m *DownloadResponse) String() string { return proto.CompactTextString(m) }
func (*DownloadResponse) ProtoMessage()    {}
func (*DownloadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f25f331c538a6a0, []int{7}
}

func (m *DownloadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadResponse.Unmarshal(m, b)
}
func (m *DownloadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadResponse.Marshal(b, m, deterministic)
}
func (m *DownloadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadResponse.Merge(m, src)
}
func (m *DownloadResponse) XXX_Size() int {
	return xxx_messageInfo_DownloadResponse.Size(m)
}
func (m *DownloadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadResponse proto.InternalMessageInfo

func (m *DownloadResponse) GetImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

func init() {
	proto.RegisterType((*FileInfo)(nil), "protobuf.FileInfo")
	proto.RegisterType((*Image)(nil), "protobuf.Image")
	proto.RegisterType((*UploadRequest)(nil), "protobuf.UploadRequest")
	proto.RegisterType((*UploadResponse)(nil), "protobuf.UploadResponse")
	proto.RegisterType((*ListFilesRequest)(nil), "protobuf.ListFilesRequest")
	proto.RegisterType((*ListFilesResponse)(nil), "protobuf.ListFilesResponse")
	proto.RegisterType((*DownloadRequest)(nil), "protobuf.DownloadRequest")
	proto.RegisterType((*DownloadResponse)(nil), "protobuf.DownloadResponse")
}

func init() {
	proto.RegisterFile("structure.proto", fileDescriptor_1f25f331c538a6a0)
}

var fileDescriptor_1f25f331c538a6a0 = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xdf, 0x4a, 0xeb, 0x40,
	0x10, 0xc6, 0x4f, 0x4e, 0xdb, 0x73, 0x92, 0x69, 0xb5, 0x75, 0x10, 0x8c, 0x11, 0xff, 0xb0, 0xa0,
	0xd4, 0x9b, 0x22, 0x15, 0x04, 0x2f, 0x4b, 0x4b, 0xa1, 0x20, 0x5e, 0x44, 0xbc, 0x2e, 0xdb, 0x64,
	0x23, 0x81, 0x26, 0x1b, 0xb3, 0x1b, 0x7d, 0x01, 0x5f, 0xd5, 0xf7, 0x90, 0xdd, 0x64, 0x9b, 0xda,
	0x16, 0xf1, 0x2a, 0x3b, 0xf3, 0xdb, 0xf9, 0xe6, 0xcb, 0xb7, 0xd0, 0x15, 0x32, 0x2f, 0x02, 0x59,
	0xe4, 0x6c, 0x90, 0xe5, 0x5c, 0x72, 0xb4, 0xf5, 0x67, 0x51, 0x44, 0xe4, 0xc3, 0x02, 0x7b, 0x1a,
	0x2f, 0xd9, 0x2c, 0x8d, 0x38, 0x7a, 0x60, 0x47, 0xf1, 0x92, 0xa5, 0x34, 0x61, 0xae, 0x75, 0x61,
	0xf5, 0x1d, 0x7f, 0x55, 0xe3, 0x19, 0x80, 0x3a, 0x47, 0x3c, 0x4f, 0xa8, 0x74, 0xff, 0x6a, 0xba,
	0xd6, 0xc1, 0x53, 0x80, 0x20, 0x67, 0x54, 0xb2, 0x70, 0x4e, 0xa5, 0xdb, 0xd0, 0xdc, 0xa9, 0x3a,
	0x23, 0x8d, 0x8b, 0x2c, 0x34, 0xb8, 0x59, 0xe2, 0xaa, 0x33, 0x92, 0x64, 0x0c, 0xad, 0x59, 0x42,
	0x5f, 0x18, 0x5e, 0x41, 0x33, 0x4e, 0x23, 0xae, 0xd7, 0xb7, 0x87, 0x38, 0x30, 0x46, 0x07, 0xc6,
	0xa4, 0xaf, 0x39, 0x22, 0x34, 0x43, 0x2a, 0xa9, 0x36, 0xd2, 0xf1, 0xf5, 0x99, 0xdc, 0xc1, 0xde,
	0x73, 0xb6, 0xe4, 0x34, 0xf4, 0xd9, 0x6b, 0xc1, 0x84, 0xc4, 0x4b, 0x68, 0xc5, 0x4a, 0xb5, 0x52,
	0xeb, 0xd6, 0x6a, 0x7a, 0x99, 0x5f, 0x52, 0x32, 0x81, 0x7d, 0x33, 0x27, 0x32, 0x9e, 0x0a, 0x86,
	0x2e, 0xfc, 0x17, 0x45, 0x10, 0x30, 0x21, 0xf4, 0xa8, 0xed, 0x9b, 0x52, 0x91, 0x84, 0x09, 0xa1,
	0x44, 0xcb, 0x0c, 0x4c, 0x49, 0x10, 0x7a, 0x0f, 0xb1, 0x90, 0xca, 0xa7, 0xa8, 0x0c, 0x90, 0x6b,
	0x38, 0x58, 0xeb, 0x55, 0xe2, 0x87, 0xd0, 0x52, 0xb9, 0x29, 0xe9, 0x46, 0xdf, 0xf1, 0xcb, 0x82,
	0x3c, 0x42, 0x77, 0xc2, 0xdf, 0xd3, 0x75, 0xfb, 0x3f, 0x3d, 0xc7, 0x39, 0xb4, 0xd5, 0x79, 0xbe,
	0xfd, 0x1e, 0x53, 0xdd, 0x21, 0xf7, 0xd0, 0xab, 0xf5, 0xaa, 0xcd, 0xbf, 0xcb, 0x63, 0xf8, 0x69,
	0x41, 0x5b, 0x59, 0x7e, 0x62, 0xf9, 0x5b, 0x1c, 0x30, 0x1c, 0x03, 0x94, 0xf9, 0xa8, 0x26, 0x1e,
	0xd5, 0x53, 0xdf, 0xd2, 0xf6, 0xdc, 0x6d, 0x50, 0xee, 0x25, 0x7f, 0xfa, 0x16, 0x4e, 0xc1, 0x59,
	0x45, 0x81, 0x5e, 0x7d, 0x75, 0x33, 0x33, 0xef, 0x64, 0x27, 0x33, 0x4a, 0x38, 0x83, 0x8e, 0xf9,
	0x2f, 0x6d, 0xe7, 0xb8, 0xbe, 0xbe, 0x91, 0x9f, 0xe7, 0xed, 0x42, 0x46, 0xe8, 0xc6, 0x5a, 0xfc,
	0xd3, 0xf8, 0xf6, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x5b, 0xc0, 0x62, 0x18, 0x1f, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FileServiceClient is the client API for FileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FileServiceClient interface {
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (FileService_UploadFileClient, error)
	ListFiles(ctx context.Context, in *ListFilesRequest, opts ...grpc.CallOption) (*ListFilesResponse, error)
	DownloadFile(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (FileService_DownloadFileClient, error)
}

type fileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileServiceClient(cc grpc.ClientConnInterface) FileServiceClient {
	return &fileServiceClient{cc}
}

func (c *fileServiceClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (FileService_UploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FileService_serviceDesc.Streams[0], "/protobuf.FileService/UploadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileServiceUploadFileClient{stream}
	return x, nil
}

type FileService_UploadFileClient interface {
	Send(*UploadRequest) error
	CloseAndRecv() (*UploadResponse, error)
	grpc.ClientStream
}

type fileServiceUploadFileClient struct {
	grpc.ClientStream
}

func (x *fileServiceUploadFileClient) Send(m *UploadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileServiceUploadFileClient) CloseAndRecv() (*UploadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileServiceClient) ListFiles(ctx context.Context, in *ListFilesRequest, opts ...grpc.CallOption) (*ListFilesResponse, error) {
	out := new(ListFilesResponse)
	err := c.cc.Invoke(ctx, "/protobuf.FileService/ListFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) DownloadFile(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (FileService_DownloadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FileService_serviceDesc.Streams[1], "/protobuf.FileService/DownloadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileServiceDownloadFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileService_DownloadFileClient interface {
	Recv() (*DownloadResponse, error)
	grpc.ClientStream
}

type fileServiceDownloadFileClient struct {
	grpc.ClientStream
}

func (x *fileServiceDownloadFileClient) Recv() (*DownloadResponse, error) {
	m := new(DownloadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileServiceServer is the server API for FileService service.
type FileServiceServer interface {
	UploadFile(FileService_UploadFileServer) error
	ListFiles(context.Context, *ListFilesRequest) (*ListFilesResponse, error)
	DownloadFile(*DownloadRequest, FileService_DownloadFileServer) error
}

// UnimplementedFileServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFileServiceServer struct {
}

func (*UnimplementedFileServiceServer) UploadFile(srv FileService_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (*UnimplementedFileServiceServer) ListFiles(ctx context.Context, req *ListFilesRequest) (*ListFilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFiles not implemented")
}
func (*UnimplementedFileServiceServer) DownloadFile(req *DownloadRequest, srv FileService_DownloadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}

func RegisterFileServiceServer(s *grpc.Server, srv FileServiceServer) {
	s.RegisterService(&_FileService_serviceDesc, srv)
}

func _FileService_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileServiceServer).UploadFile(&fileServiceUploadFileServer{stream})
}

type FileService_UploadFileServer interface {
	SendAndClose(*UploadResponse) error
	Recv() (*UploadRequest, error)
	grpc.ServerStream
}

type fileServiceUploadFileServer struct {
	grpc.ServerStream
}

func (x *fileServiceUploadFileServer) SendAndClose(m *UploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileServiceUploadFileServer) Recv() (*UploadRequest, error) {
	m := new(UploadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FileService_ListFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).ListFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.FileService/ListFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).ListFiles(ctx, req.(*ListFilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_DownloadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileServiceServer).DownloadFile(m, &fileServiceDownloadFileServer{stream})
}

type FileService_DownloadFileServer interface {
	Send(*DownloadResponse) error
	grpc.ServerStream
}

type fileServiceDownloadFileServer struct {
	grpc.ServerStream
}

func (x *fileServiceDownloadFileServer) Send(m *DownloadResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _FileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.FileService",
	HandlerType: (*FileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListFiles",
			Handler:    _FileService_ListFiles_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadFile",
			Handler:       _FileService_UploadFile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DownloadFile",
			Handler:       _FileService_DownloadFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "structure.proto",
}
