// Code generated by goctl. DO NOT EDIT.
package handler

import (
	oauthuser "micro/api/internal/handler/oauth/user"
	"micro/api/internal/svc"
	"net/http"

	"github.com/tal-tech/go-zero/rest"
)

// 路由注册
func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/oauth/login",
				Handler: oauthuser.LoginHandler(serverCtx),
			},
		},
	)
}
