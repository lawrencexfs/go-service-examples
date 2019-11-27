package mytest

import (
	"testing"
)

func Test_test1(t *testing.T) {
	a := []int{0, 1, 2, 3, 4}
	func(b []int) {
		b[1] = 1000
	}(a)

	if a[1] != 1000 {
		t.Error("test1 error!")
	}
	t.Log("a[1] = ", a[1])
}

func Test_test2(t *testing.T) {
	a := make(map[int]int)
	func(b map[int]int) {
		b[1] = 1000
	}(a)

	if val, ok := a[1]; !ok || val != 1000 {
		t.Error("test2 error!")
	}
	t.Log("a[1] = ", a[1])
}
