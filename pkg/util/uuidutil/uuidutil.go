package uuidutil

import (
	"github.com/google/uuid"
	"github.com/segmentio/ksuid"
	"github.com/teris-io/shortid"
	"strings"
	"time"
)

// GenUUID 生成随机字符串，eg: 76d27e8c-a80e-48c8-ad20-e5562e0f67e4
func GenUUID() string {
	u, _ := uuid.NewRandom()
	return u.String()
}

// GenShortID 生成一个id
func GenShortID() (string, error) {
	return shortid.Generate()
}

func KSUid() string {
	return ksuid.New().String()
}

func AssetKey() string {
	s, _ := ksuid.NewRandomWithTime(time.Now())
	return strings.ToUpper(s.String())
}
