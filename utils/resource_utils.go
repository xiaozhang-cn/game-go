package utils

import "reflect"
import "test2/game/config/parse"

type FieldType struct {
	Type string
}

// 类型解析器
var type2Parser = make(map[FieldType]parse.ResourceParse)

func ParseResourceBean(nameFiled reflect.Value, fieldString map[string]string) {
	createInstance := CreateInstance(nameFiled.Type())
	FillObjFieldValue(fieldString, &createInstance)
}
