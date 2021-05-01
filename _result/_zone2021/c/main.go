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
var nsl [][5]int

func main() {

	defer flush()

	n := ni()
	min := math.MaxInt64
	max := 0
	nsl = make([][5]int, n)
	for i := 0; i < n; i++ {
		a := ni()
		b := ni()
		c := ni()
		d := ni()
		e := ni()
		nsl[i] = [5]int{a, b, c, d, e}
		for _, v := range nsl[i] {
			if v < min {
				min = v
			}
			if v > max {
				max = v
			}
		}
	}
	ok := min
	ng := max + 1
	for {
		if ng-ok == 1 {
			break
		}
		checknum := (ng + ok) / 2
		if check(checknum) {
			ok = checknum
		} else {
			ng = checknum
		}
	}

	out(ok)
}

func check(x int) bool {
	m := make(map[int]struct{})
	for _, v := range nsl {
		c := 0
		for i := 0; i < 5; i++ {
			c = c << 1
			if v[i] >= x {
				c = (c | 1)
			}
		}
		m[c] = struct{}{}
	}
	for a := range m {
		for b := range m {
			for c := range m {
				if a|b|c == 31 {
					return true
				}
			}
		}
	}
	return false
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
