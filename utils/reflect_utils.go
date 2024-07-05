package utils

import (
	"reflect"
)

func FillObjFieldValue(fieldString map[string]string, object interface{}) {
	// 获取接口的反射值
	val := reflect.ValueOf(object)

	for {
		if val.Kind() == reflect.Interface { // 是interface先取值
			val = val.Elem()
		} else if val.Kind() == reflect.Ptr { // 获取指针指向的值
			val = val.Elem()
		} else {
			break
		}
	}

	// 获取实际类型
	typ := val.Type()
	// 输出结构体的类型
	log.Debug("obj type: ", typ.Name())

}

func CreateInstance(t reflect.Type) interface{} {
	return reflect.New(t).Elem().Interface()
}
