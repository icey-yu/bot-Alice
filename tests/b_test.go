package tests

import (
	"fmt"
	"testing"
)

func TestB(t *testing.T) {

    x := 5
    y := ~x
    fmt.Println(y) // 输出 -6

    z := int(0)
    fmt.Println(z) // 输出 -1
}
