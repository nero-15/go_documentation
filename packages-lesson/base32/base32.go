package main

import (
	"encoding/base32"
	"fmt"
	"os"
)

func main() {
	// newEncoder()
	// decode()
	encode()
}

func encode() {
	data := []byte("any + old & data")
	str := base32.StdEncoding.EncodeToString(data)
	fmt.Println(str)
}

func decode() {
	str := "ONXW2ZJAMRQXIYJAO5UXI2BAAAQGC3TEEDX3XPY="
	data, err := base32.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("%q\n", data)
}

func newEncoder() {
	input := []byte("foo\x00bar")
	encoder := base32.NewEncoder(base32.StdEncoding, os.Stdout)
	encoder.Write(input)
	// Must close the encoder when finished to flush any partial blocks.
	// If you comment out the following line, the last partial block "r"
	// won't be encoded.
	encoder.Close()
}
