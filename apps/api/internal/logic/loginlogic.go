package logic

import (
	"context"

	"genuine_douyin/apps/api/internal/svc"
	"genuine_douyin/apps/api/internal/types"
	"genuine_douyin/apps/user/user"
	"genuine_douyin/utils/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Login 用户登陆
// 通过username获得用户密码，然后比对密码
// token 先查redis 是否存在，如果存在，则直接返回，如果不存在，则生成token，并存入redis
// 并返回userId，token
func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	res, err := l.svcCtx.UserRpc.Login(l.ctx, &user.DouyinUserRegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		logx.Errorf("login failed: %v", err.Error())
		return &types.LoginResponse{
			Status: types.Status{
				StatusCode: xerr.ERR,
				StatusMsg:  "login failed",
			},
		}, nil
	}

	return &types.LoginResponse{
		Status: types.Status{
			StatusCode: xerr.OK,
		},
		UserId: res.UserId,
		Token:  res.Token,
	}, nil
}
