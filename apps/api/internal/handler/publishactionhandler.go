package handler

import (
	"net/http"

	"genuine_douyin/apps/api/internal/logic"
	"genuine_douyin/apps/api/internal/svc"
	"genuine_douyin/apps/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PublishActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PublishActionRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewPublishActionLogic(r.Context(), svcCtx)
		resp, err := l.PublishAction(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
