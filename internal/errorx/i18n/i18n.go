package i18n

import "strconv"

const (
	ZhCN = "zh-CN" // zh_CN 简体中文-中国
	ZhTW = "zh-TW" // zh_TW 繁体中文-中国
	EnUS = "en-US" // en_US 英文-美国
	KoKR = "ko-KR" // ko_KR 朝鲜语-韩国
)

var Languages = []string{
	ZhCN,
	ZhTW,
	EnUS,
	KoKR,
}

func GetMessage(code int, lang string) string {
	var msg string
	switch lang {
	case ZhCN:
		msg = ZhCNMap[strconv.Itoa(code)]
	case ZhTW:
		msg = ZhTWMap[strconv.Itoa(code)]
	case KoKR:
		msg = KoKRMap[strconv.Itoa(code)]
	case EnUS:
		msg = EnUSMap[strconv.Itoa(code)]
	default:
		msg = ZhCNMap[strconv.Itoa(code)]
	}
	return msg
}
