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

	if len(os.Args) > 1 && os.Args[1] == "i" {
		b, e := ioutil.ReadFile("./input")
		if e != nil {
			panic(e)
		}
		sc = bufio.NewScanner(strings.NewReader(strings.Replace(string(b), " ", "\n", -1)))
	}

	// defaut
	a := ns()
	l := len(a)
	h := l / 2
	var o int
	if l%2 == 0 {
		o = int(math.Pow(10, float64(h-1))) - 1
		f, _ := strconv.Atoi(a[:h])
		s, _ := strconv.Atoi(a[h:])
		o += f - int(math.Pow(10, float64(h-1))) + 1
		if f > s {
			o--
		}
	} else {
		o = int(math.Pow(10, float64(h))) - 1
	}

	fmt.Printf("%d\n", o)
	_ = wtr.Flush()
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
