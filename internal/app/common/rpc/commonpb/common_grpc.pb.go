// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: common.proto

package commonpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CommonClient is the client API for Common service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommonClient interface {
	//验证码
	CaptchaGenerate(ctx context.Context, in *CaptchaGenerateReq, opts ...grpc.CallOption) (*CaptchaGenerateResp, error)
	//验证码校验
	CaptchaCheck(ctx context.Context, in *CaptchaCheckReq, opts ...grpc.CallOption) (*CaptchaCheckResp, error)
	//短信发送
	SmsSend(ctx context.Context, in *SmsSendReq, opts ...grpc.CallOption) (*SmsSendResp, error)
	//短信发送校验
	SmsCheck(ctx context.Context, in *SmsCheckReq, opts ...grpc.CallOption) (*SmsCheckResp, error)
	//短信记录
	SmsRecord(ctx context.Context, in *SmsRecordReq, opts ...grpc.CallOption) (*SmsRecordResp, error)
	//敏感词检测
	SensitiveWordCheck(ctx context.Context, in *SensitiveWordCheckReq, opts ...grpc.CallOption) (*SensitiveWordCheckResp, error)
	//敏感词查询
	SensitiveWordSearch(ctx context.Context, in *SensitiveWordSearchReq, opts ...grpc.CallOption) (*SensitiveWordSearchResp, error)
	//敏感词添加
	SensitiveWordAdd(ctx context.Context, in *SensitiveWordAddReq, opts ...grpc.CallOption) (*SensitiveWordAddResp, error)
	//敏感词删除
	SensitiveWordDel(ctx context.Context, in *SensitiveWordDelReq, opts ...grpc.CallOption) (*SensitiveWordDelResp, error)
	//敏感词缓存
	SensitiveWordCache(ctx context.Context, in *SensitiveWordCacheReq, opts ...grpc.CallOption) (*SensitiveWordCacheResp, error)
}

type commonClient struct {
	cc grpc.ClientConnInterface
}

func NewCommonClient(cc grpc.ClientConnInterface) CommonClient {
	return &commonClient{cc}
}

func (c *commonClient) CaptchaGenerate(ctx context.Context, in *CaptchaGenerateReq, opts ...grpc.CallOption) (*CaptchaGenerateResp, error) {
	out := new(CaptchaGenerateResp)
	err := c.cc.Invoke(ctx, "/commonpb.common/CaptchaGenerate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonClient) CaptchaCheck(ctx context.Context, in *CaptchaCheckReq, opts ...grpc.CallOption) (*CaptchaCheckResp, error) {
	out := new(CaptchaCheckResp)
	err := c.cc.Invoke(ctx, "/commonpb.common/CaptchaCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonClient) SmsSend(ctx context.Context, in *SmsSendReq, opts ...grpc.CallOption) (*SmsSendResp, error) {
	out := new(SmsSendResp)
	err := c.cc.Invoke(ctx, "/commonpb.common/SmsSend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonClient) SmsCheck(ctx context.Context, in *SmsCheckReq, opts ...grpc.CallOption) (*SmsCheckResp, error) {
	out := new(SmsCheckResp)
	err := c.cc.Invoke(ctx, "/commonpb.common/SmsCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonClient) SmsRecord(ctx context.Context, in *SmsRecordReq, opts ...grpc.CallOption) (*SmsRecordResp, error) {
	out := new(SmsRecordResp)
	err := c.cc.Invoke(ctx, "/commonpb.common/SmsRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonClient) SensitiveWordCheck(ctx context.Context, in *SensitiveWordCheckReq, opts ...grpc.CallOption) (*SensitiveWordCheckResp, error) {
	out := new(SensitiveWordCheckResp)
	err := c.cc.Invoke(ctx, "/commonpb.common/SensitiveWordCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonClient) SensitiveWordSearch(ctx context.Context, in *SensitiveWordSearchReq, opts ...grpc.CallOption) (*SensitiveWordSearchResp, error) {
	out := new(SensitiveWordSearchResp)
	err := c.cc.Invoke(ctx, "/commonpb.common/SensitiveWordSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonClient) SensitiveWordAdd(ctx context.Context, in *SensitiveWordAddReq, opts ...grpc.CallOption) (*SensitiveWordAddResp, error) {
	out := new(SensitiveWordAddResp)
	err := c.cc.Invoke(ctx, "/commonpb.common/SensitiveWordAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonClient) SensitiveWordDel(ctx context.Context, in *SensitiveWordDelReq, opts ...grpc.CallOption) (*SensitiveWordDelResp, error) {
	out := new(SensitiveWordDelResp)
	err := c.cc.Invoke(ctx, "/commonpb.common/SensitiveWordDel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonClient) SensitiveWordCache(ctx context.Context, in *SensitiveWordCacheReq, opts ...grpc.CallOption) (*SensitiveWordCacheResp, error) {
	out := new(SensitiveWordCacheResp)
	err := c.cc.Invoke(ctx, "/commonpb.common/SensitiveWordCache", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommonServer is the server API for Common service.
// All implementations must embed UnimplementedCommonServer
// for forward compatibility
type CommonServer interface {
	//验证码
	CaptchaGenerate(context.Context, *CaptchaGenerateReq) (*CaptchaGenerateResp, error)
	//验证码校验
	CaptchaCheck(context.Context, *CaptchaCheckReq) (*CaptchaCheckResp, error)
	//短信发送
	SmsSend(context.Context, *SmsSendReq) (*SmsSendResp, error)
	//短信发送校验
	SmsCheck(context.Context, *SmsCheckReq) (*SmsCheckResp, error)
	//短信记录
	SmsRecord(context.Context, *SmsRecordReq) (*SmsRecordResp, error)
	//敏感词检测
	SensitiveWordCheck(context.Context, *SensitiveWordCheckReq) (*SensitiveWordCheckResp, error)
	//敏感词查询
	SensitiveWordSearch(context.Context, *SensitiveWordSearchReq) (*SensitiveWordSearchResp, error)
	//敏感词添加
	SensitiveWordAdd(context.Context, *SensitiveWordAddReq) (*SensitiveWordAddResp, error)
	//敏感词删除
	SensitiveWordDel(context.Context, *SensitiveWordDelReq) (*SensitiveWordDelResp, error)
	//敏感词缓存
	SensitiveWordCache(context.Context, *SensitiveWordCacheReq) (*SensitiveWordCacheResp, error)
	mustEmbedUnimplementedCommonServer()
}

// UnimplementedCommonServer must be embedded to have forward compatible implementations.
type UnimplementedCommonServer struct {
}

func (UnimplementedCommonServer) CaptchaGenerate(context.Context, *CaptchaGenerateReq) (*CaptchaGenerateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CaptchaGenerate not implemented")
}
func (UnimplementedCommonServer) CaptchaCheck(context.Context, *CaptchaCheckReq) (*CaptchaCheckResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CaptchaCheck not implemented")
}
func (UnimplementedCommonServer) SmsSend(context.Context, *SmsSendReq) (*SmsSendResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SmsSend not implemented")
}
func (UnimplementedCommonServer) SmsCheck(context.Context, *SmsCheckReq) (*SmsCheckResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SmsCheck not implemented")
}
func (UnimplementedCommonServer) SmsRecord(context.Context, *SmsRecordReq) (*SmsRecordResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SmsRecord not implemented")
}
func (UnimplementedCommonServer) SensitiveWordCheck(context.Context, *SensitiveWordCheckReq) (*SensitiveWordCheckResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SensitiveWordCheck not implemented")
}
func (UnimplementedCommonServer) SensitiveWordSearch(context.Context, *SensitiveWordSearchReq) (*SensitiveWordSearchResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SensitiveWordSearch not implemented")
}
func (UnimplementedCommonServer) SensitiveWordAdd(context.Context, *SensitiveWordAddReq) (*SensitiveWordAddResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SensitiveWordAdd not implemented")
}
func (UnimplementedCommonServer) SensitiveWordDel(context.Context, *SensitiveWordDelReq) (*SensitiveWordDelResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SensitiveWordDel not implemented")
}
func (UnimplementedCommonServer) SensitiveWordCache(context.Context, *SensitiveWordCacheReq) (*SensitiveWordCacheResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SensitiveWordCache not implemented")
}
func (UnimplementedCommonServer) mustEmbedUnimplementedCommonServer() {}

// UnsafeCommonServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommonServer will
// result in compilation errors.
type UnsafeCommonServer interface {
	mustEmbedUnimplementedCommonServer()
}

func RegisterCommonServer(s grpc.ServiceRegistrar, srv CommonServer) {
	s.RegisterService(&Common_ServiceDesc, srv)
}

func _Common_CaptchaGenerate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CaptchaGenerateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServer).CaptchaGenerate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commonpb.common/CaptchaGenerate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServer).CaptchaGenerate(ctx, req.(*CaptchaGenerateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Common_CaptchaCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CaptchaCheckReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServer).CaptchaCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commonpb.common/CaptchaCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServer).CaptchaCheck(ctx, req.(*CaptchaCheckReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Common_SmsSend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmsSendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServer).SmsSend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commonpb.common/SmsSend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServer).SmsSend(ctx, req.(*SmsSendReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Common_SmsCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmsCheckReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServer).SmsCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commonpb.common/SmsCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServer).SmsCheck(ctx, req.(*SmsCheckReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Common_SmsRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmsRecordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServer).SmsRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commonpb.common/SmsRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServer).SmsRecord(ctx, req.(*SmsRecordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Common_SensitiveWordCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SensitiveWordCheckReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServer).SensitiveWordCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commonpb.common/SensitiveWordCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServer).SensitiveWordCheck(ctx, req.(*SensitiveWordCheckReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Common_SensitiveWordSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SensitiveWordSearchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServer).SensitiveWordSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commonpb.common/SensitiveWordSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServer).SensitiveWordSearch(ctx, req.(*SensitiveWordSearchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Common_SensitiveWordAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SensitiveWordAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServer).SensitiveWordAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commonpb.common/SensitiveWordAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServer).SensitiveWordAdd(ctx, req.(*SensitiveWordAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Common_SensitiveWordDel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SensitiveWordDelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServer).SensitiveWordDel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commonpb.common/SensitiveWordDel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServer).SensitiveWordDel(ctx, req.(*SensitiveWordDelReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Common_SensitiveWordCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SensitiveWordCacheReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServer).SensitiveWordCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commonpb.common/SensitiveWordCache",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServer).SensitiveWordCache(ctx, req.(*SensitiveWordCacheReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Common_ServiceDesc is the grpc.ServiceDesc for Common service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Common_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "commonpb.common",
	HandlerType: (*CommonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CaptchaGenerate",
			Handler:    _Common_CaptchaGenerate_Handler,
		},
		{
			MethodName: "CaptchaCheck",
			Handler:    _Common_CaptchaCheck_Handler,
		},
		{
			MethodName: "SmsSend",
			Handler:    _Common_SmsSend_Handler,
		},
		{
			MethodName: "SmsCheck",
			Handler:    _Common_SmsCheck_Handler,
		},
		{
			MethodName: "SmsRecord",
			Handler:    _Common_SmsRecord_Handler,
		},
		{
			MethodName: "SensitiveWordCheck",
			Handler:    _Common_SensitiveWordCheck_Handler,
		},
		{
			MethodName: "SensitiveWordSearch",
			Handler:    _Common_SensitiveWordSearch_Handler,
		},
		{
			MethodName: "SensitiveWordAdd",
			Handler:    _Common_SensitiveWordAdd_Handler,
		},
		{
			MethodName: "SensitiveWordDel",
			Handler:    _Common_SensitiveWordDel_Handler,
		},
		{
			MethodName: "SensitiveWordCache",
			Handler:    _Common_SensitiveWordCache_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common.proto",
}