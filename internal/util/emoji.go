package util

import "regexp"

var emojiRegx = regexp.MustCompile(`[\p{So}\p{C}]`) // \p{So} 代表符号表情，\p{C} 代表其他控制字符

func PopEmojis(input string) string {
	// 替换所有匹配的 emoji 表情为空字符串
	return emojiRegx.ReplaceAllString(input, "")
}
