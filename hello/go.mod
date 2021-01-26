// https://golang.org/doc/tutorial/call-module-code
module hello

go 1.14

replace example.com/greetings => ../greetings

require example.com/greetings v1.1.0
