package main

import (
	"crypto/aes"
	"fmt"
)

func main() {
	key := []byte("abcdefghijklmnopqrstuvwxyz123456")
	c, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	fmt.Println(c)
}
