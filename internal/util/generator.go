package util

import "math/rand"

func GenerateRandomString(n int) string {
	var charsets = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	var letters = make([]rune, n)

	for i := range letters {
		letters[i] = charsets[rand.Intn(len(charsets))]
	}

	return string(letters)
}
