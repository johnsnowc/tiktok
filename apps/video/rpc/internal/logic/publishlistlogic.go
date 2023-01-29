package logic

import (
	"context"

	"genuine_douyin/apps/video/rpc/internal/svc"
	"genuine_douyin/apps/video/rpc/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishListLogic {
	return &PublishListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishListLogic) PublishList(in *video.DouyinPublishListRequest) (*video.DouyinPublishListResponse, error) {
	//videos, err := l.svcCtx.VideoModel.FindAllByUid(l.ctx, in.UserId)
	//if err != nil {
	//	return &video.DouyinPublishListResponse{}, err
	//}

	return &video.DouyinPublishListResponse{}, nil
}
