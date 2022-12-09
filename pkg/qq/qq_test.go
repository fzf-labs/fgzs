package qq

import (
	"fmt"
	"testing"
)

func TestAuthQq_GetUserInfo(t *testing.T) {
	authQq := NewAuthQq(&AuthConfig{
		ClientId:     "",
		ClientSecret: "",
		RedirectUrl:  "",
		ApplyUnionId: "1",
	})
	token := ""
	userInfo, err := authQq.GetUserInfo(token)
	fmt.Println(err)
	if err != nil {
		return
	}
	fmt.Println(userInfo)
}
