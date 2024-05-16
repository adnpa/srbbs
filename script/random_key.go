package main

//生成随机密钥 用于jwt

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func main() {
	// 生成 256 位的随机密钥
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		fmt.Println("Error generating random key:", err)
		return
	}

	// 将密钥转换为 base64 编码的字符串
	keyBase64 := base64.StdEncoding.EncodeToString(key)
	fmt.Println("Random key (base64):", keyBase64)
}
