package test

import (
	"fmt"
	"test2/game/config/parse"
	"testing"
)

func TestDefer(t *testing.T) {
	parse := parse.Int32Parse{}
	i, err := parse.Parse("1-")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(i)
}
