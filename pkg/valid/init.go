package valid

import (
	"fmt"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/pkg/errors"
	"reflect"
	"strings"
	"sync"
)

type Validator struct {
	once      sync.Once
	validator *validator.Validate
	trans     ut.Translator
}

var Validate = &Validator{}

func (v *Validator) NewValidator() *Validator {
	v.once.Do(func() {
		v.validator = validator.New()
		// 注册一个获取json tag的自定义方法
		v.validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		// 在校验器注册自定义的校验方法
		if err := RegisterValidation(v.validator); err != nil {
			panic(err)
		}
	})
	return v
}

func (v *Validator) RegisterTranslation(locale string) *Validator {
	zhT := zh.New() // 中文翻译器
	enT := en.New() // 英文翻译器

	// 第一个参数是备用（fallback）的语言环境
	// 后面的参数是应该支持的语言环境（支持多个）
	// uni := ut.New(zhT, zhT) 也是可以的
	uni := ut.New(enT, zhT, enT)

	// locale 通常取决于 http 请求头的 'Accept-Language'
	var ok bool
	// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
	v.trans, ok = uni.GetTranslator(locale)
	if !ok {
		panic(fmt.Sprintf("Validator GetTranslator(%s) failed", locale))
	}
	var err error
	// 注册翻译器
	switch locale {
	case "en":
		err = enTranslations.RegisterDefaultTranslations(v.validator, v.trans)
	case "zh":
		err = zhTranslations.RegisterDefaultTranslations(v.validator, v.trans)
	default:
		err = enTranslations.RegisterDefaultTranslations(v.validator, v.trans)
	}
	if err != nil {
		panic(fmt.Sprintf("Validator RegisterDefaultTranslations(%s) failed", locale))
	}
	// 注意！因为要使用到trans实例
	// 所以这一步注册要放到trans初始化的后面
	if err := RegisterTranslation(v); err != nil {
		panic(fmt.Sprintf("Validator RegisterTranslation(%s) failed", locale))
	}
	return v
}

// Validate validate
func (v *Validator) Validate(obj interface{}) error {
	if obj == nil {
		return nil
	}
	value := reflect.ValueOf(obj)
	switch value.Kind() {
	case reflect.Ptr:
		err := v.validator.Struct(value.Elem().Interface())
		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				return errors.New(err.Translate(v.trans))
			}
			return err
		}
	case reflect.Struct:
		err := v.validator.Struct(obj)
		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				return errors.New(err.Translate(v.trans))
			}
			return err
		}
	case reflect.Slice, reflect.Array:
		count := value.Len()
		for i := 0; i < count; i++ {
			if err := v.validator.Struct(value.Index(i).Interface()); err != nil {
				for _, err := range err.(validator.ValidationErrors) {
					return errors.New(err.Translate(v.trans))
				}
				return err
			}
		}
	default:
		return nil
	}
	return nil
}
