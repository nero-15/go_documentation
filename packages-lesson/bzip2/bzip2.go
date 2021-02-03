package main

import (
	"compress/bzip2"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	text := "BZh91AY&SYA\xaf\x82\r\x00\x00\x01\x01\x80\x02\xc0\x02\x00 \x00!\x9ah3M\x07<]\xc9\x14\xe1BA\x06\xbe\x084"
	st := strings.NewReader(text)
	bz := bzip2.NewReader(st)
	_, err := io.Copy(os.Stdout, bz)
	if err != nil {
		fmt.Println(err)
	}

}
