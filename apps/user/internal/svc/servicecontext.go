package svc

import (
	"genuine_douyin/apps/user/dal"
	"genuine_douyin/apps/user/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config     config.Config
	UserModel  *gorm.DB
	RedisCache *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		UserModel:  dal.NewGorm(c.Mysql.DataSource),
		RedisCache: c.RedisCacheConf.NewRedis(),
	}
}
