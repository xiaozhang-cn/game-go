package common

type ConfigHead struct {
	Heads []string   // 表头
	Data  [][]string // 所有行数据
}

func (configHead ConfigHead) GetFieldString(lineIndex int) map[string]string {
	lineLength := len(configHead.Data)
	if lineIndex > lineLength-1 {
		return nil
	}

	fieldString := make(map[string]string)
	for i, field := range configHead.Heads {
		fieldString[field] = configHead.Data[lineIndex][i]
	}
	return fieldString
}
