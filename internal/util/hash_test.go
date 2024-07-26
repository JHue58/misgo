package util

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	str := "asdasfafsdfjghfhfha"
	hash := NewHashBuilder([]byte(str))
	fmt.Printf("md5:%s\nsha1:%s\nnsha256:%s", hash.MD5(), hash.SHA1(), hash.SHA256())
}
