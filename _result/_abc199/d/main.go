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
var cmap map[int][]int

func main() {

	defer flush()

	n := ni()
	m := ni()
	o := 1

	cmap = make(map[int][]int)
	color := make([]int, n, n)
	checked := make([]bool, n, n)
	for i := 0; i < n; i++ {
		cmap[i] = []int{}
	}

	for i := 0; i < m; i++ {
		a := ni() - 1
		b := ni() - 1
		cmap[a] = append(cmap[a], b)
		cmap[b] = append(cmap[b], a)
	}

	for i := 0; i < n; i++ {
		if len(cmap[i]) == 0 {
			o *= 3
			continue
		}
		if checked[i] {
			continue
		}
		rmap := make(map[int]struct{})
		getRoute(i, rmap)
		rs := make([]int, len(rmap))
		rmapidx := 0
		for v := range rmap {
			rs[rmapidx] = v
			checked[v] = true
			rmapidx++
		}

		rslen := len(rs)
		o *= check(0, rslen, rs, color)
	}

	out(o)
}

func check(routei int, routelen int, route []int, color []int) int {
	pattern := 0
	i := route[routei]
	routei++
	for c := 1; c <= 3; c++ {
		color[i] = c
		colorError := false
		for _, v := range cmap[i] {
			if color[v] == color[i] {
				colorError = true
				break
			}
		}
		if colorError {
			continue
		}
		if routei < routelen {
			pattern += check(routei, routelen, route, color)
		} else {
			pattern++
		}
	}
	color[i] = 0
	return pattern
}

func getRoute(i int, route map[int]struct{}) {
	route[i] = struct{}{}
	for _, v := range cmap[i] {
		if _, ok := route[v]; ok {
			continue
		}
		getRoute(v, route)
	}
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
