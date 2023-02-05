package logic

import (
	"context"

	"genuine_douyin/apps/comment/comment"
	"genuine_douyin/apps/comment/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentListLogic) CommentList(in *comment.DouyinCommentListRequest) (*comment.DouyinCommentListResponse, error) {
	// todo: add your logic here and delete this line

	return &comment.DouyinCommentListResponse{}, nil
}
