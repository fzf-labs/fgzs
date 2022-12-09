package baiduai

import (
	"fmt"
	"testing"
)

func TestTextTrans(t *testing.T) {
	accessToken := GetAccessToken()
	trans, err := TextTransBySli(accessToken, []string{"你好", "测试"}, "ko-KR")
	if err != nil {
		return
	}
	fmt.Println(trans)
}
