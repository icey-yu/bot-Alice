package tests

import (
	"strings"
	"testing"
)

func TestStr(t *testing.T) {
	//s := "@问遥 我有问题想问问你"
	s := "我有问题想@问遥 问问你"
	//exS := strings.Trim(s, "@问遥")
	exS := strings.Replace(s, "@问遥", "", -1)
	println(exS)
}
