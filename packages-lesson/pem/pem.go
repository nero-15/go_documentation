package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
)

func main() {
	// decode()
	encode()
}

func encode() {
	block := &pem.Block{
		Type: "MESSAGE",
		Headers: map[string]string{
			"Animal": "Gopher",
		},
		Bytes: []byte("test"),
	}

	if err := pem.Encode(os.Stdout, block); err != nil {
		log.Fatal(err)
	}
}

func decode() {
	var pubPEMData = []byte(``)

	block, rest := pem.Decode(pubPEMData)
	if block == nil || block.Type != "PUBLIC KEY" {
		log.Fatal("failed to decode PEM block containing public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Got a %T, with remaining data: %q", pub, rest)
}
