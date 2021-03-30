package main

import (
	"bufio"
	"fmt"
	"os"
)

var wtr = bufio.NewWriter(os.Stdout)

func main() {
	var s string
	fmt.Scanf("%s", &s)

	fmt.Fprintln(wtr, string(s[1:])+string(s[0]))
	_ = wtr.Flush()
}
