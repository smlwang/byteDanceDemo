package tool

import (
	"crypto/sha256"
	"encoding/hex"
)

//使用sha256生成token
func Token(str string) string {
	data := sha256.Sum256([]byte(str))
	return hex.EncodeToString(data[:])
}
