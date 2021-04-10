package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var wtr = bufio.NewWriter(os.Stdout)

func main() {

	if len(os.Args) > 1 && os.Args[1] == "i" {
		b, e := ioutil.ReadFile("./input")
		if e != nil {
			panic(e)
		}
		sc = bufio.NewScanner(strings.NewReader(strings.Replace(string(b), " ", "\n", -1)))
	}

	// defaut
	a := ni()
	b := ni()
	c := ni()
	s := ns()
	fmt.Printf("%d %s\n", a+b+c, s)

	// get with loop
	n := ni()
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = ni()
	}

	// output
	fmt.Fprintln(wtr, "Hello")
	fmt.Fprintln(wtr, "World")
	_ = wtr.Flush()
}

func init() {
	sc.Split(bufio.ScanWords)
}

func ni() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func nf() float64 {
	sc.Scan()
	f, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return f
}

func ns() string {
	sc.Scan()
	return sc.Text()
}
