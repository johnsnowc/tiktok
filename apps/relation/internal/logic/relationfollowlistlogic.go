package logic

import (
	"context"

	"genuine_douyin/apps/relation/internal/svc"
	"genuine_douyin/apps/relation/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationFollowListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRelationFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationFollowListLogic {
	return &RelationFollowListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RelationFollowListLogic) RelationFollowList(in *relation.DouyinRelationFollowListRequest) (*relation.DouyinRelationFollowListResponse, error) {
	// todo: add your logic here and delete this line

	return &relation.DouyinRelationFollowListResponse{}, nil
}
