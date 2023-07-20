package svc

import (
	"genuine_douyin/apps/asynqJob/server/internal/config"
	"genuine_douyin/apps/favorite/favoritesrv"
	"genuine_douyin/apps/relation/relationsrv"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	AsynqServer *asynq.Server
	RedisCache  *redis.Redis

	FavoriteSvcRpcClient favoritesrv.FavoriteSrv
	RelationSvcRpcClient relationsrv.RelationSrv

	ScriptREMTag string
}

const scriptLoadREM = "local arr = redis.call('SMEMBERS', KEYS[1])  for i=1, #arr do redis.call('SREM', KEYS[1], arr[i] ) end return arr"

func NewServiceContext(c config.Config) *ServiceContext {
	ServiceContext := &ServiceContext{
		Config:      c,
		AsynqServer: newAsynqServer(c),
		RedisCache:  c.RedisCacheConf.NewRedis(),

		FavoriteSvcRpcClient: favoritesrv.NewFavoriteSrv(zrpc.MustNewClient(c.UserOptServiceConf)),
		RelationSvcRpcClient: relationsrv.NewRelationSrv(zrpc.MustNewClient(c.UserOptServiceConf)),
	}

	ServiceContext.ScriptREMTag, _ = ServiceContext.RedisCache.ScriptLoad(scriptLoadREM)

	return ServiceContext
}
