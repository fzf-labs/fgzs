package urlutil

import (
	"github.com/gookit/goutil/dump"
	"testing"
)

func TestUrlDecode(t *testing.T) {
	str := "username=tizi&password=12345&type=100"
	decode := UrlDecode(str)
	dump.Println(decode)
}

func TestUrlEncode(t *testing.T) {
	m := map[string]string{"username": "tizi", "password": "12345", "type": "100", "a": "1231"}
	encode := UrlEncode(m)
	dump.Println(encode)
}
