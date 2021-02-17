package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {
	// elastic()
	trailingtab()
}

func trailingtab() {
	// Observe that the third line has no trailing tab,
	// so its final cell is not part of an aligned column.
	const padding = 3
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, '-', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "a\tb\taligned\t")
	fmt.Fprintln(w, "aa\tbb\taligned\t")
	fmt.Fprintln(w, "aaa\tbbb\tunaligned") // no trailing tab
	fmt.Fprintln(w, "aaaa\tbbbb\taligned\t")
	w.Flush()
}

func elastic() {
	// Observe how the b's and the d's, despite appearing in the
	// second cell of each line, belong to different columns.
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, '.', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "a\tb\tc")
	fmt.Fprintln(w, "aa\tbb\tcc")
	fmt.Fprintln(w, "aaa\t") // trailing tab
	fmt.Fprintln(w, "aaaa\tdddd\teeee")
	w.Flush()
}
