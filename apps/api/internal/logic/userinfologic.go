package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"genuine_douyin/apps/api/internal/svc"
	"genuine_douyin/apps/api/internal/types"
	"genuine_douyin/apps/user/user"
	myToken "genuine_douyin/utils/jwt"
	"genuine_douyin/utils/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	fromId := l.ctx.Value(myToken.CurrentUserId("CurrentUserId")).(int64)
	if fromId == 0 {
		logx.Errorf("parse token from id == 0")
	}
	res, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &user.DouyinUserRequest{
		UserId: req.UserId,
		FromId: fromId,
	})
	if err != nil {
		logx.Errorf("get user info failed: %v", err.Error())
		return &types.UserInfoResponse{
			Status: types.Status{
				StatusCode: xerr.ERR,
				StatusMsg:  "get user info failed",
			},
		}, nil
	}

	var u types.User
	_ = copier.Copy(&u, res.User)

	return &types.UserInfoResponse{
		Status: types.Status{
			StatusCode: xerr.OK,
		},
		User: u,
	}, nil
}
