package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type GenToken struct{}

func (g *GenToken) GenToken(iat time.Time, userId int64, payloads map[string]interface{}) (string, error) {
	claims := Claims{
		UserId:   userId,
		ExpireAt: iat.Add(time.Second * AccessExpire).Unix(),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "johnsnowc",
		},
	}
	for k, v := range payloads {
		claims.Else[k] = v
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(AccessSecret))
}
