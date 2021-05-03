package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var wtr = bufio.NewWriter(os.Stdout)

func main() {

	defer flush()

	n := ni()
	m := ni()
	o := 1

	ns := make([]int, n)
	for i := 0; i < n; i++ {
		ns[i] |= 1 << i
	}
	for i := 0; i < m; i++ {
		x := ni() - 1
		y := ni() - 1
		ns[x] |= 1 << y
		ns[y] |= 1 << x
	}

	c := 2 << n
	for i := 3; i < c; i++ {
		tmp := i
		tmpch := i
		for j := 0; j < n; j++ {
			if (tmp>>j)&1 == 1 {
				tmpch &= ns[j]
			}
		}
		bc := bits.OnesCount64(uint64(tmpch))
		o = chmax(o, bc)
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

func chmax(a, b int) int {
	if a > b {
		return a
	}
	return b
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
