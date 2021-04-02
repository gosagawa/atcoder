package main

import (
	"bufio"
	"fmt"
	"os"
)

var wtr = bufio.NewWriter(os.Stdout)

func main() {

	// defaut
	var a, b, c int
	var s string
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d %d", &b, &c)
	fmt.Scanf("%s", &s)
	fmt.Printf("%d %s\n", a+b+c, s)

	// get with loop
	var n int
	fmt.Scanf("%d", &n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &nums[i])
	}

	// output
	fmt.Fprintln(wtr, "Hello")
	fmt.Fprintln(wtr, "World")
	_ = wtr.Flush()
}
