package tests

import (
	"testing"

	"github.com/dlclark/regexp2"
)

func TestRegexp(t *testing.T) {
	match := regexp2.MustCompile(`^(?:very )?good[！|!]?$`, regexp2.IgnoreCase)
	//match := regexp2.MustCompile(`^你?真?(可以|厉害|棒|不错)[！|!]?$`, regexp2.IgnoreCase)
	ok, err := match.MatchString("good")
	if err != nil {
		return
	}
	println(ok)
}
