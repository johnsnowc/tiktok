package jwt

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	UserId   int64 `json:"user_id"`
	ExpireAt int64 `json:"expire_at"`
	Else     jwt.MapClaims
	jwt.RegisteredClaims
}

type CurrentUserId string

const (
	AccessSecret = "genuine_douyin"
	AccessExpire = 86400
)
