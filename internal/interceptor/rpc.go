package interceptor

import (
	"context"
	"encoding/json"
	"fgzs/internal/define/vars"
	"fgzs/internal/errorx"
	"fgzs/pkg/conv"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Validator interface {
	ValidateAll() error
}

// RpcErrInterceptor Rpc  错误拦截器 rpcinterceptor
func RpcErrInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	//参数校验
	if r, ok := req.(Validator); ok {
		if err := r.ValidateAll(); err != nil {
			logx.WithContext(ctx).Errorf("【RPC-ERR】 %+v", err)
			e := errorx.ParamErr.WithDetail(err)
			//转成grpc err
			return nil, status.Error(codes.Code(e.GetBusinessCode()), e.Error())
		}
	}
	resp, err = handler(ctx, req)
	//错误拦截
	if err != nil {
		causeErr := errors.Cause(err)                    // err类型
		if e, ok := causeErr.(*errorx.BusinessErr); ok { //自定义错误类型
			logx.WithContext(ctx).Errorf("【RPC-ERR】 %+v", err)
			//转成grpc err
			err = status.Error(codes.Code(e.GetBusinessCode()), e.Error())
		} else {
			logx.WithContext(ctx).Errorf("【RPC-ERR】 %+v", err)
		}

	}
	return resp, err
}

// MetaDataInterceptor grpc metadata 拦截器
// zrpc.MustNewClient(c.BehaviorRpc, zrpc.WithUnaryClientInterceptor(interceptor.MetaDataInterceptor))
//func MetaDataInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
//	channelID := contextKey(ctx, constants.ChannelID)
//	deviceType := contextKey(ctx, constants.DeviceType)
//	realIP := contextKey(ctx, constants.RealIP)
//	//md := metadata.Pairs()
//	//md.Set(constants.RealIP, realIP)
//	//md.Set(constants.ChannelID, channelID)
//	//md.Set(constants.DeviceType, deviceType)
//	ctx = metadata.AppendToOutgoingContext(ctx, constants.RealIP, realIP, constants.ChannelID, channelID, constants.DeviceType, deviceType)
//	err := invoker(ctx, method, req, reply, cc, opts...)
//	if err != nil {
//		return err
//	}
//	return nil
//}

func ContextKey(ctx context.Context, t string) string {
	return conv.String(ctx.Value(vars.ContextWithValueKey(t)))
}

func SpanInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	span := oteltrace.SpanFromContext(ctx)
	if span.IsRecording() {
		marshal, _ := json.Marshal(req)
		span.SetAttributes(attribute.Key("rpc.req").String(string(marshal)))
		span.End()
	}
	resp, err = handler(ctx, req)
	return resp, err
}
