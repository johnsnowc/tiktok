package logic

import (
	"context"

	"genuine_douyin/apps/favorite/favorite"
	"genuine_douyin/apps/favorite/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteListLogic) FavoriteList(in *favorite.DouyinFavoriteListRequest) (*favorite.DouyinFavoriteListResponse, error) {
	// todo: add your logic here and delete this line

	return &favorite.DouyinFavoriteListResponse{}, nil
}
