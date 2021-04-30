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

	k := ni()
	kl := []int{k, k, k, k, k, k, k, k, k}

	_ = k
	s1 := ns()
	s1 = s1[:4]
	s1n, e := strconv.Atoi(s1)
	if e != nil {
		panic(e)
	}
	s1num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1num[s1n%10-1] *= 10
	s1num[(s1n/10)%10-1] *= 10
	s1num[(s1n/100)%10-1] *= 10
	s1num[s1n/1000-1] *= 10
	s1sum := 0
	for i := 0; i < 9; i++ {
		s1sum += s1num[i]
	}
	kl[s1n%10-1]--
	kl[(s1n/10)%10-1]--
	kl[(s1n/100)%10-1]--
	kl[s1n/1000-1]--

	s2 := ns()
	s2 = s2[:4]
	s2n, e := strconv.Atoi(s2)
	if e != nil {
		panic(e)
	}
	s2num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2num[s2n%10-1] *= 10
	s2num[(s2n/10)%10-1] *= 10
	s2num[(s2n/100)%10-1] *= 10
	s2num[s2n/1000-1] *= 10
	kl[s2n%10-1]--
	kl[(s2n/10)%10-1]--
	kl[(s2n/100)%10-1]--
	kl[s2n/1000-1]--
	s2sum := 0
	for i := 0; i < 9; i++ {
		s2sum += s2num[i]
	}

	pt := 0
	for i := 1; i <= 9; i++ {
		s1sum += s1num[i-1] * 9
		for j := 1; j <= 9; j++ {
			s2sum += s2num[j-1] * 9
			if s1sum > s2sum {
				if i == j {
					if kl[i-1] > 1 {
						pt += kl[i-1] * (kl[i-1] - 1)
					}
				} else {
					pt += kl[i-1] * kl[j-1]
				}
			}
			s2sum -= s2num[j-1] * 9
		}
		s1sum -= s1num[i-1] * 9
	}
	allpt := (9*k - 8) * (9*k - 9)
	out(float64(pt) / float64(allpt))
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
