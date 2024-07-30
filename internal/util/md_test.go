package util

import (
	"fmt"
	"testing"
)

func TestMd(t *testing.T) {
	val := "## why \n |Title|Time|\n|-|-|\n|123|as|"
	md, _ := ToMarkDownHtml(val)
	fmt.Println(md)
}
