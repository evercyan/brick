package xauth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
{
     "iss": "jwt",         // 签发者
     "iat": 1595838971,    // 签发时间
     "exp": 1595838972,    // 过期时间
     "nbf": 1595838972,    // 校验时间, 该时间前此 token 无效
     "sub": "www.xxx.com", // 面向用户
     "jti": "xxxx",        // 该 token 唯一标识
     "xxx": "xxx",         // 可附加信息
}
*/

var (
	key     = "sdflIerl34i^flkj"
	expire  = 24 * 3600
	payload = map[string]interface{}{
		"name": "hello",
	}
)

func TestJwt(t *testing.T) {
	{
		encrypted, err := GetJWT(key, payload, expire)
		assert.Nil(t, err)
		assert.NotEmpty(t, encrypted)

		decrypted, err := ParseJWT(key, encrypted)
		assert.Nil(t, err)
		assert.NotNil(t, decrypted)
		assert.Equal(t, "hello", decrypted["name"])

	}
	{
		_, err := ParseJWT(key, "abcdefg")
		assert.NotNil(t, err)
	}
}
