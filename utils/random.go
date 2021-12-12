package utils

import (
	"math/rand"
	"time"
)

func RandomString(n int) string{
	var letters = []byte("abcdefghijklmnopqrstuvwxyz")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for k, _ := range result {
		result[k] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
