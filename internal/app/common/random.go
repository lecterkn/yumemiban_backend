package common

import (
	"math/rand"
	"time"
)

// 生成対象文字列
const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// seed
var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// ランダムなユーザー名を生成
func GenerateUserName() string {
	return random(7)
}

// ランダムなパスワードを生成
func GeneratePassword() string {
	return random(9)
}

// 文字数とcharsetから文字列をランダム生成
func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// 文字数を指定しランダムな文字列を生成
func random(length int) string {
	return stringWithCharset(length, charset)
}
