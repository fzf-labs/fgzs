package meta

import (
	"context"
	"encoding/json"
	"fgzs/internal/define/constant"
	"fgzs/internal/define/vars"
	"fgzs/pkg/conv"
	"fgzs/pkg/util/urlutil"
	"net/http"
)

// GetUid  从Context中获取用户ID
func GetUid(ctx context.Context) string {
	return conv.String(ctx.Value(vars.ContextWithValueKey(constant.ContextUID)))
}

// SetUid 设置用户ID到Context
func SetUid(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, vars.ContextWithValueKey(constant.ContextUID), value)
}

// GetXParams  从Context中获取header自定义参数
func GetXParams(ctx context.Context) *XParams {
	xParams := conv.String(ctx.Value(vars.ContextWithValueKey(constant.ContextXParams)))
	return TransformXParams(xParams)
}

func SetXParams(ctx context.Context, value *XParams) context.Context {
	return context.WithValue(ctx, vars.ContextWithValueKey(constant.ContextXParams), value)
}

// GetXParamsByHeader 从Context中获取header自定义参数
func GetXParamsByHeader(r *http.Request) *XParams {
	xParams := r.Header.Get(constant.HeaderXParams)
	return TransformXParams(xParams)
}

type XParams struct {
	Language   string `json:"language"`    //语言 例如: zh-CN
	AppVersion string `json:"app_version"` //app版本号 例如：1.0.0
	DeviceType string `json:"device_type"` //设备类型 例如: ios  android
	DeviceID   string `json:"device_id"`   //设备ID 例如：123456789
	Tz         string `json:"tz"`          //时区 例如：Asia/Shanghai
	Time       string `json:"time"`        //时间 例如：1657598308
	Platform   string `json:"platform"`    //平台 例如：ios  android h5 cocos
	ChannelID  string `json:"channel_id"`  //渠道ID 例如：default
}

func TransformXParams(str string) *XParams {
	params := new(XParams)
	urlDecode := urlutil.UrlDecode(str)
	marshal, err := json.Marshal(urlDecode)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(marshal, params)
	if err != nil {
		return nil
	}
	return params
}

func (x *XParams) Check() bool {
	return true
}
