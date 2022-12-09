package response

import (
	"encoding/json"
	"fgzs/internal/errorx"
	"fgzs/pkg/util/validutil"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"net/http"
)

// H is a shortcut for map[string]interface{}
type H map[string]interface{}

type HttpResponse struct {
	Code    int         `json:"code"`              // HTTP Code
	Message string      `json:"message"`           // 描述信息
	Data    interface{} `json:"data"`              // 返回信息
	ErrMsg  string      `json:"err_msg,omitempty"` // 错误信息
	Err     string      `json:"err,omitempty"`     // 错误
}

func HttpSuccess(resp interface{}, lang string) *HttpResponse {
	r := &HttpResponse{
		Code:    errorx.Success.GetBusinessCode(),
		Message: errorx.Success.GetMessage(lang),
		Data:    resp,
	}
	if validutil.IsZero(r.Data) {
		r.Data = H{}
	}
	return r
}
func HttpBusinessError(err *errorx.BusinessErr, lang string) *HttpResponse {
	r := &HttpResponse{
		Code:    err.GetBusinessCode(),
		Message: err.GetMessage(lang),
		Data:    err.GetErrData(),
		ErrMsg:  err.GetErrMsg(),
		Err:     err.GetDetail(),
	}
	return r
}

func Http(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	lang := r.Header.Get("Accept-Language")
	if err == nil {
		httpx.OkJson(w, HttpSuccess(resp, lang))
		return
	}
	Err(r, w, err)
}

func Err(r *http.Request, w http.ResponseWriter, err error) {
	lang := r.Header.Get("Accept-Language")
	var e *errorx.BusinessErr
	causeErr := errors.Cause(err)
	//自定义错误类型检测
	if businessErr, ok := causeErr.(*errorx.BusinessErr); ok {
		e = businessErr
	} else if grpcErr, ok := status.FromError(causeErr); ok {
		// grpc err错误
		if _, ok := errorx.BusinessErrs[int(grpcErr.Code())]; ok {
			_ = json.Unmarshal([]byte(grpcErr.Message()), &e)
		} else {
			e = errorx.InternalServerError.WithDetail(err)
		}
	} else {
		e = errorx.InternalServerError.WithDetail(err)
	}
	if e == nil {
		e = errorx.InternalServerError.WithDetail(err)
	}
	if e.ErrLevel == errorx.WarnLevel || e.ErrLevel == errorx.ErrLevel {
		logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v ", e)
	}
	httpx.WriteJson(w, e.GetHttpCode(), HttpBusinessError(e, lang))
}
