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

	n := ni()
	a := ni()
	b := ni()
	o := 0
	d := 0
	for i := 1; i <= n; i++ {
		if i%10 == 0 {
			d -= 9
		}
		if i%100 == 0 {
			d -= 9
		}
		if i%1000 == 0 {
			d -= 9
		}
		if i%10000 == 0 {
			d -= 9
		}
		d++
		if a <= d && d <= b {
			o += i
		}

	}

	fmt.Fprintln(wtr, o)
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
