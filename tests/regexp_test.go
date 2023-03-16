package tests

import (
	"github.com/dlclark/regexp2"
	"testing"
)

func TestRegexp(t *testing.T) {
	match := regexp2.MustCompile(`^(?:very )?good[！|!]?$`, regexp2.IgnoreCase)
	//match := regexp2.MustCompile(`^你?真?(可以|厉害|棒|不错)[！|!]?$`, regexp2.IgnoreCase)
	ok, err := match.MatchString("GOOD")
	if err != nil {
		return
	}
	println(ok)
}
