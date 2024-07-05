package utils

import (
	"fmt"
	"reflect"
	"testing"
)

type Point struct {
	X, Y int
}

func TestUtils(t *testing.T) {

	var a int32 = 10
	typeOfA := reflect.TypeOf(a)
	log.Info("type:", typeOfA)

	p1 := Point{X: 1, Y: 2}
	p2 := Point{X: 1, Y: 2}
	p3 := Point{X: 2, Y: 3}

	fmt.Println(p1 == p2) // true
	fmt.Println(p1 == p3) // false
}
