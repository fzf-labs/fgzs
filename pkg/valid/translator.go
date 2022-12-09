package valid

import (
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func RegisterTranslation(v *Validator) error {
	if err := v.validator.RegisterTranslation("required_if", v.trans, registerTranslator("required_if", "{0}未传"), translate); err != nil {
		return err
	}

	if err := v.validator.RegisterTranslation("phone", v.trans, registerTranslator("phone", "{0}错误的手机号"), translate); err != nil {
		return err
	}

	if err := v.validator.RegisterTranslation("dateGt", v.trans, registerTranslator("dateGt", "{0}必须要大于当前日期"), translate); err != nil {
		return err
	}
	if err := v.validator.RegisterTranslation("dateLt", v.trans, registerTranslator("dateLt", "{0}必须要小于当前日期"), translate); err != nil {
		return err
	}
	return nil
}

// removeTopStruct 去除字段名中的结构体名称标识
// refer from:https://github.com/go-playground/validator/issues/633#issuecomment-654382345
func RemoveTopStruct(fields map[string]string) string {
	var res string
	for field, err := range fields {
		res += field[strings.Index(field, ".")+1:] + ":" + err + ";"
	}
	res = strings.TrimRight(res, ";") + "."
	return res
}

// registerTranslator 为自定义字段添加翻译功能
func registerTranslator(tag string, msg string) validator.RegisterTranslationsFunc {
	return func(trans ut.Translator) error {
		if err := trans.Add(tag, msg, false); err != nil {
			return err
		}
		return nil
	}
}

// translate 自定义字段的翻译方法
func translate(trans ut.Translator, fe validator.FieldError) string {
	msg, err := trans.T(fe.Tag(), fe.Field())
	if err != nil {
		panic(fe.(error).Error())
	}
	return msg
}
