package xauth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

// GetJWT 生成 jwt token
func GetJWT(key string, claims map[string]interface{}, expire int) (string, error) {
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Duration(expire) * time.Second).Unix()
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(claims)).SignedString([]byte(key))
}

// ParseJWT 解析 jwt token
func ParseJWT(key, token string) (claims map[string]interface{}, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("error parsing JWT: %v", e)
		}
	}()
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}
	if !t.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	return claims, nil
}
