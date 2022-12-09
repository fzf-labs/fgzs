// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"fgzs/internal/app/user/api/internal/handler/auth"
	"fgzs/internal/app/user/api/internal/svc"
	"net/http"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/users/login",
				Handler: auth.UserLoginHandler(serverCtx),
			},
		},
	)
}
