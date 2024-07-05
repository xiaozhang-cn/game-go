package parse

import (
	"strconv"
)

type Int32Parse struct {
}

func (parse *Int32Parse) Parse(text string) (interface{}, error) {
	i, err := strconv.Atoi(text)
	if err != nil {
		return nil, err
	} else {
		result := int32(i)
		return result, nil
	}
}
