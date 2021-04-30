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

	n := ni()
	s := ns()

	q := ni()
	lens := len(s)
	o := make([]byte, lens, lens)
	pos := make([]int, lens, lens)
	for i := 0; i < lens; i++ {
		pos[i] = i
	}

	turn := false
	for i := 0; i < q; i++ {

		t := ni()
		a := ni() - 1
		b := ni() - 1
		switch t {
		case 1:
			if turn {
				if a >= n {
					a = a - n
				} else {
					a = a + n
				}
				if b >= n {
					b = b - n
				} else {
					b = b + n
				}
			}
			pos[a], pos[b] = pos[b], pos[a]
		case 2:
			if turn == false {
				turn = true
			} else {
				turn = false
			}
		}
	}
	if turn {
		pos = append(pos[n:2*n], pos[:n]...)
	}
	for i := 0; i < lens; i++ {
		o[i] = s[pos[i]]
	}
	out(string(o))

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
