package parse

// ResourceParse 解析器
type ResourceParse interface {
	// Parse 解析成值
	Parse(text string) (interface{}, error)
}
