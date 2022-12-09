package baiduai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func GetAccessToken() string {
	var host = "https://aip.baidubce.com/oauth/2.0/token"
	var param = map[string]string{
		"grant_type":    "",
		"client_id":     "",
		"client_secret": "",
	}
	uri, err := url.Parse(host)
	if err != nil {
		fmt.Println(err)
	}
	query := uri.Query()
	for k, v := range param {
		query.Set(k, v)
	}
	uri.RawQuery = query.Encode()

	response, err := http.Get(uri.String())
	if err != nil {
		fmt.Println(err)
	}
	result, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	return gjson.Get(string(result), "access_token").String()
}

type TextTransResp struct {
	Result struct {
		TransResult []struct {
			Dst string `json:"dst"`
			Src string `json:"src"`
		} `json:"trans_result"`
		From string `json:"from"`
		To   string `json:"to"`
	} `json:"result"`
	LogId int64 `json:"log_id"`
}

// 韩语 kor
// 繁体 cht
// 英语 en
// 有频率限制
func TextTrans(accessToken string, q string, to string) (string, error) {
	time.Sleep(time.Second)
	langTo := map[string]string{
		"zh-TW": "cht", // zh_TW 繁体中文-中国
		"en-US": "en",  // en_US 英文-美国
		"ko-KR": "kor", // ko_KR 朝鲜语-韩国
	}
	var host = "https://aip.baidubce.com/rpc/2.0/mt/texttrans/v1"
	uri, err := url.Parse(host)
	if err != nil {
		return "", err
	}
	query := uri.Query()
	query.Set("access_token", accessToken)
	uri.RawQuery = query.Encode()
	param := map[string]string{
		"from": "auto",
		"to":   langTo[to],
		"q":    q,
	}
	jsonByte, err := json.Marshal(param)
	if err != nil {
		return "", err
	}
	client := &http.Client{}
	request, err := http.NewRequest("POST", uri.String(), bytes.NewBuffer(jsonByte))
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", "application/json;charset=utf-8")
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body)

	result, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	fmt.Println(string(result))
	return gjson.Get(string(result), "result.trans_result.0.dst").String(), nil
}

type TransResult struct {
	Dst string `json:"dst"`
	Src string `json:"src"`
}

func TextTransBySli(accessToken string, qs []string, to string) (map[string]string, error) {
	time.Sleep(time.Second)
	resp := make(map[string]string)
	if len(qs) == 0 || to == "" {
		return resp, fmt.Errorf("%s", "参数错误")
	}
	langTo := map[string]string{
		"zh-TW": "cht", // zh_TW 繁体中文-中国
		"en-US": "en",  // en_US 英文-美国
		"ko-KR": "kor", // ko_KR 朝鲜语-韩国
	}
	var host = "https://aip.baidubce.com/rpc/2.0/mt/texttrans/v1"
	uri, err := url.Parse(host)
	if err != nil {
		return nil, err
	}
	q := ""
	for _, v := range qs {
		q += v + "\n"
	}
	q = strings.TrimRight(q, "\n")
	query := uri.Query()
	query.Set("access_token", accessToken)
	uri.RawQuery = query.Encode()
	param := map[string]string{
		"from": "auto",
		"to":   langTo[to],
		"q":    q,
	}
	jsonByte, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	request, err := http.NewRequest("POST", uri.String(), bytes.NewBuffer(jsonByte))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json;charset=utf-8")
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body)

	result, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	transResults := make([]TransResult, 0)
	s := gjson.Get(string(result), "result.trans_result").String()
	if s == "" {
		return nil, err
	}
	err = json.Unmarshal([]byte(s), &transResults)
	if err != nil {
		return nil, err
	}
	for _, transResult := range transResults {
		resp[transResult.Src] = transResult.Dst
	}
	return resp, nil
}
