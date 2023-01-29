package logic

import (
	"context"

	"genuine_douyin/apps/user/rpc/internal/svc"
	"genuine_douyin/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.DouyinUserRegisterRequest) (*user.DouyinUserRegisterResponse, error) {
	// todo: add your logic here and delete this line

	return &user.DouyinUserRegisterResponse{}, nil
}
