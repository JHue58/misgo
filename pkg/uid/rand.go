package uid

import (
	"crypto/rand"
	"encoding/base64"
)

var r = randManager{length: 64}

func NewUid() (string, error) {
	return r.Generate()
}

type randManager struct {
	length int
}

func (r randManager) Generate() (Uid string, err error) {
	// 每个字符需要6个比特（base64的每个字符代表6个比特），所以我们需要的字节数是 length * 6 / 8。
	// 这里我们使用 base64.URLEncoding 字符集。
	bytes := make([]byte, (r.length*6+7)/8)

	_, err = rand.Read(bytes)
	if err != nil {
		return "", err
	}

	// 使用 base64.URLEncoding 编码字节数组
	encoded := base64.URLEncoding.EncodeToString(bytes)

	// 截取到指定长度
	if r.length > len(encoded) {
		return encoded, nil
	}
	return encoded[:r.length], nil
}
