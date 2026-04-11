package auth

import (
	"errors"
	"strconv"

	"github.com/golang-jwt/jwt"
)

// ErrTokenInvalidClaims is returned when the token claims are invalid
var ErrTokenInvalidClaims = errors.New("token invalid claims")

// @secretKey: JWT 加解密密钥
// @iat: 时间戳
// @seconds: 过期时间，单位秒
// @payload: 数据载体
func GetJwtToken(secretKey string, iat, seconds int64, payload string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["payload"] = payload
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

// ParseToken 解析 JWT Token
func ParseToken(tokenString string, secretKey string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalidClaims
}

// GetUserIdFromPayload 从 JWT payload 中获取用户 ID
func GetUserIdFromPayload(payload string) (int64, error) {
	userId, err := strconv.ParseInt(payload, 10, 64)
	if err != nil {
		return 0, err
	}
	return userId, nil
}