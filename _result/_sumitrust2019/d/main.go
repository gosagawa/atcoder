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
	o := 0
	_ = n
	l1 := make([]bool, 10, 10)
	l2 := make([][]bool, 10, 10)
	l3 := make([][][]bool, 10, 10)
	for i := 0; i < 10; i++ {
		l2[i] = make([]bool, 10, 10)
		l3[i] = make([][]bool, 10, 10)
		for j := 0; j < 10; j++ {
			l3[i][j] = make([]bool, 10, 10)
		}
	}

	for _, v := range s {
		i, e := strconv.Atoi(string(v))
		if e != nil {
			panic(e)
		}
		for l1i, l1v := range l1 {
			if l1v {
				for l2i, l2v := range l2[l1i] {
					if l2v {
						l3[l1i][l2i][i] = true
					}
				}
				l2[l1i][i] = true
			}
		}
		l1[i] = true
	}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			for m := 0; m < 10; m++ {
				if l3[i][j][m] {
					o++
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
