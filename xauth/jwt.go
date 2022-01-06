package xauth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

// GetJWT 生成 jwt token
func GetJWT(key string, val map[string]interface{}, expire int) (string, error) {
	val["iat"] = time.Now().Unix()
	val["exp"] = time.Now().Add(time.Duration(expire) * time.Second).Unix()
	return jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims(val),
	).SignedString([]byte(key))
}

// ParseJWT 解析 jwt token
func ParseJWT(key, token string) (map[string]interface{}, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}
	if !t.Valid {
		return nil, errors.New("token does not valid")
	}
	val, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("token claims does not valid")
	}
	return val, nil
}
