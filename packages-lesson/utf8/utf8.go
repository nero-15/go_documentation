package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// decodeRune()
	// encodeRune()
	valid()
}

func valid() {
	valid := []byte("Hello, 世界")
	invalid := []byte{0xff, 0xfe, 0xfd}

	fmt.Println(utf8.Valid(valid))
	fmt.Println(utf8.Valid(invalid))
}

func encodeRune() {
	r := '世'
	buf := make([]byte, 3)

	n := utf8.EncodeRune(buf, r)

	fmt.Println(buf)
	fmt.Println(n)
}

func decodeRune() {
	b := []byte("Hello, 世界")

	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		fmt.Printf("%c %v\n", r, size)

		b = b[size:]
	}
}
