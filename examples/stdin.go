package main

import (
	"bufio"
	"fmt"
	"os"
)

var wtr = bufio.NewWriter(os.Stdout)

func main() {
	var a, b, c int
	var s string
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d %d", &b, &c)
	fmt.Scanf("%s", &s)
	fmt.Printf("%d %s\n", a+b+c, s)

	fmt.Fprintln(wtr, "Hello")
	fmt.Fprintln(wtr, "World")
	_ = wtr.Flush()
}
