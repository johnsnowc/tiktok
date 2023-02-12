package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"genuine_douyin/apps/api/internal/types"
	myToken "genuine_douyin/utils/jwt"
	"genuine_douyin/utils/xerr"
)

type IsLoginMiddleware struct {
}

func NewIsLoginMiddleware() *IsLoginMiddleware {
	return &IsLoginMiddleware{}
}

// Handle 没有登录也是可以调用feed接口的
func (m *IsLoginMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status := new(types.Status)
		token := r.FormValue("token")
		if token != "" {
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

			// 过期时间点 小于当前时间 表示过期
			if claims.ExpireAt < time.Now().Unix() {
				status.StatusCode = xerr.REUQEST_PARAM_ERROR
				status.StatusMsg = "please login again"
				res, _ := json.Marshal(status)
				_, _ = w.Write(res)
				return
			}

			r = r.Clone(context.WithValue(r.Context(), myToken.CurrentUserId("LoginUserId"), claims.UserId))

			// 把r传递给下一个handler
		}
		next(w, r)
	}
}
