package parse

import (
	"strconv"
)

type Int64Parse struct {
}

func (parse *Int64Parse) Parse(text string) (interface{}, error) {
	i, err := strconv.Atoi(text)
	if err != nil {
		return nil, err
	} else {
		result := int64(i)
		return result, nil
	}
}
