package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/api/parser"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var DirName = flag.String("d", "", "")

func main() {
	flag.Parse()
	dir, err := FilePathWalkDir(*DirName)
	if err != nil {
		return
	}
	all := make([]Api, 0)
	if len(dir) > 0 {
		for _, v := range dir {
			apis := ParseContent(v)
			all = append(all, apis...)
		}
	}
	if len(all) > 0 {
		sql := "INSERT INTO system_apis (api_group,`method`,`path`,description) VALUES"
		for _, api := range all {
			sql += fmt.Sprintf("('%s','%s','%s','%s'),", api.ApiGroup, api.Method, api.Path, api.Description)
		}
		err := WriteWithIo("./api.sql", strings.Trim(sql, ","))
		if err != nil {
			return
		}
	}
}

type Api struct {
	ApiGroup    string `json:"api_group"`
	Path        string `json:"path"`
	Description string `json:"description"`
	Method      string `json:"method"`
}

func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(path, ".api") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
func ParseContent(api string) []Api {
	apis := make([]Api, 0)
	sp, err := parser.Parse(api)
	if err != nil {
		return nil
	}

	if len(sp.Service.Groups) > 0 {
		for _, group := range sp.Service.Groups {
			groupName := group.Annotation.Properties["group"]
			if len(group.Routes) > 0 {
				for _, route := range group.Routes {
					method := route.Method
					path := route.Path
					handlerComment := route.HandlerComment
					comment := ""
					if len(handlerComment) > 0 {
						comment = handlerComment[0]
					}
					a := Api{
						ApiGroup:    groupName,
						Path:        path,
						Description: strings.TrimSpace(strings.TrimLeft(comment, "/")),
						Method:      strings.ToUpper(method),
					}
					apis = append(apis, a)
				}
			}
		}
	}
	return apis
}

// WriteWithIo 使用io.WriteString()函数进行数据的写入，不存在则创建
func WriteWithIo(filePath, content string) error {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return err
	}
	defer file.Close()

	if content != "" {
		_, err := io.WriteString(file, content)
		if err != nil {
			return err
		}
	}
	return nil
}
