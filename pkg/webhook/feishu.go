package webhook

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type FeiShuConfig struct {
	Url  string
	Sign string
}

type FeiShu struct {
	cfg *FeiShuConfig
}

func NewFeiShu(cfg *FeiShuConfig) *FeiShu {
	return &FeiShu{cfg: cfg}
}

func NewFeiShuByCfg(cfg *FeiShuConfig) *FeiShu {
	return &FeiShu{cfg: cfg}
}

type SendMsg struct {
	Timestamp string `json:"timestamp"`
	Sign      string `json:"sign"`
	MsgType   string `json:"msg_type"`
	Content   struct {
		Text string `json:"text"`
	} `json:"content"`
}

// https://open.feishu.cn/open-apis/bot/v2/hook/790d95f6-47f8-45b0-8cfe-7635dec6c1d4
// i25mfl9o1MT2vpUOkCtEZc
func (f *FeiShu) SendMsg(msg string) error {
	timestamp := time.Now().Unix()
	sign, err := f.GenSign(f.cfg.Sign, timestamp)
	if err != nil {
		return err
	}
	sendMsg := SendMsg{
		Timestamp: strconv.FormatInt(timestamp, 10),
		Sign:      sign,
		MsgType:   "text",
		Content: struct {
			Text string `json:"text"`
		}{
			Text: msg,
		},
	}
	marshal, err := json.Marshal(sendMsg)
	if err != nil {
		return err
	}
	// request
	result, err := http.Post(f.cfg.Url, "application/json", strings.NewReader(string(marshal)))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return err
	}
	defer result.Body.Close()
	return nil
}

func (f *FeiShu) GenSign(secret string, timestamp int64) (string, error) {
	//timestamp + key 做sha256, 再进行base64 encode
	stringToSign := fmt.Sprintf("%v", timestamp) + "\n" + secret

	var data []byte
	h := hmac.New(sha256.New, []byte(stringToSign))
	_, err := h.Write(data)
	if err != nil {
		return "", err
	}
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature, nil
}
