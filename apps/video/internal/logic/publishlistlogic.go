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
	list := make([]*dal.Video, 0)
	if err := l.svcCtx.VideoModel.WithContext(l.ctx).Where("uid = ?", in.UserId).Find(&list).Error; err != nil {
		return &video.DouyinPublishListResponse{}, err
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
		// todo favorite
		returnList = append(returnList, &v)
	}

	return &video.DouyinPublishListResponse{VideoList: returnList}, nil
}
