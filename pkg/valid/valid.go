package valid

import (
	"fgzs/pkg/util/validutil"
	"github.com/go-playground/validator/v10"
	"time"
)

// RegisterValidation 在校验器注册自定义的校验方法
func RegisterValidation(v *validator.Validate) error {
	if err := v.RegisterValidation("required_if", RequiredIf); err != nil {
		return err
	}
	if err := v.RegisterValidation("phone", Phone); err != nil {
		return err
	}
	if err := v.RegisterValidation("dateGt", DateGt); err != nil {
		return err
	}
	if err := v.RegisterValidation("dateLt", DateLt); err != nil {
		return err
	}
	return nil
}
func RequiredIf(fl validator.FieldLevel) bool {

	return true
}

// Phone 校验是否是手机号
func Phone(fl validator.FieldLevel) bool {
	if phone, ok := fl.Field().Interface().(string); ok {
		return validutil.IsPhone(phone)
	}
	return false
}

// DateGt 大于日期
func DateGt(fl validator.FieldLevel) bool {
	date, err := time.Parse("2006-01-02", fl.Field().String())
	if err != nil {
		return false
	}
	if date.Before(time.Now()) {
		return false
	}
	return true
}

// DateLt 大于日期
func DateLt(fl validator.FieldLevel) bool {
	date, err := time.Parse("2006-01-02", fl.Field().String())
	if err != nil {
		return false
	}
	if date.After(time.Now()) {
		return false
	}
	return true
}

// 校验银行卡
func CheckBankCard(s string) bool {
	return checkLuHn(s)
}

// Luhn算法
// Luhn算法会通过校验码对一串数字进行验证，校验码通常会被加到这串数字的末尾处，从而得到一个完整的身份识别码
func checkLuHn(value string) bool {
	var (
		sum     = 0
		nDigits = len(value)
		parity  = nDigits % 2
	)
	for i := 0; i < nDigits; i++ {
		var digit = int(value[i] - 48)
		if i%2 == parity {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}
	return sum%10 == 0
}
