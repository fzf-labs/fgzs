package ddm

import (
	"fmt"
	"strings"
)

func Mobile(m string) string {
	if len(m) != 11 {
		return ""
	}
	return fmt.Sprintf("%s****%s", m[:3], m[len(m)-4:])
}

func BankCard(bc string) string {
	if len(bc) > 19 || len(bc) < 16 {
		return ""
	}
	return fmt.Sprintf("%s******%s", bc[:6], bc[len(bc)-4:])
}

func IDCard(s string) string {
	if len(s) != 18 {
		return ""
	}

	return fmt.Sprintf("%s******%s", s[:1], s[len(s)-1:])
}

func IDName(name string) string {
	if len(name) < 1 {
		return ""
	}
	nameRune := []rune(name)
	return fmt.Sprintf("*%s", string(nameRune[1:]))
}

func PassWord() string {
	return "******"
}

func Email(e string) string {
	if !strings.Contains(string(e), "@") {
		return ""
	}
	split := strings.Split(string(e), "@")
	if len(split[0]) < 1 || len(split[1]) < 1 {
		return ""
	}
	return fmt.Sprintf("%s***%s", split[0][:1], split[0][len(split[0])-1:])
}
