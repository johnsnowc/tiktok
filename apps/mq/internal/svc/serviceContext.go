package svc

import (
	"genuine_douyin/apps/comment/commentsrv"
	"genuine_douyin/apps/favorite/favoritesrv"
	"genuine_douyin/apps/mq/internal/config"
	"genuine_douyin/apps/relation/relationsrv"
	"genuine_douyin/apps/video/videosrv"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	FavoriteSvcRpcClient favoritesrv.FavoriteSrv
	RelationSvcRpcClient relationsrv.RelationSrv
	CommentSvcRpcClient  commentsrv.CommentSrv
	VideoSvcRpcClient    videosrv.VideoSrv

	RedisCache *redis.Redis
	ScriptADD  string // 在Mqs中初始化
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		FavoriteSvcRpcClient: favoritesrv.NewFavoriteSrv(zrpc.MustNewClient(c.UserOptServiceConf)),
		RelationSvcRpcClient: relationsrv.NewRelationSrv(zrpc.MustNewClient(c.UserOptServiceConf)),
		CommentSvcRpcClient:  commentsrv.NewCommentSrv(zrpc.MustNewClient(c.UserOptServiceConf)),
		VideoSvcRpcClient:    videosrv.NewVideoSrv(zrpc.MustNewClient(c.VideoService)),

		RedisCache: c.RedisCacheConf.NewRedis(),
	}
}
