package parse

type StringParse struct {
}

func (parse *StringParse) Parse(text string) (interface{}, error) {
	return text, nil
}
