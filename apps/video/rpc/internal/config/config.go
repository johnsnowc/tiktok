package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	Mysql struct {
		DataSource string
	}

	CacheRedis cache.CacheConf

	UserRpc zrpc.RpcClientConf

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
