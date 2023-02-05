package logic

import (
	"context"

	"genuine_douyin/apps/favorite/favorite"
	"genuine_douyin/apps/favorite/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteActionLogic {
	return &FavoriteActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteActionLogic) FavoriteAction(in *favorite.DouyinFavoriteActionRequest) (*favorite.DouyinFavoriteActionResponse, error) {
	// todo: add your logic here and delete this line

	return &favorite.DouyinFavoriteActionResponse{}, nil
}
