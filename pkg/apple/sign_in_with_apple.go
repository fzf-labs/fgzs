package apple

import (
	"context"
	"fgzs/pkg/util/strutil"
	"fmt"
	"github.com/Timothylock/go-signin-with-apple/apple"
	"github.com/pyihe/apple_validator"
)

//https://blog.csdn.net/qq_36770474/article/details/118340500

type SignInWithAppleConfig struct {
	Secret      string
	KeyId       string
	TeamId      string
	ClientId    string
	RedirectUrl string
}

type SignInWithApple struct {
	Cfg *SignInWithAppleConfig
}

func NewSignInWithApple(cfg *SignInWithAppleConfig) *SignInWithApple {
	return &SignInWithApple{Cfg: cfg}
}

// 苹果验证都两种方式:
// 1.验证jwt 需要去服务端获取公钥,然后来验证格式是否正确
// 2.验证authorizationCode 直接请求校验接口
func (s *SignInWithApple) CheckByAuthorizationCode(code string) (uniqueID string, err error) {
	privateKey := strutil.FormatPrivateKey(s.Cfg.Secret)
	// 生成用于向 Apple 验证服务器进行身份验证的客户端密码
	clientSecret, err := apple.GenerateClientSecret(privateKey, s.Cfg.TeamId, s.Cfg.ClientId, s.Cfg.KeyId)
	if err != nil {
		return "", fmt.Errorf("error generating secret: " + err.Error())
	}
	// 生成新的验证客户端
	client := apple.New()
	vReq := apple.AppValidationTokenRequest{
		ClientID:     s.Cfg.ClientId,
		ClientSecret: clientSecret,
		Code:         code,
	}
	var resp apple.ValidationResponse
	// 进行验证
	err = client.VerifyAppToken(context.Background(), vReq, &resp)
	if err != nil {
		return "", fmt.Errorf("error verifying: " + err.Error())
	}
	if resp.Error != "" {
		return "", fmt.Errorf("apple returned an error: %s - %s\n", resp.Error, resp.ErrorDescription)
	}
	// Get the unique user ID
	unique, err := apple.GetUniqueID(resp.IDToken)
	if err != nil {
		return "", fmt.Errorf("failed to get unique ID: " + err.Error())
	}
	return unique, nil
}

func (s *SignInWithApple) CheckIdentityToken(token string) (uniqueID string, err error) {
	validator := apple_validator.NewValidator()
	jwtToken, err := validator.CheckIdentityToken(token)
	if err != nil {
		return "", fmt.Errorf("CheckIdentityToken err: " + err.Error())
	}
	ok, err := jwtToken.IsValid()
	if !ok {
		return "", fmt.Errorf("CheckIdentityToken IsValid err: " + err.Error())
	}
	return jwtToken.Sub(), nil
}
