// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.20.3
// source: library.proto

package generated

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	LibraryService_AddBook_FullMethodName    = "/library.LibraryService/AddBook"
	LibraryService_GetBooks_FullMethodName   = "/library.LibraryService/GetBooks"
	LibraryService_BorrowBook_FullMethodName = "/library.LibraryService/BorrowBook"
	LibraryService_DeleteBook_FullMethodName = "/library.LibraryService/DeleteBook"
	LibraryService_UpdateBook_FullMethodName = "/library.LibraryService/UpdateBook"
)

// LibraryServiceClient is the client API for LibraryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LibraryServiceClient interface {
	AddBook(ctx context.Context, in *AddBookRequest, opts ...grpc.CallOption) (*AddBookResponse, error)
	GetBooks(ctx context.Context, in *GetBooksRequest, opts ...grpc.CallOption) (*GetBooksResponse, error)
	BorrowBook(ctx context.Context, in *BorrowBookRequest, opts ...grpc.CallOption) (*BorrowBookResponse, error)
	DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*DeleteBookResponse, error)
	UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*UpdateBookResponse, error)
}

type libraryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLibraryServiceClient(cc grpc.ClientConnInterface) LibraryServiceClient {
	return &libraryServiceClient{cc}
}

func (c *libraryServiceClient) AddBook(ctx context.Context, in *AddBookRequest, opts ...grpc.CallOption) (*AddBookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddBookResponse)
	err := c.cc.Invoke(ctx, LibraryService_AddBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) GetBooks(ctx context.Context, in *GetBooksRequest, opts ...grpc.CallOption) (*GetBooksResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBooksResponse)
	err := c.cc.Invoke(ctx, LibraryService_GetBooks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) BorrowBook(ctx context.Context, in *BorrowBookRequest, opts ...grpc.CallOption) (*BorrowBookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BorrowBookResponse)
	err := c.cc.Invoke(ctx, LibraryService_BorrowBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*DeleteBookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteBookResponse)
	err := c.cc.Invoke(ctx, LibraryService_DeleteBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*UpdateBookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateBookResponse)
	err := c.cc.Invoke(ctx, LibraryService_UpdateBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LibraryServiceServer is the server API for LibraryService service.
// All implementations must embed UnimplementedLibraryServiceServer
// for forward compatibility.
type LibraryServiceServer interface {
	AddBook(context.Context, *AddBookRequest) (*AddBookResponse, error)
	GetBooks(context.Context, *GetBooksRequest) (*GetBooksResponse, error)
	BorrowBook(context.Context, *BorrowBookRequest) (*BorrowBookResponse, error)
	DeleteBook(context.Context, *DeleteBookRequest) (*DeleteBookResponse, error)
	UpdateBook(context.Context, *UpdateBookRequest) (*UpdateBookResponse, error)
	mustEmbedUnimplementedLibraryServiceServer()
}

// UnimplementedLibraryServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLibraryServiceServer struct{}

func (UnimplementedLibraryServiceServer) AddBook(context.Context, *AddBookRequest) (*AddBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBook not implemented")
}
func (UnimplementedLibraryServiceServer) GetBooks(context.Context, *GetBooksRequest) (*GetBooksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBooks not implemented")
}
func (UnimplementedLibraryServiceServer) BorrowBook(context.Context, *BorrowBookRequest) (*BorrowBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BorrowBook not implemented")
}
func (UnimplementedLibraryServiceServer) DeleteBook(context.Context, *DeleteBookRequest) (*DeleteBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBook not implemented")
}
func (UnimplementedLibraryServiceServer) UpdateBook(context.Context, *UpdateBookRequest) (*UpdateBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBook not implemented")
}
func (UnimplementedLibraryServiceServer) mustEmbedUnimplementedLibraryServiceServer() {}
func (UnimplementedLibraryServiceServer) testEmbeddedByValue()                        {}

// UnsafeLibraryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LibraryServiceServer will
// result in compilation errors.
type UnsafeLibraryServiceServer interface {
	mustEmbedUnimplementedLibraryServiceServer()
}

func RegisterLibraryServiceServer(s grpc.ServiceRegistrar, srv LibraryServiceServer) {
	// If the following call pancis, it indicates UnimplementedLibraryServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LibraryService_ServiceDesc, srv)
}

func _LibraryService_AddBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).AddBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LibraryService_AddBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).AddBook(ctx, req.(*AddBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_GetBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).GetBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LibraryService_GetBooks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).GetBooks(ctx, req.(*GetBooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_BorrowBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BorrowBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).BorrowBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LibraryService_BorrowBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).BorrowBook(ctx, req.(*BorrowBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_DeleteBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).DeleteBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LibraryService_DeleteBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).DeleteBook(ctx, req.(*DeleteBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_UpdateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).UpdateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LibraryService_UpdateBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).UpdateBook(ctx, req.(*UpdateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LibraryService_ServiceDesc is the grpc.ServiceDesc for LibraryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LibraryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "library.LibraryService",
	HandlerType: (*LibraryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddBook",
			Handler:    _LibraryService_AddBook_Handler,
		},
		{
			MethodName: "GetBooks",
			Handler:    _LibraryService_GetBooks_Handler,
		},
		{
			MethodName: "BorrowBook",
			Handler:    _LibraryService_BorrowBook_Handler,
		},
		{
			MethodName: "DeleteBook",
			Handler:    _LibraryService_DeleteBook_Handler,
		},
		{
			MethodName: "UpdateBook",
			Handler:    _LibraryService_UpdateBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "library.proto",
}