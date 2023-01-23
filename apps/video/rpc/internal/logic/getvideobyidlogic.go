package logic

import (
	"context"

	"genuine_douyin/apps/video/rpc/internal/svc"
	"genuine_douyin/apps/video/rpc/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoByIdLogic {
	return &GetVideoByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVideoByIdLogic) GetVideoById(in *video.VideoIdRequest) (*video.Video, error) {
	// todo: add your logic here and delete this line

	return &video.Video{}, nil
}
