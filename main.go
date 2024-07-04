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

var log = utils.GetLog()

func main() {
	var packageName = "./game/config/bean"
	types, err := utils.ExtractStructTypes(packageName)
	if err != nil {
		log.Error(err)
	}
	log.Info("scan types: ", types)
	testExcel2Bean()
}

func testExcel2Bean() {
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
		log.Info("configResource: ", configResource)

		configMap[strings.ToLower(configResource)] = excel
	}

	var resourceBeanHolder = bean.GetInstance()

	// 获取包的类型
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
			createInstance := utils.CreateInstance(nameFiled.Type())
			utils.FillObjFieldValue(fieldString, &createInstance)
		}
	}

}

func getResourceBeanName(filename string) {

}
