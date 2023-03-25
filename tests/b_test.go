package tests

import (
	"testing"
)

func TestB(t *testing.T) {
	li := make([]int, 0)
	add(li)
	println(li)
}

func add(li []int) []int {
	li = append(li, 5)
	return li
}
