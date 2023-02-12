package svc

import (
	"genuine_douyin/apps/api/internal/config"
	"genuine_douyin/apps/api/internal/middleware"
	user "genuine_douyin/apps/user/usersrv"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	UserRpc user.UserSrv
	//VideoRpc video.VideoSrv

	AuthJWT rest.Middleware
	IsLogin rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		UserRpc: user.NewUserSrv(zrpc.MustNewClient(c.UserRpc)),
		//VideoRpc: video.NewVideoSrv(zrpc.MustNewClient(c.VideoRpc)),

		AuthJWT: middleware.NewAuthJWTMiddleware().Handle,
		IsLogin: middleware.NewIsLoginMiddleware().Handle,
	}
}
