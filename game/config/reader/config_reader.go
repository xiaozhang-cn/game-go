package reader

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"test2/common"
	"test2/utils"
)

var log = utils.GetLog()

func init() {
	fmt.Println("init reader")
}

func NewConfigHead() common.ConfigHead {
	return common.ConfigHead{
		Heads: make([]string, 0),
		Data:  make([][]string, 0),
	}
}

func ReadExcel(filepath string) common.ConfigHead {
	log.Debug("read file: ", filepath)
	f, err := excelize.OpenFile(filepath)
	if err != nil {
		log.Error(err)
	}

	// 关闭文件
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
		log.Debug("close file: ", filepath)
	}()

	configHead := NewConfigHead()
	sheetList := f.GetSheetList()
	if len(sheetList) == 0 {
		log.Warnf("no sheets found")
	}

	firstSheet := sheetList[0]
	rows, err := f.GetRows(firstSheet)
	if err != nil {
		log.Error("read sheet error", err)
	}

	configHead.Heads = rows[0][:]
	log.Info("配置:", f.Path, " 表头:", configHead.Heads)

	for _, row := range rows[1:] {
		lineData := row[:]
		appendLength := len(configHead.Heads) - len(lineData)
		// 剩余列数据用""填充
		if appendLength > 0 {
			appendSlice := make([]string, appendLength)
			lineData = append(lineData, appendSlice...)
		}
		configHead.Data = append(configHead.Data, lineData)
	}

	return configHead
}
