package logic

import (
	"context"

	"genuine_douyin/apps/relation/internal/svc"
	"genuine_douyin/apps/relation/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRelationActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationActionLogic {
	return &RelationActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RelationActionLogic) RelationAction(in *relation.DouyinRelationActionRequest) (*relation.DouyinRelationActionResponse, error) {
	// todo: add your logic here and delete this line

	return &relation.DouyinRelationActionResponse{}, nil
}
