package main

import (
	"fmt"
	"strconv"
)

func main() {
	// append()
	// atoi()
	// canBackquote()
	// format()
	// isPrint()
	// itoa()
	parse()
}

func parse() {
	//bool
	v := "true"
	if s, err := strconv.ParseBool(v); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}

	//int
	v32 := "-354634382"
	if s, err := strconv.ParseInt(v32, 10, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	if s, err := strconv.ParseInt(v32, 16, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}

	v64 := "-3546343826724305832"
	if s, err := strconv.ParseInt(v64, 10, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	if s, err := strconv.ParseInt(v64, 16, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
}

func itoa() {
	i := 10
	s := strconv.Itoa(i)
	fmt.Printf("%T, %v\n", s, s)

}

func isPrint() {
	c := strconv.IsPrint('\u263a')
	fmt.Println(c)

	bel := strconv.IsPrint('\007')
	fmt.Println(bel)
}

func format() {
	//bool
	// v := true
	// s := strconv.FormatBool(v)
	// fmt.Printf("%T, %v\n", s, s)

	//int
	v := int64(-42)

	s10 := strconv.FormatInt(v, 10)
	fmt.Printf("%T, %v\n", s10, s10)

	s16 := strconv.FormatInt(v, 16)
	fmt.Printf("%T, %v\n", s16, s16)
}

func canBackquote() {
	fmt.Println(strconv.CanBackquote("Fran & Freddie's Diner ☺"))
	fmt.Println(strconv.CanBackquote("`can't backquote this`"))
}

func atoi() {
	v := "10"
	if s, err := strconv.Atoi(v); err == nil {
		fmt.Printf("%T, %v", s, s)
	}
}

func append() {
	//bool
	b := []byte("bool:")
	b = strconv.AppendBool(b, true)
	fmt.Println(string(b))

	//float
	b32 := []byte("float32:")
	b32 = strconv.AppendFloat(b32, 3.1415926535, 'E', -1, 32)
	fmt.Println(string(b32))

	b64 := []byte("float64:")
	b64 = strconv.AppendFloat(b64, 3.1415926535, 'E', -1, 64)
	fmt.Println(string(b64))

	//int
	b10 := []byte("int (base 10):")
	b10 = strconv.AppendInt(b10, -42, 10)
	fmt.Println(string(b10))

	b16 := []byte("int (base 16):")
	b16 = strconv.AppendInt(b16, -42, 16)
	fmt.Println(string(b16))

	//quote
	q := []byte("quote:")
	q = strconv.AppendQuote(q, `"Fran & Freddie's Diner"`)
	fmt.Println(string(q))

	//rune
	r := []byte("rune:")
	r = strconv.AppendQuoteRune(r, '☺')
	fmt.Println(string(r))

}
