package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//practice bytes
	//https://golang.org/pkg/bufio/#example_Scanner_Bytes
	bytes()
}

func bytes() {
	scanner := bufio.NewScanner(strings.NewReader("gopher"))
	for scanner.Scan() {
		fmt.Println(len(scanner.Bytes()) == 6)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "shouldn't see an error scanning a string")
	}
}
