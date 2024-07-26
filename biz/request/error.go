package request

import "fmt"

type errorPrefix string

func (c errorPrefix) joinMsg(msg string) string {
	return fmt.Sprintf(string(c), msg)
}

var (
	parmaError    errorPrefix = "参数错误: %s"
	dataBaseError errorPrefix = "数据库错误: %s"
	serverError   errorPrefix = "服务器错误: %s"
)
