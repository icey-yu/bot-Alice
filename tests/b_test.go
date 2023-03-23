package tests

import (
	"testing"
)

func TestB(t *testing.T) {
	ch := make(chan int, 100)
	ch <- 1
	println(len(ch))
}
