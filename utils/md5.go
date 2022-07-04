package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 进行md5编码
func MD5(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}
