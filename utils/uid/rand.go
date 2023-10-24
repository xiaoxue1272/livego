package uid

import "math/rand"

var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		// rand 伪随机 seed 固定 1
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
