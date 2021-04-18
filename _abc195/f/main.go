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
var pn = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71}
var pnl = 20
var m = [][]bool{}
var mr = [][]int{}

func main() {

	defer flush()

	a := ni()
	b := ni()
	l := b - a + 1
	_ = a
	_ = b

	m = make([][]bool, l, l)
	mr = make([][]int, pnl, pnl)
	mods := make([]int, pnl, pnl)
	for i, v := range pn {
		mods[i] = a % v
	}

	mniso := make([]bool, l, l)
	for i := 0; i < l; i++ {
		m[i] = make([]bool, pnl, pnl)
	}
	for i, v := range pn {
		pos := 0
		if mods[i] != 0 {
			pos = -mods[i] + v
		}
		if pos > l {
			continue
		}
		c := 0
		for ip := pos; ip < l; ip += v {
			if c == 0 {
				c++
				continue
			}
			if c == 1 {
				mniso[ip-v] = true
				m[ip-v][i] = true
				mr[i] = append(mr[i], ip-v)
			}
			m[ip][i] = true
			mniso[ip] = true
			mr[i] = append(mr[i], ip)
			c++
		}
	}
	isoc := 0
	for i := 0; i < l; i++ {
		if !mniso[i] {
			isoc++
		}
	}
	mulisoc := int(math.Pow(2, float64(isoc)))

	/*
		out(mniso)
		out(isoc)
		out(mulisoc)
		out(mr)
	*/

	o := mulisoc + (l-isoc)*mulisoc

	for i := range pn {
		if len(mr[i]) == 0 {
			continue
		}
		check := make([]bool, pnl, pnl)
		o += count(i, check, false) * mulisoc
	}

	/*
		for i, v := range m {
			out(i+a, v[:5])
		}
	*/

	out(o)
}
func count(i int, check []bool, hasParent bool) int {
	c := 0

	for _, v := range mr[i] {
		able := true
		for ii, ib := range m[v] {
			if ib && check[ii] {
				able = false
				break
			}
			if ib && ii < i {
				able = false
				break
			}
		}
		if !able {
			continue
		}
		for ii, ib := range m[v] {
			if ib {
				check[ii] = ib
			}
		}
		if hasParent {
			c++
		}

		for ni := i + 1; ni < pnl; ni++ {
			if len(mr[ni]) == 0 {
				continue
			}
			c += count(ni, check, true)
		}
		for ii, ib := range m[v] {
			if ib {
				check[ii] = false
			}
		}
	}
	return c
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
