package util

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

type Hash []byte

func (h Hash) String() string {
	return hex.EncodeToString(h)
}

type HashBuilder interface {
	MD5() Hash
	SHA1() Hash
	SHA256() Hash
}

type hashBuilder struct {
	data []byte
}

func NewHashBuilder(data []byte) HashBuilder {
	return &hashBuilder{data: data}
}

// MD5 computes the MD5 hash of the data
func (h *hashBuilder) MD5() Hash {
	hash := md5.Sum(h.data)
	return hash[:]
}

// SHA1 computes the SHA-1 hash of the data and returns the first 128 bits
func (h *hashBuilder) SHA1() Hash {
	hash := sha1.Sum(h.data)
	return hash[:] // First 128 bits
}

// SHA256 computes the SHA-256 hash of the data
func (h *hashBuilder) SHA256() Hash {
	hash := sha256.Sum256(h.data)
	return hash[:]
}
