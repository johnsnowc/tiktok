package logic

import (
	"context"
	"strconv"
	"time"

	"genuine_douyin/apps/user/dal"
	"genuine_douyin/apps/user/internal/svc"
	"genuine_douyin/apps/user/user"
	"genuine_douyin/utils/cryptx"
	"genuine_douyin/utils/jwt"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.DouyinUserRegisterRequest) (*user.DouyinUserRegisterResponse, error) {
	logx.Infof("i'm in")
	u := dal.User{}
	if err := l.svcCtx.UserModel.WithContext(l.ctx).Where("name = ?", in.Username).First(&u).Error; err == nil {
		return &user.DouyinUserRegisterResponse{}, errors.New("user exists")
	}

	passWord := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
	newUser := dal.User{
		Name:     in.Username,
		Password: passWord,
	}
	if err := l.svcCtx.UserModel.WithContext(l.ctx).Create(&newUser).Error; err != nil {
		return &user.DouyinUserRegisterResponse{}, err
	}

	var genToken *jwt.GenToken
	now := time.Now()
	tokenString, err := genToken.GenToken(now, newUser.Id, nil)
	if err != nil {
		logx.Errorf("gen token error: %s", err.Error())
		return &user.DouyinUserRegisterResponse{}, errors.Wrapf(err, "genToken error")
	}

	_, err = l.svcCtx.RedisCache.SetnxExCtx(l.ctx, "token:"+strconv.FormatInt(newUser.Id, 10), tokenString, jwt.AccessExpire)
	if err != nil {
		logx.Errorf("set token to redis error: %s", err.Error())
		return &user.DouyinUserRegisterResponse{}, errors.Wrapf(err, "set redis token error")
	}

	return &user.DouyinUserRegisterResponse{
		UserId: newUser.Id,
		Token:  tokenString,
	}, nil
}
