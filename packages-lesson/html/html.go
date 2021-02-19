package main

import (
	"fmt"
	"html"
)

func main() {
	// escape()
	unescape()
}

func unescape() {
	const s = `&quot;Fran &amp; Freddie&#39;s Diner&quot; &lt;tasty@example.com&gt;`
	fmt.Println(html.UnescapeString(s))
}

func escape() {
	const s = `"Fran & Freddie's Diner" <tasty@example.com>`
	fmt.Println(html.EscapeString(s))
}
