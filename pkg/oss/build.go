package oss

import (
	"github.com/segmentio/ksuid"
	"strings"
	"time"
)

func BuildNewNameAndPath(ext string, category ...string) (newFileName string, filePath string) {
	//日期
	date := time.Now().Format("20060102")
	join := make([]string, 0)
	join = append(join, category...)
	join = append(join, date)
	dir := strings.Trim(strings.Join(join, "/"), "/")
	newFileName = ksuid.New().String() + ext
	filePath = dir + "/" + newFileName
	return
}
