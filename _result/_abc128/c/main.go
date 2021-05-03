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
	m := ni()
	o := 0

	ns := make([]int, n)
	b := 1
	for i := 0; i < m; i++ {
		k := ni()
		if i > 0 {
			b = b << 1
		}
		for j := 0; j < k; j++ {
			ki := ni()
			ns[ki-1] |= b
		}
	}
	ps := 0
	b = 1
	for i := 0; i < m; i++ {
		if i > 0 {
			b = b << 1
		}
		p := ni()
		ps |= p * b
	}

	check := int(math.Pow(2, float64(n)))
	for i := 0; i < check; i++ {
		tmp := i
		tmpcheck := 0
		for j := 0; j < n; j++ {
			if tmp&1 == 1 {
				tmpcheck ^= ns[j]
			}
			tmp = tmp >> 1
		}
		if tmpcheck == ps {
			o++
		}
	}

	out(o)
}

func sw(b bool) bool {
	return !b
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
