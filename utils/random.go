package utils

import "math/rand"

func RandomString(l int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes := make([]byte, l)
	for i := range bytes {
		bytes[i] = chars[rand.Intn(l)%len(chars)]
	}
	return string(bytes)
}
