package main

import (
	"crypto/sha512"
	"fmt"
)

func main() {
	h := sha512.New()
	h.Write([]byte("hello world\n"))
	fmt.Printf("%x", h.Sum(nil))
}
