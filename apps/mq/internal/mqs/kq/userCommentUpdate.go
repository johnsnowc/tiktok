package kq

import (
	"context"
	"encoding/json"
	"genuine_douyin/apps/comment/comment"
	"genuine_douyin/apps/mq/internal/svc"
	"genuine_douyin/common/messageTypes"
	"github.com/zeromicro/go-zero/core/logx"
)

/*
Listening to the payment flow status change notification message queue
*/
type UserCommentOpt struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCommentUpdateMq(ctx context.Context, svcCtx *svc.ServiceContext) *UserCommentOpt {
	return &UserCommentOpt{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCommentOpt) Consume(_, val string) error {
	var message messageTypes.UserCommentOptMessage
	if err := json.Unmarshal([]byte(val), &message); err != nil {
		logx.WithContext(l.ctx).Error("UserCommentOptMessage->Consume Unmarshal err : %v , val : %s", err, val)
		return err
	}

	if err := l.execService(message); err != nil {
		logx.WithContext(l.ctx).Error("UserCommentOptMessage->execService  err : %v , val : %s , message:%+v", err, val, message)
		return err
	}

	return nil
}

// 处理逻辑
func (l *UserCommentOpt) execService(message messageTypes.UserCommentOptMessage) error {
	// 调用rpc 更新user_comment表
	_, err := l.svcCtx.CommentSvcRpcClient.CommentAction(l.ctx, &comment.DouyinCommentActionRequest{
		FromId:      message.UserId,
		VideoId:     message.VideoId,
		UserId:      message.UserId,
		CommentId:   &message.CommentId,
		CommentText: &message.CommentText,
		ActionType:  int32(message.ActionType),
	})

	logx.Error("UserCommentOptMessage->execService xxxxxxxxxxx")

	if err != nil {
		logx.Errorf("UserCommentOptMessage->execService  err : %v , val : %s , message:%+v", err, message)
		return err
	}

	return nil
}
