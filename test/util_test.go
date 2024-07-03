package test

import (
	"fmt"
	"test2/common"
	"test2/game/config/bean"
	"test2/utils"
	"testing"
)

func TestUtils(t *testing.T) {
	slice := make([]int, 0)
	var p1 = slice
	var p2 = slice

	fmt.Printf("p1 address: %p\n", &p1)
	fmt.Printf("p2 address: %p\n", &p2)

	//utils.GetLog().Info(utils.IsMemoryEquals(p1, p2))

	resource := bean.SkillResource{}
	var resourceInterface interface{} = resource
	utils.FillObjFieldValue(common.ConfigHead{}, &resourceInterface)
}
