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

	a := ni()
	b := ni()
	c := ni()
	x := ni()
	y := ni()
	o := 0
	if x > y {
		diff := x - y
		if a+b < 2*c {
			o += a*y + b*y
		} else {
			o += 2 * c * y
		}
		if a < 2*c {
			o += diff * a
		} else {
			o += diff * 2 * c
		}
	} else {
		diff := y - x
		if a+b < 2*c {
			o += a*x + b*x
		} else {
			o += 2 * c * x
		}
		if b < 2*c {
			o += diff * b
		} else {
			o += diff * 2 * c
		}

	}

	out(o)
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
