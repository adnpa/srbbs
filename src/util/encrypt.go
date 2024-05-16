package util

import (
	"crypto/sha256"
	"fmt"
)

// EncryptPassword 加密用户密码
func EncryptPassword(data []byte) (res string) {
	h := sha256.New()
	h.Write(data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

//func encryptPassword(data []byte) (result string) {
//	h := md5.New()
//	h.Write([]byte(secret))
//	return hex.EncodeToString(h.Sum(data))
//}
