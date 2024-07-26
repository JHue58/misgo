package clipboard

import "fmt"

var ClipEmptyError = fmt.Errorf("clip不能为空")
var ContentEmptyError = fmt.Errorf("content不能为空")
var FileToLargeError = fmt.Errorf("文件过大")
var FileNameOrHashError = fmt.Errorf("content(filename)或hash不正确")
