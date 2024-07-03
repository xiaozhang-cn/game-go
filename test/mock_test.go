package test

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")

	panic("panic error!")
}
