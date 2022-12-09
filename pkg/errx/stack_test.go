package errx

import (
	"fmt"
	"runtime"
	"testing"
)

func TestFunc_FileLine(t *testing.T) {
	//str := "abc"
	caller, file, line, ok := runtime.Caller(0)
	fmt.Println(caller, file, line, ok)
}
