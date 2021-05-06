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
var xs = []int{}
var lenx = 0
var m = 0

func main() {

	defer flush()

	x := ns()
	m = ni()
	lenx = len(x)
	xs = make([]int, lenx)
	d := 0
	for i, v := range x {
		n, e := strconv.Atoi(string(v))
		if e != nil {
			panic(e)
		}
		xs[i] = n
		if n > d {
			d = n
		}
	}
	ok := d
	ng := m + 1
	switch lenx {
	case 1:
		if xs[0] <= m {
			out(1)
		} else {
			out(0)
		}
		return
	case 2:
		out(chmax(0, (m-xs[1])/(xs[0])-d))
		return
	case 3:
		ng = int(math.Sqrt(float64(m/xs[0]))) + 1
	case 4:
		ng = int(math.Cbrt(float64(m/xs[0]))) + 1
	case 5, 6, 7, 8:
		ng = int(math.Sqrt(math.Sqrt(float64(m)))) + 1
	default:
		ng = int(math.Sqrt(math.Sqrt(math.Sqrt(float64(m))))) + 1
	}

	o := chmax(0, culc(ok, ng)-d)
	out(o)
}
func chmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func chmin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func culc(ok, ng int) int {
	//out(ok, ng)
	if ng <= ok {
		return 0
	}
	if ng-ok == 1 {
		return ok
	}
	check := (ok + ng) / 2
	cs := formatInt(m, check)
	//	out(cs)
	lencs := len(cs)
	res := 0
	if lenx < lencs {
		res = 1
	} else if lenx > lencs {
		res = -1
	} else {
		for i := 0; i < lenx; i++ {
			if xs[i] < cs[i] {
				res = 1
				break
			}
			if xs[i] > cs[i] {
				res = -1
				break
			}
		}
	}
	if res >= 0 {
		ok = check
	} else {
		ng = check
	}

	return culc(ok, ng)

}

func formatInt(i, base int) []int {
	ret := make([]int, 0)

	for {
		ret = append([]int{i % base}, ret...)
		if i < base {
			break
		}
		i /= base
	}
	return ret
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
