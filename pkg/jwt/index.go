package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const SecretKey = "SecretKey"

type Claims struct {
	Data map[string]interface{}
	jwt.RegisteredClaims
}

// CreateJWTToken 生成token.
// data: 数据
// durationSecond: 有效期，单位秒
func CreateToken(data map[string]interface{}, durationSecond uint64) (string, error) {
	claims := Claims{
		data,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(durationSecond) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			// Issuer:    "test",
			// Subject:   "somebody",
			// ID:        time.Now().GoString(),
			// Audience:  []string{"somebody_else"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(SecretKey))
	return tokenString, err
}

// ParseToken 解析token.
func ParseToken(token string) (data map[string]interface{}, err error) {
	var tokenInfo *jwt.Token

	if tokenInfo, err = jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(SecretKey), err
	}); err != nil {
		return
	}
	if claims, ok := tokenInfo.Claims.(*Claims); ok && tokenInfo.Valid {
		data = claims.Data
	} else {
		err = errors.New("parsing data exception")
	}
	return
}

// RefreshToken 刷新token有效期.
func RefreshToken(token string) (data map[string]interface{}, err error) {
	// TODO
	return
}
