package utils

import (
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func ToJson(object interface{}) (string, error) {
	jsonData, err := json.Marshal(object)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil

}
