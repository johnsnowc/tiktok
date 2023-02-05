package logic

import (
	"context"
	"time"

	"genuine_douyin/apps/user/user"
	"genuine_douyin/apps/video/dal"
	"genuine_douyin/apps/video/internal/svc"
	"genuine_douyin/apps/video/video"

	"github.com/jinzhu/copier"
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
	var timeline time.Time
	if in.LatestTime == nil {
		timeline = time.Now()
	} else {
		timeline = time.Unix(*in.LatestTime, 0)
	}

	list := make([]*dal.Video, 0)
	if err := l.svcCtx.VideoModel.WithContext(l.ctx).Limit(30).Order("create_time desc").
		Where("create_time < ?", timeline).Find(&list).Error; err != nil {
		return nil, err
	}

	returnList := make([]*video.Video, 0, len(list))
	for i := 0; i < len(list); i++ {
		u, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &user.DouyinUserRequest{
			UserId: list[i].Uid,
		})
		if err != nil {
			l.Logger.Error("request user rpc failed")
		}
		var v video.Video
		_ = copier.Copy(&v, list[i])
		v.Author = u.User
		//todo favorite
		returnList = append(returnList, &v)
	}
	nextTime := list[len(list)-1].CreatedAt.Unix()

	return &video.DouyinFeedResponse{
		VideoList: returnList,
		NextTime:  &nextTime,
	}, nil
}
