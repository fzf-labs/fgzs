package middleware

import (
	"fgzs/internal/errorx"
	"fgzs/internal/meta"
	"fgzs/internal/response"
	"fgzs/pkg/webhook"
	"net/http"
)

type ArgumentMiddleware struct {
	Mode    string
	WebHook *webhook.WebHookConfig
}

func NewArgumentMiddleware(mode string, webHook *webhook.WebHookConfig) *ArgumentMiddleware {
	return &ArgumentMiddleware{
		Mode:    mode,
		WebHook: webHook,
	}
}

func (m *ArgumentMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//自定义参数设置
		xParams := meta.GetXParamsByHeader(r)
		if xParams == nil {
			response.Err(r, w, errorx.ParamErr)
			return
		}
		r = r.WithContext(meta.SetXParams(r.Context(), xParams))
		next(w, r)
	}
}
