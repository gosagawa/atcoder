package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var wtr = bufio.NewWriter(os.Stdout)
var ans = make(map[int]int)

func main() {

	defer flush()

	n := ni()
	k := ni()
	if k == 0 {
		out(n)
		return
	}
	o := culc(n, k)

	out(o)
}

func culc(n, k int) int {
	if v, ok := ans[n]; ok {
		if k == 1 {
			return v
		}
		k--
		return culc(v, k)
	}
	sl := []int{}
	tmp := n
	for i := 0; i < 11; i++ {
		num := tmp % 10
		tmp /= 10
		sl = append(sl, num)
		if tmp == 0 {
			break
		}
	}
	sort.Slice(sl, func(i, j int) bool { return sl[i] > sl[j] })
	o := 0
	switch len(sl) {
	case 0:
		o = 0
	case 1:
		o = 0
	case 2:
		o = 9 * (sl[0] - sl[1])
	case 3:
		o = 99 * (sl[0] - sl[2])
	case 4:
		o = 999*(sl[0]-sl[3]) + 90*(sl[1]-sl[2])
	case 5:
		o = 9999*(sl[0]-sl[4]) + 990*(sl[1]-sl[3])
	case 6:
		o = 99999*(sl[0]-sl[5]) + 9990*(sl[1]-sl[4])
		o += 900 * (sl[2] - sl[3])
	case 7:
		o = 999999*(sl[0]-sl[6]) + 99990*(sl[1]-sl[5])
		o += 9900 * (sl[2] - sl[4])
	case 8:
		o = 9999999*(sl[0]-sl[7]) + 999990*(sl[1]-sl[6])
		o += 99900*(sl[2]-sl[5]) + 9000*(sl[3]-sl[4])
	case 9:
		o = 99999999*(sl[0]-sl[8]) + 9999990*(sl[1]-sl[7])
		o += 999900*(sl[2]-sl[6]) + 99000*(sl[3]-sl[5])
	case 10:
		o = 999999999*(sl[0]-sl[9]) + 99999990*(sl[1]-sl[8])
		o += 9999900*(sl[2]-sl[7]) + 999000*(sl[3]-sl[6])
		o += 90000 * (sl[4] - sl[5])
	case 11:
		o = 9999999999*(sl[0]-sl[10]) + 999999990*(sl[1]-sl[9])
		o += 99099900*(sl[2]-sl[8]) + 9990000*(sl[3]-sl[7])
		o += 990000 * (sl[4] - sl[6])
	}
	ans[n] = o
	//out(sl, n, o)
	if k == 1 {
		return o
	}
	k--
	return culc(o, k)
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
