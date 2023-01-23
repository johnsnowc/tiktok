package svc

import (
	user "genuine_douyin/apps/user/rpc/usersrv"
	"genuine_douyin/apps/video/model"
	"genuine_douyin/apps/video/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	VideoModel model.VideoModel

	UserRpc user.UserSrv
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		VideoModel: model.NewVideoModel(conn, c.CacheRedis),
		UserRpc:    user.NewUserSrv(zrpc.MustNewClient(c.UserRpc)),
	}
}
