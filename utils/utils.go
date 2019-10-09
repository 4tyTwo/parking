package utils

import (
	"log"
	"math/rand"
)

const dictionary = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GenerateRandomString - generates random string of n length
func GenerateRandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = dictionary[rand.Intn(len(dictionary))]
	}
	return string(b)
}

// CheckErr does nothing if err is nil, calls log.Fatal otherwise
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
