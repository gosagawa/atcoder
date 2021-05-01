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

	s := ns()
	rev := false
	lens := len(s)
	o := make([]rune, 2*lens)
	posfrom := lens - 1
	posto := lens
	for _, st := range s {
		if string(st) == "R" {
			if rev {
				rev = false
			} else {
				rev = true
			}
		} else {
			if rev {
				if st == o[posfrom+1] {
					posfrom++
				} else {
					o[posfrom] = st
					posfrom--
				}
			} else {
				if st == o[posto-1] {
					posto--
				} else {
					o[posto] = st
					posto++
				}
			}
		}
	}
	o = o[posfrom+1 : posto]
	l := len(o)
	if rev {
		for i := 0; i < l/2; i++ {
			o[0+i], o[l-1-i] = o[l-1-i], o[0+i]
		}
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
