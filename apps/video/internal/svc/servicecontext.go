package svc

import (
	favorite "genuine_douyin/apps/favorite/favoritesrv"
	user "genuine_douyin/apps/user/usersrv"
	"genuine_douyin/apps/video/dal"
	"genuine_douyin/apps/video/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config

	VideoModel *gorm.DB
	RedisCache *redis.Redis

	UserRpc     user.UserSrv
	FavoriteRpc favorite.FavoriteSrv

	MinioClient *minio.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
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
		VideoModel:  dal.NewGorm(c.Mysql.DataSource),
		RedisCache:  c.RedisCacheConf.NewRedis(),
		UserRpc:     user.NewUserSrv(zrpc.MustNewClient(c.UserRpc)),
		FavoriteRpc: favorite.NewFavoriteSrv(zrpc.MustNewClient(c.FavoriteRpc)),
		MinioClient: minioClient,
	}
}
