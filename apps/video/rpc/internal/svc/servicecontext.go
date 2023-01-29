package svc

import (
	user "genuine_douyin/apps/user/rpc/usersrv"
	"genuine_douyin/apps/video/model"
	"genuine_douyin/apps/video/rpc/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	VideoModel model.VideoModel

	UserRpc user.UserSrv

	MinioClient *minio.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	minioClient, err := minio.New(c.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(c.Minio.AccessKey, c.Minio.SecretKey, ""),
		Secure: c.Minio.UseSSL,
	})
	if err != nil {
		logx.Errorf("init... new minio client failed, err: %v", err)
		panic(err)
	}
	return &ServiceContext{
		Config:      c,
		VideoModel:  model.NewVideoModel(conn, c.CacheRedis),
		UserRpc:     user.NewUserSrv(zrpc.MustNewClient(c.UserRpc)),
		MinioClient: minioClient,
	}
}
