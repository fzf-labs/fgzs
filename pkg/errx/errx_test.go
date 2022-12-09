package errx

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	stack := callersStack(2, 1)
	s1 := stack.String()

	fmt.Println(s1)
}
