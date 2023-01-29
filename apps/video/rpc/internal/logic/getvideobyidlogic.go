package logic

import (
	"context"

	"genuine_douyin/apps/user/rpc/user"
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
	v, err := l.svcCtx.VideoModel.FindOne(l.ctx, in.VideoId)
	if err != nil {
		return &video.Video{}, err
	}

	u, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &user.DouyinUserRequest{
		UserId: v.Uid,
	})
	if err != nil {
		return &video.Video{}, err
	}

	return &video.Video{
		Id: v.Id,
		Author: &user.User{
			Id:            u.User.Id,
			Name:          u.User.Name,
			FollowCount:   u.User.FollowCount,
			FollowerCount: u.User.FollowerCount,
			IsFollow:      u.User.IsFollow,
		},
		PlayUrl:       v.PlayUrl,
		CoverUrl:      v.CoverUrl,
		FavoriteCount: v.FavoriteCount,
		CommentCount:  v.CommentCount,
		//IsFavorite:    , todo
		Title: v.Title,
	}, nil
}
