package main

import (
	"expvar"
	"fmt"
)

func main() {
	// Package expvar provides a standardized interface to public variables, such as operation counters in servers. It exposes these variables via HTTP at /debug/vars in JSON format.

	fmt.Println(expvar.Handler())
}
