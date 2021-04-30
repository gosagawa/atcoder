package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var wtr = bufio.NewWriter(os.Stdout)

func main() {

	defer flush()

	o := math.MaxInt64

	n := ni()
	var mina, minb int
	for i := 0; i < n; i++ {
		a := ni()
		b := ni()
		if i == 0 {
			mina = a
			minb = b
			o = a + b
			continue
		}
		if high(mina, b) < o {
			o = high(mina, b)
		}
		if high(minb, a) < o {
			o = high(minb, a)
		}
		if a+b < o {
			o = a + b
		}
		if mina > a {
			mina = a
		}
		if minb > b {
			minb = b
		}

	}

	out(o)
}

func high(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
	if len(os.Args) > 1 && os.Args[1] == "i" {
		b, e := ioutil.ReadFile("./input")
		if e != nil {
			panic(e)
		}
		sc = bufio.NewScanner(strings.NewReader(strings.Replace(string(b), " ", "\n", -1)))
	}
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

func out(v ...interface{}) {
	_, e := fmt.Fprintln(wtr, v...)
	if e != nil {
		panic(e)
	}
}

func flush() {
	e := wtr.Flush()
	if e != nil {
		panic(e)
	}
}
