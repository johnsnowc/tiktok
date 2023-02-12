package logic

import (
	"context"
	myToken "genuine_douyin/utils/jwt"
	"genuine_douyin/utils/xerr"
	"strconv"
	"time"

	"genuine_douyin/apps/user/dal"
	"genuine_douyin/apps/user/internal/svc"
	"genuine_douyin/apps/user/user"
	"genuine_douyin/utils/cryptx"

	"github.com/pkg/errors"
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
	// 查询用户是否存在
	u := dal.User{}
	if err := l.svcCtx.UserModel.WithContext(l.ctx).Where("name = ?", in.Username).First(&u).Error; err != nil {
		logx.Errorf("find user failed, err: %s", err.Error())
		return &user.DouyinUserRegisterResponse{}, errors.Wrap(err, "find user failed")
	}

	// 判断密码是否正确
	password := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
	logx.Infof("in.password: %s, after crypt: %s, u.password: %s", in.Password, password, u.Password)
	if password != u.Password {
		logx.Errorf("password not match")
		return &user.DouyinUserRegisterResponse{}, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "password not match")
	}

	// 通过userId查 redis 是否有此token
	token, err := l.svcCtx.RedisCache.GetCtx(l.ctx, "token:"+strconv.FormatInt(u.Id, 10))
	if err != nil {
		logx.Errorf("get token from redis failed, err: %s", err.Error())
		return &user.DouyinUserRegisterResponse{}, errors.Wrap(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "get token from redis failed")
	}
	// 如果存在，则直接返回
	if token != "" {
		return &user.DouyinUserRegisterResponse{
			UserId: u.Id,
			Token:  token,
		}, nil
	}

	//如果不存在，则生成token，并存入redis
	var genToken myToken.GenToken
	now := time.Now()
	token, err = genToken.GenToken(now, u.Id, nil)
	_, err = l.svcCtx.RedisCache.SetnxExCtx(l.ctx, "token:"+strconv.FormatInt(u.Id, 10), token, myToken.AccessExpire)
	if err != nil {
		logx.Errorf("set token to redis failed, err: %s", err.Error())
		return &user.DouyinUserRegisterResponse{}, errors.Wrapf(xerr.NewErrCode(xerr.TOKEN_GENERATE_ERROR), "set token to redis error")
	}

	return &user.DouyinUserRegisterResponse{
		UserId: u.Id,
		Token:  token,
	}, nil
}
