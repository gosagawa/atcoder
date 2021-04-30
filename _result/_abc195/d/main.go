package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var wtr = bufio.NewWriter(os.Stdout)

type buggage struct {
	size  int
	value int
}
type box struct {
	size int
	pos  int
}

func main() {

	n := ni()
	m := ni()
	q := ni()
	ws := make([]buggage, n)
	for i := 0; i < n; i++ {
		w := ni()
		v := ni()
		ws[i] = buggage{w, v}
	}
	sort.Slice(ws, func(i, j int) bool { return ws[i].size > ws[j].size })
	sort.Slice(ws, func(i, j int) bool { return ws[i].value > ws[j].value })

	xs := make([]box, m)
	xssaved := make([]box, m)
	for i := 0; i < m; i++ {
		xs[i] = box{ni(), i}
	}
	sort.Slice(xs, func(i, j int) bool { return xs[i].size < xs[j].size })
	copy(xssaved, xs)
	posMap := make([]int, m)
	for i, v := range xs {
		posMap[v.pos] = i
	}

	for i := 0; i < q; i++ {
		o := 0
		l := ni() - 1
		r := ni() - 1
		for b := l; b <= r; b++ {
			xs[posMap[b]].size = 0
		}
		used := []int{}
		for _, w := range ws {
			for i, x := range xs {
				if w.size <= x.size {
					o += w.value
					used = append(used, i)
					xs[i].size = 0
					break
				}
			}
		}

		out(o)
		for b := l; b <= r; b++ {
			xs[posMap[b]].size = xssaved[posMap[b]].size
		}
		for _, v := range used {
			xs[v].size = xssaved[v].size
		}
	}

	flush()
}

func init() {
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

func out(v interface{}) {
	_, e := fmt.Fprintln(wtr, v)
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
