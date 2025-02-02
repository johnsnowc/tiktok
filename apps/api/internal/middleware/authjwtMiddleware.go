package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"genuine_douyin/apps/api/internal/types"
	myToken "genuine_douyin/utils/jwt"
	"genuine_douyin/utils/xerr"
)

type AuthJWTMiddleware struct {
}

func NewAuthJWTMiddleware() *AuthJWTMiddleware {
	return &AuthJWTMiddleware{}
}

/*
	这里前端有bug 用户信息、投稿接口、发布列表应该都带有user_id 才能对比token解析出来的user_id是否一致
	但是前端的投稿接口没有user_id 只有token 所以这里没有办法判断
*/

func (m *AuthJWTMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status := new(types.Status)
		token := r.FormValue("token")

		if token == "" {
			status.StatusCode = xerr.REUQEST_PARAM_ERROR
			status.StatusMsg = "no token"
			res, _ := json.Marshal(status)
			_, _ = w.Write(res)
			return
		}
		// 解析token 判断是否有效
		var parseClaims myToken.ParseToken
		claims, err := parseClaims.ParseToken(token)
		if err != nil {
			status.StatusCode = xerr.REUQEST_PARAM_ERROR
			status.StatusMsg = "param error " + err.Error()
			res, _ := json.Marshal(status)
			_, _ = w.Write(res)
			return
		}

		id := r.FormValue("user_id")
		var userId int64
		if id != "" {
			res, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				status.StatusCode = xerr.REUQEST_PARAM_ERROR
				status.StatusMsg = fmt.Sprintf("param error user_id : %d", res)
				res, _ := json.Marshal(status)
				_, _ = w.Write(res)
				return
			}
			userId = res
		}

		if userId != 0 && userId != claims.UserId {
			status.StatusCode = xerr.REUQEST_PARAM_ERROR
			status.StatusMsg = "user_id not match"
			res, _ := json.Marshal(status)
			_, _ = w.Write(res)
			return
		}

		// 过期时间点 小于当前时间 表示过期
		if claims.ExpireAt < time.Now().Unix() {
			status.StatusCode = xerr.REUQEST_PARAM_ERROR
			status.StatusMsg = "please login again"
			res, _ := json.Marshal(status)
			_, _ = w.Write(res)
			return
		}
		r = r.Clone(context.WithValue(r.Context(), myToken.CurrentUserId("CurrentUserId"), claims.UserId))

		next(w, r)
	}
}
