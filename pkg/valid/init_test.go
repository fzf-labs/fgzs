package valid

import (
	"fmt"
	"testing"
)

func TestValidate(t *testing.T) {
	type Man struct {
		Name    string `json:"name" validate:"required"`
		Version string `json:"version" validate:"required"`
		Phone   string `json:"phone" validate:"phone"`
		Info    string `json:"info"`
	}
	man := Man{
		Version: "121",
		Info:    "12312",
	}
	err := Validate.NewValidator().RegisterTranslation("zh").Validate(man)

	fmt.Println(err)
}
