package utils

import "unicode"

func Capitalize(s string) string {
	if s == "" {
		return s
	}
	// 获取首字母
	firstRune := []rune(s)[0]
	// 将首字母转换为大写
	firstRune = unicode.ToUpper(firstRune)
	// 返回首字母大写后的字符串
	return string(firstRune) + s[1:]
}
