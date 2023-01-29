package logic

import (
	"context"

	"genuine_douyin/apps/user/rpc/internal/svc"
	"genuine_douyin/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
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
	// todo: add your logic here and delete this line

	return &user.DouyinUserResponse{}, nil
}
