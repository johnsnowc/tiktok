package logic

import (
	"context"

	"genuine_douyin/apps/video/rpc/internal/svc"
	"genuine_douyin/apps/video/rpc/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFeedLogic {
	return &GetUserFeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserFeedLogic) GetUserFeed(in *video.DouyinFeedRequest) (*video.DouyinFeedResponse, error) {
	// todo: add your logic here and delete this line

	return &video.DouyinFeedResponse{}, nil
}
