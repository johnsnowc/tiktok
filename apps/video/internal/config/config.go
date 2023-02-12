package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	Mysql struct {
		DataSource string
	}
	RedisCacheConf redis.RedisConf

	UserRpc     zrpc.RpcClientConf
	FavoriteRpc zrpc.RpcClientConf

	Minio struct {
		Endpoint    string
		AccessKey   string
		SecretKey   string
		UseSSL      bool
		VideoBucket string
		CoverBucket string
		Location    string
		ContentType string
	}
}
