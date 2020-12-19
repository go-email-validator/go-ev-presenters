// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// EmailValidationClient is the client API for EmailValidation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmailValidationClient interface {
	SingleValidation(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*EmailResponse, error)
}

type emailValidationClient struct {
	cc grpc.ClientConnInterface
}

func NewEmailValidationClient(cc grpc.ClientConnInterface) EmailValidationClient {
	return &emailValidationClient{cc}
}

func (c *emailValidationClient) SingleValidation(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*EmailResponse, error) {
	out := new(EmailResponse)
	err := c.cc.Invoke(ctx, "/github.com.go_email_validator.go_ev_presenters.api.v1.EmailValidation/singleValidation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmailValidationServer is the server API for EmailValidation service.
// All implementations must embed UnimplementedEmailValidationServer
// for forward compatibility
type EmailValidationServer interface {
	SingleValidation(context.Context, *EmailRequest) (*EmailResponse, error)
	mustEmbedUnimplementedEmailValidationServer()
}

// UnimplementedEmailValidationServer must be embedded to have forward compatible implementations.
type UnimplementedEmailValidationServer struct {
}

func (UnimplementedEmailValidationServer) SingleValidation(context.Context, *EmailRequest) (*EmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SingleValidation not implemented")
}
func (UnimplementedEmailValidationServer) mustEmbedUnimplementedEmailValidationServer() {}

// UnsafeEmailValidationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmailValidationServer will
// result in compilation errors.
type UnsafeEmailValidationServer interface {
	mustEmbedUnimplementedEmailValidationServer()
}

func RegisterEmailValidationServer(s grpc.ServiceRegistrar, srv EmailValidationServer) {
	s.RegisterService(&_EmailValidation_serviceDesc, srv)
}

func _EmailValidation_SingleValidation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailValidationServer).SingleValidation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.go_email_validator.go_ev_presenters.api.v1.EmailValidation/singleValidation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailValidationServer).SingleValidation(ctx, req.(*EmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EmailValidation_serviceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.go_email_validator.go_ev_presenters.api.v1.EmailValidation",
	HandlerType: (*EmailValidationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "singleValidation",
			Handler:    _EmailValidation_SingleValidation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ev.proto",
}