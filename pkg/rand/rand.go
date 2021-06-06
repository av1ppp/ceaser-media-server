package rand

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var symbolRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_")

// Generation a random string with a specified length.
func RandString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = symbolRunes[rand.Intn(len(symbolRunes))]
	}
	return string(b)
}
