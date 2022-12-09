package baiduai

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// BaiduImageAIConfig 图片审核配置（百度）
type BaiduImageAIConfig struct {
	ClientID     string
	ClientSecret string
}

// Token 访问配置
type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

// GetToken 获取百度审核接口的访问token
func GetToken(clientID, clientSecret string) (*Token, error) {
	url := fmt.Sprintf("https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id=%s&client_secret=%s", clientID, clientSecret)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var token Token
	if err = json.Unmarshal(body, &token); err != nil {
		return nil, err
	}
	return &token, nil
}

// AIResult 审核结果
type AIResult struct {
	LogID          int64 `json:"log_id"`         // 请求唯一ID
	ErrorCode      int64 `json:"error_code"`     // 错误码
	ConclusionType int32 `json:"conclusionType"` // 审核结果1:合规,2:不合规,3:疑似,4:审核失败
}

// GetAIResult 获取审核结果
func GetAIResult(accessToken string, fileBytes []byte) (*AIResult, error) {
	var host = "https://aip.baidubce.com/rest/2.0/solution/v1/img_censor/v2/user_defined"

	uri, err := url.Parse(host)
	if err != nil {
		fmt.Println(err)
	}
	query := uri.Query()
	query.Set("access_token", accessToken)
	uri.RawQuery = query.Encode()

	image := base64.StdEncoding.EncodeToString(fileBytes)
	sendBody := http.Request{}
	err = sendBody.ParseForm()
	if err != nil {
		return nil, err
	}
	sendBody.Form.Add("image", image)
	sendData := sendBody.Form.Encode()
	client := &http.Client{}
	request, err := http.NewRequest("POST", uri.String(), strings.NewReader(sendData))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)
	result, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var aiResult AIResult
	if err = json.Unmarshal(result, &aiResult); err != nil {
		return nil, err
	}
	return &aiResult, nil
}
