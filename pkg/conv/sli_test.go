package conv

import (
	"github.com/gookit/goutil/dump"
	"testing"
)

func TestSliToMap(t *testing.T) {
	sli1 := []string{"1", "2", "3"}

	sliToMap := SliToMap(sli1, func(s string) (string, string) {
		return s, s
	})
	dump.Println(sliToMap)
}
