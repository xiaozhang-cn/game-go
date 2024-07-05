package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"test2/common"
	"test2/game/config/bean"
	"test2/game/config/reader"
	"test2/utils"
)

const (
	BeanPath = "./game/config/bean"
)

var log = utils.GetLog()

func main() {
	excel2Bean()
}

func excel2Bean() {
	dirPath := "resource/excel"
	dir, err := os.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}

	configMap := make(map[string]common.ConfigHead)
	for _, file := range dir {
		fileAbPath := filepath.Join(dirPath, file.Name())
		excel := reader.ReadExcel(fileAbPath)
		excelJson, _ := utils.ToJson(excel)
		log.Debug("excel text: ", excelJson)
		var configResource = strings.TrimSuffix(file.Name(), filepath.Ext(file.Name())) + "_" + "resource"
		configMap[strings.ToLower(configResource)] = excel
	}

	var resourceBeanHolder = bean.GetInstance()

	pkg := reflect.TypeOf(resourceBeanHolder)
	elem := reflect.ValueOf(&resourceBeanHolder).Elem()
	//log.Info(pkg)

	fieldNum := elem.NumField()
	for i := 0; i < fieldNum; i++ {
		nameFiled := elem.Field(i)
		fieldName := pkg.Field(i).Name
		if !strings.Contains(fieldName, "Resource") {
			continue
		}
		fieldValue := nameFiled.Interface()
		jsonString, _ := utils.ToJson(fieldValue)
		log.WithFields(logrus.Fields{
			"fieldIndex": i,
			"fieldName":  fieldName,
			"fieldValue": jsonString,
		})

		var resourceName = strings.ToLower(fieldName)
		parts := strings.Split(resourceName, "resource")
		resourceName = strings.Join(parts, "_resource")
		resourceName = strings.ToLower(resourceName)

		configHead := configMap[resourceName]
		for index := range configHead.Data {
			fieldString := configHead.GetFieldString(index)
			utils.ParseResourceBean(nameFiled, fieldString)
		}
	}

}
