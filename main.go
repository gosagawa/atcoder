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

	slen, s := stouis(ns())
	tlen, t := stouis(ns())

	fmt.Println(maxdistance(s, t, slen, tlen))
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

func nb() []byte {
	sc.Scan()
	return sc.Bytes()
}

func stouis(s string) (int, []uint64) {
	ls := len(s)
	c := (ls-1)/64 + 1
	uis := make([]uint64, c, c)
	for i := 0; i < c; i++ {
		f := ls - (i+1)*64
		t := ls - i*64
		if f < 0 {
			f = 0
		}
		if t < 0 {
			t = ls
		}
		subs := s[f:t]
		u, e := strconv.ParseUint(subs, 2, 64)
		if e != nil {
			panic(e)
		}
		uis[i] = u
	}
	return ls, uis
}

func maxdistance(sl, tl []uint64, ss, ts int) int {

	d := ts
	var k uint64 = 1
	var sActive uint64 = sl[0]
	allc := ss - ts + 1
	var checkerC, slcounter int

	checker := make([]uint64, ts, ts)
	checkerd := make([]int, ts, ts)
	checkeri := make([]int, ts, ts)

	activeChecker := 1
	for si := 0; si < ss; si++ {
		if si < allc {
			checker[checkerC] = tl[0]
		}
		for c := 0; c < activeChecker; c++ {
			if checker[c]&k != sActive&k {
				checkerd[c]++
			}
			if checkeri[c]%64 == 63 {
				nextI := checker[c] % 64
				checker[c] = tl[nextI]
			} else {
				checker[c] >>= 1
			}
			checkeri[c]++
		}
		/*
		           fmt.Println("--------------")
		   		fmt.Println(checker)
		   		fmt.Println(checkerd)
		   		fmt.Println(checkeri)
		*/
		if si >= ts-1 {
			finishI := (si - ts + 1) % ts
			/*
				fmt.Println("-------")
				fmt.Println(finishI, checkerd[finishI])
			*/
			if checkerd[finishI] < d {
				d = checkerd[finishI]
			}
		}

		if si < ts-1 && si < allc-1 {
			activeChecker++
		}

		if si%64 == 63 {
			slcounter++
			sActive = sl[slcounter]
		} else {
			sActive >>= 1
		}

		if si%ts == ts-1 {
			checkerC = 0
			checkerd[0] = 0
			checkeri[0] = 0
		} else {
			checkerC++
		}

	}
	return d
}
func distance(a, b uint64) int {

	d := 0
	var k uint64 = 1
	for i := 0; i < 64; i++ {
		if a&k != b&k {
			d++
		}
		k <<= 1
	}
	return d
}
