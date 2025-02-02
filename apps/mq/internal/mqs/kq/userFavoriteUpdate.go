package kq

import (
	"context"
	"encoding/json"
	"fmt"
	"genuine_douyin/apps/favorite/favorite"
	"genuine_douyin/apps/mq/internal/svc"
	"genuine_douyin/common/globalkey"
	"genuine_douyin/common/messageTypes"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

/*
Listening to the payment flow status change notification message queue
*/
type UserFavoriteOpt struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFavoriteUpdateMq(ctx context.Context, svcCtx *svc.ServiceContext) *UserFavoriteOpt {
	return &UserFavoriteOpt{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFavoriteOpt) Consume(_, val string) error {
	var message messageTypes.UserFavoriteOptMessage

	if err := json.Unmarshal([]byte(val), &message); err != nil {
		logx.WithContext(l.ctx).Error("UserFavoriteOptMessage->Consume Unmarshal err : %v , val : %s", err, val)
		return err
	}

	if err := l.execService(message); err != nil {
		logx.WithContext(l.ctx).Error("UserFavoriteOptMessage->execService  err : %v , val : %s , message:%+v", err, val, message)
		logx.Errorf("UserFavoriteOptMessage->execService  err : %v , val : %s , message:%+v", err, val, message)
		return err
	}
	return nil
}

// 处理逻辑
func (l *UserFavoriteOpt) execService(message messageTypes.UserFavoriteOptMessage) error {

	logx.Infof("UserFavoriteOptMessage message : %+v\n", message)

	var req favorite.DouyinFavoriteActionRequest
	_ = copier.Copy(&req, &message)

	// 构造redis的数据
	dataKey := fmt.Sprintf(globalkey.FavoriteSetValTpl, message.VideoId)
	favoriteSetVal := fmt.Sprintf(globalkey.FavoriteSetValTpl, message.VideoId)
	dataVal := fmt.Sprintf(globalkey.ExistDataValTpl, message.UserId, message.ActionType)

	// 消息取出来之后无非是点赞或者取消点赞 0，1，那么打到redis也是0，1
	_, err := l.svcCtx.RedisCache.EvalShaCtx(l.ctx, l.svcCtx.ScriptADD, []string{globalkey.FavoriteSetKey, dataKey}, []string{favoriteSetVal, dataVal})
	if err != redis.Nil {
		logx.Errorf("script exec err : %v", err)
		return err
	}

	return nil
}
