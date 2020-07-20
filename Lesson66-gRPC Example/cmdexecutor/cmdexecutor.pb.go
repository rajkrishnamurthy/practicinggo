// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cmdexecutor.proto

package cmdexecutor

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Command is the message which will contains the commands from the client system.
type Command struct {
	Cmd                   string   `protobuf:"bytes,1,opt,name=cmd,proto3" json:"cmd,omitempty"`
	Args                  []string `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
	FetchObjToBeReturned  string   `protobuf:"bytes,3,opt,name=fetch_obj_to_be_returned,json=fetchObjToBeReturned,proto3" json:"fetch_obj_to_be_returned,omitempty"`
	FetchFileOrFolderPath string   `protobuf:"bytes,4,opt,name=fetch_file_or_folder_path,json=fetchFileOrFolderPath,proto3" json:"fetch_file_or_folder_path,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}
func (*Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c0da9daf7cdfa44, []int{0}
}

func (m *Command) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command.Unmarshal(m, b)
}
func (m *Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command.Marshal(b, m, deterministic)
}
func (m *Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command.Merge(m, src)
}
func (m *Command) XXX_Size() int {
	return xxx_messageInfo_Command.Size(m)
}
func (m *Command) XXX_DiscardUnknown() {
	xxx_messageInfo_Command.DiscardUnknown(m)
}

var xxx_messageInfo_Command proto.InternalMessageInfo

func (m *Command) GetCmd() string {
	if m != nil {
		return m.Cmd
	}
	return ""
}

func (m *Command) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Command) GetFetchObjToBeReturned() string {
	if m != nil {
		return m.FetchObjToBeReturned
	}
	return ""
}

func (m *Command) GetFetchFileOrFolderPath() string {
	if m != nil {
		return m.FetchFileOrFolderPath
	}
	return ""
}

// CommandOutput will return the output based on the input values
type CommandOutput struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	ErrorCode            int32    `protobuf:"varint,2,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	CmdOutput            []byte   `protobuf:"bytes,3,opt,name=cmd_output,json=cmdOutput,proto3" json:"cmd_output,omitempty"`
	OutputType           string   `protobuf:"bytes,5,opt,name=output_type,json=outputType,proto3" json:"output_type,omitempty"`
	SrcFqPath            string   `protobuf:"bytes,4,opt,name=src_fq_path,json=srcFqPath,proto3" json:"src_fq_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommandOutput) Reset()         { *m = CommandOutput{} }
func (m *CommandOutput) String() string { return proto.CompactTextString(m) }
func (*CommandOutput) ProtoMessage()    {}
func (*CommandOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c0da9daf7cdfa44, []int{1}
}

func (m *CommandOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandOutput.Unmarshal(m, b)
}
func (m *CommandOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandOutput.Marshal(b, m, deterministic)
}
func (m *CommandOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandOutput.Merge(m, src)
}
func (m *CommandOutput) XXX_Size() int {
	return xxx_messageInfo_CommandOutput.Size(m)
}
func (m *CommandOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandOutput.DiscardUnknown(m)
}

var xxx_messageInfo_CommandOutput proto.InternalMessageInfo

func (m *CommandOutput) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CommandOutput) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *CommandOutput) GetCmdOutput() []byte {
	if m != nil {
		return m.CmdOutput
	}
	return nil
}

func (m *CommandOutput) GetOutputType() string {
	if m != nil {
		return m.OutputType
	}
	return ""
}

func (m *CommandOutput) GetSrcFqPath() string {
	if m != nil {
		return m.SrcFqPath
	}
	return ""
}

func init() {
	proto.RegisterType((*Command)(nil), "Command")
	proto.RegisterType((*CommandOutput)(nil), "CommandOutput")
}

func init() { proto.RegisterFile("cmdexecutor.proto", fileDescriptor_4c0da9daf7cdfa44) }

var fileDescriptor_4c0da9daf7cdfa44 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xcb, 0x4e, 0xeb, 0x30,
	0x10, 0x86, 0x8f, 0x7b, 0x39, 0x3d, 0x99, 0x1e, 0x6e, 0x16, 0x48, 0x06, 0x09, 0xa8, 0xba, 0x0a,
	0x9b, 0x0a, 0x81, 0x84, 0x58, 0xb0, 0xa2, 0xa2, 0xdb, 0xa2, 0xa8, 0x7b, 0x2b, 0xb1, 0x27, 0xbd,
	0xa8, 0xee, 0xa4, 0x13, 0x47, 0xa2, 0x8f, 0xc3, 0x8e, 0xc7, 0x44, 0x75, 0x52, 0x09, 0x76, 0x93,
	0xf9, 0xf2, 0xff, 0xfa, 0x6c, 0xc3, 0x99, 0x71, 0x16, 0x3f, 0xd0, 0x54, 0x9e, 0x78, 0x54, 0x30,
	0x79, 0x1a, 0x7e, 0x0a, 0xe8, 0x8d, 0xc9, 0xb9, 0x74, 0x63, 0xe5, 0x29, 0xb4, 0x8d, 0xb3, 0x4a,
	0x0c, 0x44, 0x1c, 0x25, 0xfb, 0x51, 0x4a, 0xe8, 0xa4, 0x3c, 0x2f, 0x55, 0x6b, 0xd0, 0x8e, 0xa3,
	0x24, 0xcc, 0xf2, 0x09, 0x54, 0x8e, 0xde, 0x2c, 0x34, 0x65, 0x2b, 0xed, 0x49, 0x67, 0xa8, 0x19,
	0x7d, 0xc5, 0x1b, 0xb4, 0xaa, 0x1d, 0xa2, 0xe7, 0x81, 0x4f, 0xb3, 0xd5, 0x8c, 0x5e, 0x31, 0x69,
	0x98, 0x7c, 0x86, 0xcb, 0x3a, 0x97, 0x2f, 0xd7, 0xa8, 0x89, 0x75, 0x4e, 0x6b, 0x8b, 0xac, 0x8b,
	0xd4, 0x2f, 0x54, 0x27, 0x04, 0x2f, 0xc2, 0x0f, 0x93, 0xe5, 0x1a, 0xa7, 0x3c, 0x09, 0xf4, 0x3d,
	0xf5, 0x8b, 0xe1, 0x97, 0x80, 0xa3, 0xc6, 0x71, 0x5a, 0xf9, 0xa2, 0xf2, 0x52, 0x41, 0xcf, 0x61,
	0x59, 0xa6, 0x73, 0x6c, 0x6c, 0x0f, 0x9f, 0xf2, 0x1a, 0x00, 0x99, 0x89, 0xb5, 0x21, 0x8b, 0xaa,
	0x35, 0x10, 0x71, 0x37, 0x89, 0xc2, 0x66, 0x4c, 0x36, 0x60, 0xe3, 0xac, 0xa6, 0x50, 0x13, 0x74,
	0xff, 0x27, 0x91, 0x71, 0x87, 0xde, 0x5b, 0xe8, 0xd7, 0x48, 0xfb, 0x5d, 0x81, 0xaa, 0x1b, 0xba,
	0xa1, 0x5e, 0xcd, 0x76, 0x05, 0xca, 0x1b, 0xe8, 0x97, 0x6c, 0x74, 0xbe, 0xfd, 0xa9, 0x1d, 0x95,
	0x6c, 0x26, 0xdb, 0xbd, 0xea, 0xc3, 0x0b, 0x9c, 0x34, 0xa6, 0x6f, 0xcd, 0x3d, 0xcb, 0x3b, 0xe8,
	0xd5, 0x33, 0xca, 0x7f, 0xa3, 0x06, 0x5e, 0x1d, 0x8f, 0x7e, 0x1d, 0x68, 0xf8, 0x27, 0x16, 0xf7,
	0x22, 0xfb, 0x1b, 0xde, 0xe4, 0xf1, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xa6, 0xdb, 0x7d, 0x8b, 0xa8,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CommandExecutorClient is the client API for CommandExecutor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommandExecutorClient interface {
	Execute(ctx context.Context, opts ...grpc.CallOption) (CommandExecutor_ExecuteClient, error)
}

type commandExecutorClient struct {
	cc *grpc.ClientConn
}

func NewCommandExecutorClient(cc *grpc.ClientConn) CommandExecutorClient {
	return &commandExecutorClient{cc}
}

func (c *commandExecutorClient) Execute(ctx context.Context, opts ...grpc.CallOption) (CommandExecutor_ExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CommandExecutor_serviceDesc.Streams[0], "/CommandExecutor/Execute", opts...)
	if err != nil {
		return nil, err
	}
	x := &commandExecutorExecuteClient{stream}
	return x, nil
}

type CommandExecutor_ExecuteClient interface {
	Send(*Command) error
	Recv() (*CommandOutput, error)
	grpc.ClientStream
}

type commandExecutorExecuteClient struct {
	grpc.ClientStream
}

func (x *commandExecutorExecuteClient) Send(m *Command) error {
	return x.ClientStream.SendMsg(m)
}

func (x *commandExecutorExecuteClient) Recv() (*CommandOutput, error) {
	m := new(CommandOutput)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CommandExecutorServer is the server API for CommandExecutor service.
type CommandExecutorServer interface {
	Execute(CommandExecutor_ExecuteServer) error
}

// UnimplementedCommandExecutorServer can be embedded to have forward compatible implementations.
type UnimplementedCommandExecutorServer struct {
}

func (*UnimplementedCommandExecutorServer) Execute(srv CommandExecutor_ExecuteServer) error {
	return status.Errorf(codes.Unimplemented, "method Execute not implemented")
}

func RegisterCommandExecutorServer(s *grpc.Server, srv CommandExecutorServer) {
	s.RegisterService(&_CommandExecutor_serviceDesc, srv)
}

func _CommandExecutor_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CommandExecutorServer).Execute(&commandExecutorExecuteServer{stream})
}

type CommandExecutor_ExecuteServer interface {
	Send(*CommandOutput) error
	Recv() (*Command, error)
	grpc.ServerStream
}

type commandExecutorExecuteServer struct {
	grpc.ServerStream
}

func (x *commandExecutorExecuteServer) Send(m *CommandOutput) error {
	return x.ServerStream.SendMsg(m)
}

func (x *commandExecutorExecuteServer) Recv() (*Command, error) {
	m := new(Command)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CommandExecutor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CommandExecutor",
	HandlerType: (*CommandExecutorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Execute",
			Handler:       _CommandExecutor_Execute_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "cmdexecutor.proto",
}