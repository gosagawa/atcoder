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
var m map[int]map[int]bool

func main() {

	defer flush()

	o := 0
	n := ni()
	m = make(map[int]map[int]bool)
	ps := make([][2]int, n)
	for i := 0; i < n; i++ {
		x := ni()
		y := ni()
		if _, ok := m[x]; !ok {
			m[x] = make(map[int]bool)
		}
		m[x][y] = true
		ps[i] = [2]int{x, y}
	}
	for i := 0; i < n-3; i++ {
		for j := i + 1; j < n; j++ {
			dx := ps[i][0] - ps[j][0]
			dy := ps[i][1] - ps[j][1]

			pos1x := ps[i][0] + dy
			pos1y := ps[i][1] - dx
			if pos1x >= 0 && pos1y >= 0 && check(pos1x, pos1y) {
				pos1fx := pos1x - dx
				pos1fy := pos1y - dy
				if pos1fx >= 0 && pos1fy >= 0 && check(pos1fx, pos1fy) {
					subo := dx*dx + dy*dy
					if subo > o {
						o = subo
					}
				}
			}

			pos2x := ps[i][0] - dy
			pos2y := ps[i][1] + dx
			if pos2x >= 0 && pos2y >= 0 && check(pos2x, pos2y) {
				pos2fx := pos2x - dx
				pos2fy := pos2y - dy
				if pos2fx >= 0 && pos2fy >= 0 && check(pos2fx, pos2fy) {
					subo := dx*dx + dy*dy
					if subo > o {
						o = subo
					}
				}
			}
		}
		m[ps[i][0]][ps[i][1]] = false
	}

	out(o)
}

func check(x, y int) bool {
	if _, ok := m[x]; !ok {
		return false
	}
	if v, ok := m[x][y]; ok {
		return v
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
