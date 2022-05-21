package tool

import (
	"crypto/sha256"
	"encoding/hex"
)

func Token(str string) string {
	data := sha256.Sum256([]byte(str))
	return hex.EncodeToString(data[:])
}
