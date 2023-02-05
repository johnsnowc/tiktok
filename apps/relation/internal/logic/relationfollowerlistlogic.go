package logic

import (
	"context"

	"genuine_douyin/apps/relation/internal/svc"
	"genuine_douyin/apps/relation/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationFollowerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRelationFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationFollowerListLogic {
	return &RelationFollowerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RelationFollowerListLogic) RelationFollowerList(in *relation.DouyinRelationFollowerListRequest) (*relation.DouyinRelationFollowerListResponse, error) {
	// todo: add your logic here and delete this line

	return &relation.DouyinRelationFollowerListResponse{}, nil
}
