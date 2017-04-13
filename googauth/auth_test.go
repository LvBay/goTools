package googauth_test

import (
	"goTools/googauth"
	"strconv"
	"testing"
	"time"
)

func TestAuth(t *testing.T) {

	secret := googauth.CreateSecret("test1") // 密钥
	t0 := int64(time.Now().Unix() / 30)
	code := googauth.ComputeCode(secret, t0) // 根据密钥和时间戳生成code

	// 配置
	otpconf := &googauth.OTPConfig{
		Secret:     secret,
		WindowSize: 3,
	}
	result, _ := otpconf.Authenticate(strconv.Itoa(code)) // 验证code
	t.Log(result)
}
