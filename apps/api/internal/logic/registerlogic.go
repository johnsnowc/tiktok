package logic

import (
	"context"

	"genuine_douyin/apps/api/internal/svc"
	"genuine_douyin/apps/api/internal/types"
	"genuine_douyin/apps/user/user"
	"genuine_douyin/utils/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	res, err := l.svcCtx.UserRpc.Register(l.ctx, &user.DouyinUserRegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		logx.Errorf("register failed: %s", err.Error())
		return &types.RegisterResponse{
			Status: types.Status{
				StatusCode: xerr.SECRET_ERROR,
				StatusMsg:  "register failed" + err.Error(),
			},
		}, nil
	}

	return &types.RegisterResponse{
		Status: types.Status{
			StatusCode: xerr.OK,
		},
		UserId: res.UserId,
		Token:  res.Token,
	}, nil
}
