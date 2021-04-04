package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var wtr = bufio.NewWriter(os.Stdout)

func main() {

	if len(os.Args) > 1 && os.Args[1] == "i" {
		b, e := ioutil.ReadFile("./input")
		if e != nil {
			panic(e)
		}
		sc = bufio.NewScanner(strings.NewReader(strings.Replace(string(b), " ", "\n", -1)))
	}

	h := ni()
	w := ni()
	a := ni()
	b := ni()

	m := make([][]bool, w)
	for i := 0; i < w; i++ {
		m[i] = make([]bool, h)
	}
	o := getCount(m, h, w, a, b, 0, 0)

	fmt.Fprintln(wtr, o)
	_ = wtr.Flush()
}

func getCount(m [][]bool, h, w, a, b, x, y int) int {
	//	fmt.Println(m, a, b, x, y)
	var o int
	if m[x][y] {
		nextX := x + 1
		nextY := y
		if nextX >= w {
			nextX = 0
			nextY = y + 1
		}
		if nextY == h {
			return 1
		}
		return getCount(m, h, w, a, b, nextX, nextY)
	}

	// horizen
	if x+1 < w && !m[x+1][y] && a > 0 {
		newM := make([][]bool, w)
		copy(newM, m)
		for i := 0; i < w; i++ {
			newM[i] = make([]bool, h)
			copy(newM[i], m[i])
		}
		newM[x][y] = true
		newM[x+1][y] = true
		nextX := x + 2
		nextY := y
		if nextX >= w {
			nextX = 0
			nextY = y + 1
		}
		if nextY == h {
			return 1
		}
		o += getCount(newM, h, w, a-1, b, nextX, nextY)
	}

	// vertical
	if y+1 < h && a > 0 {
		newM := make([][]bool, w)
		copy(newM, m)
		for i := 0; i < w; i++ {
			newM[i] = make([]bool, h)
			copy(newM[i], m[i])
		}
		newM[x][y] = true
		newM[x][y+1] = true
		nextX := x + 1
		nextY := y
		if nextX >= w {
			nextX = 0
			nextY = y + 1
		}
		if nextY == h {
			return 1
		}
		o += getCount(newM, h, w, a-1, b, nextX, nextY)
	}

	// small
	if b > 0 {
		newM := make([][]bool, w)
		copy(newM, m)
		for i := 0; i < w; i++ {
			newM[i] = make([]bool, h)
			copy(newM[i], m[i])
		}
		newM[x][y] = true
		nextX := x + 1
		nextY := y
		if nextX >= w {
			nextX = 0
			nextY = y + 1
		}
		if nextY == h {
			return 1
		}
		o += getCount(newM, h, w, a, b-1, nextX, nextY)
	}
	return o
}

func init() {
	sc.Split(bufio.ScanWords)
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
