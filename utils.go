package main

import (
	"log"
	"math/rand"
)

const dictionary = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateRandomUppercaseString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = dictionary[rand.Intn(len(dictionary))]
	}
	return string(b)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
