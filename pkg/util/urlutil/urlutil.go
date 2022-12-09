package urlutil

import "net/url"

func UrlEncode(m map[string]string) string {
	if len(m) == 0 {
		return ""
	}
	param := url.Values{}
	for k, v := range m {
		param.Add(k, v)
	}
	unescape, err := url.QueryUnescape(param.Encode())
	if err != nil {
		return ""
	}
	return unescape
}

func UrlDecode(str string) map[string]string {
	values, err := url.ParseQuery(str)
	if err != nil {
		return nil
	}
	m := make(map[string]string)
	for k, v := range values {
		m[k] = v[0]
	}
	return m
}
