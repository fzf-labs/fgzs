package errorx

import (
	"encoding/json"
	"fgzs/internal/errorx/i18n"
	"fgzs/pkg/baiduai"
	"fgzs/pkg/util/fileutil"

	"fmt"
	"sort"
	"strconv"
)

func Export() {
	list := data()
	exportJson(list)
	exportMarkdown(list)
}

// 数据生成
func data() []map[string]string {
	ids := make([]int, 0)
	for i := range BusinessErrs {
		ids = append(ids, i)
	}
	sort.Ints(ids)
	list := make([]map[string]string, 0)
	tranList := make(map[string]map[string]string)
	accessToken := baiduai.GetAccessToken()
	for _, i := range ids {
		m := make(map[string]string)
		m["http_code"] = strconv.Itoa(BusinessErrs[i].GetHttpCode())
		m["code"] = strconv.Itoa(BusinessErrs[i].GetBusinessCode())
		m["err_msg"] = BusinessErrs[i].GetErrMsg()
		zhCNMsg := BusinessErrs[i].GetMessage(i18n.ZhCN)
		m["zh-CN"] = zhCNMsg
		for _, v := range i18n.Languages {
			if v == i18n.ZhCN {
				continue
			}
			m[v] = BusinessErrs[i].GetMessage(v)
			if m[v] == "" {
				if zhCNMsg != "" {
					if _, ok := tranList[zhCNMsg][v]; ok {
						m[v] = tranList[zhCNMsg][v]
					} else {
						message, _ := baiduai.TextTrans(accessToken, zhCNMsg, v)
						_, ok := tranList[zhCNMsg]
						if !ok {
							tranList[zhCNMsg] = make(map[string]string)
						}
						tranList[zhCNMsg][v] = message
						m[v] = message
					}
				}
			}
		}
		list = append(list, m)
	}
	return list
}

// 导出json
func exportJson(list []map[string]string) {
	marshal, err := json.Marshal(list)
	if err != nil {
		return
	}
	err = fileutil.WriteContentCover("./storage/code/code.json", string(marshal))
	if err != nil {
		return
	}
	fmt.Println("错误码json导出成功")
	m := make(map[string]map[string]string)
	for _, v := range list {
		for _, lang := range i18n.Languages {
			if lang == i18n.ZhCN {
				continue
			}
			if _, ok := m[lang]; !ok {
				m[lang] = make(map[string]string)
			}
			m[lang][v["code"]] = v[lang]
		}
	}
	for _, lang := range i18n.Languages {
		if lang == i18n.ZhCN {
			continue
		}
		marshal, err := json.Marshal(m[lang])
		if err != nil {
			return
		}
		err = fileutil.WriteContentCover(fmt.Sprintf("./storage/code/%s.json", lang), string(marshal))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

// 导出markdown
func exportMarkdown(list []map[string]string) {
	first := `|http_code|business_code|err_msg|`
	second := `|--|--|--|`
	for _, v := range i18n.Languages {
		first += v + `|`
		second += `--|`
	}
	str := NewLine(first) + NewLine(second)

	if len(list) > 0 {
		for _, m := range list {
			tmpStr := `|` + m["http_code"] + `|` + m["code"] + `|` + m["err_msg"] + `|`
			for _, v := range i18n.Languages {
				tmpStr += m[v] + `|`
			}
			str += NewLine(tmpStr)
		}
	}

	err := fileutil.WriteContentCover("./storage/code/code.md", str)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("错误码markdown导出成功")
}

func NewLine(str string) string {
	return str + "\n"
}
