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
var mm = make(map[int]map[int]struct{})

func main() {

	defer flush()

	ox := 0
	oy := 0
	n := ni()
	ns := make([][2]int, n-1)
	var bx, by int
	for i := 0; i < n; i++ {
		x := ni()
		y := ni()
		if i == 0 {
			bx, by = x, y
			continue
		}
		ns[i-1] = [2]int{bx - x, by - y}
	}
	mm = make(map[int]map[int]struct{})
	m := ni()
	ms := make([][2]int, m)
	for i := 0; i < m; i++ {
		x := ni()
		y := ni()
		if _, ok := mm[x]; !ok {
			mm[x] = make(map[int]struct{})
		}
		mm[x][y] = struct{}{}
		ms[i] = [2]int{x, y}
	}
	for _, pt := range ms {
		match := true
		for _, npt := range ns {
			if !check(pt[0]-npt[0], pt[1]-npt[1]) {
				match = false
				break
			}
		}
		if match {
			ox = pt[0] - bx
			out(ox, oy)
			break
		}
	}

	out(fmt.Sprintf("%v %v", ox, oy))
}

func check(x, y int) bool {
	if _, ok := mm[x]; !ok {
		return false
	}
	if _, ok := mm[x][y]; !ok {
		return false
	}
	return true
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
