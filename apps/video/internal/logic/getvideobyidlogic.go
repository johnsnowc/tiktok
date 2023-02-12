package logic

import (
	"context"

	"genuine_douyin/apps/user/user"
	"genuine_douyin/apps/video/dal"
	"genuine_douyin/apps/video/internal/svc"
	"genuine_douyin/apps/video/video"

	"github.com/jinzhu/copier"
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
	v := dal.Video{}
	if err := l.svcCtx.VideoModel.WithContext(l.ctx).Where("id = ?", in.VideoId).First(&v).Error; err != nil {
		return nil, err
	}

	u, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &user.DouyinUserRequest{
		UserId: v.Uid,
	})
	if err != nil {
		return &video.Video{}, err
	}

	var res video.Video
	_ = copier.Copy(&res, v)
	res.Author = u.User
	//todo favorite
	return &res, nil
}
