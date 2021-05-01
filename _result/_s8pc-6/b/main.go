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
	ns := make([][2]int, n)
	for i := 0; i < n; i++ {
		ns[i][0] = ni()
		ns[i][1] = ni()
	}
	for i := 0; i < n; i++ {
		for i2 := 0; i2 < 2; i2++ {
			ent := ns[i][i2]
			for j := 0; j < n; j++ {
				for j2 := 0; j2 < 2; j2++ {
					ext := ns[j][j2]
					sum := 0
					for k := 0; k < n; k++ {
						sum += chmin(sec(ent, ns[k][0], ns[k][1], ext), sec(ent, ns[k][1], ns[k][0], ext))
						//out(ent, ext, i, i2, j, j2, k, sum)
					}

					o = chmin(o, sum)

				}
			}
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

func sec(a, b, c, d int) int {
	o := 0
	if a > b {
		o += a - b
	} else {
		o += b - a
	}
	if b > c {
		o += b - c
	} else {
		o += c - b
	}
	if c > d {
		o += c - d
	} else {
		o += d - c
	}
	return o
}
func chmin(a, b int) int {
	if a > b {
		return b
	}
	return a
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
