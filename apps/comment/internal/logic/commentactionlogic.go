package logic

import (
	"context"

	"genuine_douyin/apps/comment/comment"
	"genuine_douyin/apps/comment/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentActionLogic {
	return &CommentActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentActionLogic) CommentAction(in *comment.DouyinCommentActionRequest) (*comment.DouyinCommentActionResponse, error) {
	// todo: add your logic here and delete this line

	return &comment.DouyinCommentActionResponse{}, nil
}
