package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"genuine_douyin/apps/user/dal"
	"genuine_douyin/apps/user/internal/svc"
	"genuine_douyin/apps/user/user"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIdLogic) GetUserById(in *user.DouyinUserRequest) (*user.DouyinUserResponse, error) {
	u := dal.User{}
	if err := l.svcCtx.UserModel.WithContext(l.ctx).Where("id = ?", in.UserId).First(&u).Error; err != nil {
		return &user.DouyinUserResponse{}, status.Error(100, "用户不存在")
	}

	//todo 判断from_id是否follow了user_id
	var res user.User
	_ = copier.Copy(&res, u)

	return &user.DouyinUserResponse{User: &res}, nil
}
