package interceptor

import (
	"encoding/json"
	"fgzs/internal/errorx"
	"fgzs/internal/errorx/i18n"
	"fgzs/internal/response"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

func HttpErrorHandler(err error) (int, interface{}) {
	var e *errorx.BusinessErr
	if err == nil {
		e = errorx.Success
	} else {
		causeErr := errors.Cause(err)
		//自定义错误类型检测
		if businessErr, ok := causeErr.(*errorx.BusinessErr); ok {
			e = businessErr
		} else {
			// grpc err错误
			if grpcErr, ok := status.FromError(causeErr); ok {
				if _, ok := errorx.BusinessErrs[int(grpcErr.Code())]; ok {
					_ = json.Unmarshal([]byte(grpcErr.Message()), &e)
				}
			}
		}
	}
	if e == nil {
		e = errorx.InternalServerError.WithDetail(err)
	}
	if e.ErrLevel == errorx.WarnLevel {
		logx.Errorf("【API-WARN】 : %+v ", e)
	}
	if e.ErrLevel == errorx.ErrLevel {
		logx.Errorf("【API-ERR】 : %+v ", e)
	}
	return e.HttpCode, response.HttpBusinessError(e, i18n.EnUS)
}
