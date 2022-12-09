package errorx

import (
	"encoding/json"
	"fgzs/internal/errorx/i18n"
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"strings"
)

var BusinessErrs = make(map[int]*BusinessErr)

// H is a shortcut for map[string]interface{}
type H map[string]interface{}

const (
	InfoLevel = "info"
	WarnLevel = "warn"
	ErrLevel  = "err"
)

type BusinessErr struct {
	HttpCode         int           `json:"httpCode"`         // HTTP Code
	BusinessCode     int           `json:"businessCode"`     // 业务错误码
	ErrMessage       string        `json:"errMessage"`       // 错误信息
	ErrMessageFormat []interface{} `json:"errMessageFormat"` // 错误信息格式化参数
	ErrLevel         string        `json:"errLevel"`         // 错误等级
	ErrDetail        string        `json:"errDetail"`        // 错误细节
	ErrData          interface{}   `json:"errData"`          // 错误内容
}

func (e *BusinessErr) Error() string {
	marshal, _ := json.Marshal(e)
	return string(marshal)
}

// NewError 实例化一个错误
func NewError(businessCode int, errMsg string, opts ...Option) *BusinessErr {
	_, ok := BusinessErrs[businessCode]
	if ok {
		panic(fmt.Sprintf("code %d is exsit, please change one", businessCode))
	}
	e := &BusinessErr{
		HttpCode:     http.StatusOK,
		BusinessCode: businessCode,
		ErrMessage:   errMsg,
		ErrLevel:     InfoLevel,
	}
	if len(opts) > 0 {
		for _, f := range opts {
			f(e)
		}
	}
	BusinessErrs[businessCode] = e
	return e
}

type Option func(e *BusinessErr)

func WithHttpCode(httpCode int) Option {
	return func(e *BusinessErr) {
		e.HttpCode = httpCode
	}
}

func WithLevel(level string) Option {
	return func(e *BusinessErr) {
		e.ErrLevel = level
	}
}
func (e *BusinessErr) WithErrData(errData interface{}) *BusinessErr {
	newError := *e
	newError.ErrData = errData
	return &newError
}
func (e *BusinessErr) WithCustomMsg(msg string) *BusinessErr {
	newError := *e
	newError.ErrMessage = msg
	return &newError
}

func (e *BusinessErr) WithMsgFormat(msgFormat ...interface{}) *BusinessErr {
	newError := *e
	newError.ErrMessageFormat = msgFormat
	return &newError
}

// WithDetail 错误细节
func (e *BusinessErr) WithDetail(err error) *BusinessErr {
	newError := *e
	newError.ErrDetail = fmt.Sprintf("err: %s,line: %s", err.Error(), e.fileLine())
	return &newError
}

func (e *BusinessErr) GetHttpCode() int {
	return e.HttpCode
}

func (e *BusinessErr) GetBusinessCode() int {
	return e.BusinessCode
}

func (e *BusinessErr) GetErrMsg() string {
	return e.ErrMessage
}

func (e *BusinessErr) GetMessage(lang string) string {
	message := i18n.GetMessage(e.GetBusinessCode(), lang)
	if message != "" && len(e.ErrMessageFormat) > 0 && strings.Contains(message, "%") {
		message = fmt.Sprintf(message, e.ErrMessageFormat...)
	}
	return message
}

func (e *BusinessErr) GetDetail() string {
	return e.ErrDetail
}

func (e *BusinessErr) GetErrData() interface{} {
	if e.ErrData == nil {
		return H{}
	} else {
		return e.ErrData
	}
}

func (e *BusinessErr) fileLine() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return ""
	}
	//fcName := runtime.FuncForPC(caller - 1).Name()
	return file + ":" + strconv.Itoa(line)
}
