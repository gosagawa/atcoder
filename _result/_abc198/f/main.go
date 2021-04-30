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

	n := ni() - 6

	// a + b + 4c
	dpsubA := make([][]int, 10, 10)
	dpsubB := make([][]int, 10, 10)
	dpsubC := make([][]int, 10, 10)
	dpsubD := make([][]int, 10, 10)
	dpsubE := make([][]int, 10, 10)
	for i := range dpsubA {
		dpsubA[i] = make([]int, 6, 6)
		dpsubB[i] = make([]int, 6, 6)
		dpsubC[i] = make([]int, 6, 6)
		dpsubD[i] = make([]int, 6, 6)
		dpsubE[i] = make([]int, 6, 6)
	}
	for a := 0; a <= 9; a++ {
		for b := 0; b <= 9; b++ {

			x := (3*a + 3*b) % 10
			j := (3*a + 3*b) / 10
			dpsubA[x][j]++

			for c := 0; c <= 9; c++ {
				x = (a + b + 4*c) % 10
				j = (a + b + 4*c) / 10
				dpsubB[x][j]++

				x = (2*a + 2*b + 2*c) % 10
				j = (2*a + 2*b + 2*c) / 10
				dpsubC[x][j]++

				for d := 0; d <= 9; d++ {
					x = (a + b + 2*c + 2*d) % 10
					j = (a + b + 2*c + 2*d) / 10
					dpsubD[x][j]++
					for e := 0; e <= 9; e++ {
						for f := 0; f <= 9; f++ {
							x = (a + b + c + d + e + f) % 10
							j = (a + b + c + d + e + f) / 10
							dpsubE[x][j]++
						}
					}
				}
			}
		}
	}

	nl := []int{}
	tmp := n
	for i := 0; i <= 20; i++ {
		nl = append(nl, tmp%10)
		tmp /= 10
		if tmp == 0 {
			break
		}
	}
	nll := len(nl)
	dpA := make([][]int, nll, nll)
	dpB := make([][]int, nll, nll)
	dpC := make([][]int, nll, nll)
	dpD := make([][]int, nll, nll)
	dpE := make([][]int, nll, nll)
	for i, v := range nl {
		dpA[i] = make([]int, 6, 6)
		dpB[i] = make([]int, 6, 6)
		dpC[i] = make([]int, 6, 6)
		dpD[i] = make([]int, 6, 6)
		dpE[i] = make([]int, 6, 6)
		for j := 0; j <= 5; j++ {
			if i == 0 {
				dpA[i][j] = dpsubA[v][j]
				dpB[i][j] = dpsubB[v][j]
				dpC[i][j] = dpsubC[v][j]
				dpD[i][j] = dpsubD[v][j]
				dpE[i][j] = dpsubE[v][j]
				continue
			}
			for k := 0; k <= 5; k++ {
				currentv := v - k
				adjust := 0
				if currentv < 0 {
					if j == 0 {
						break
					}
					currentv += 10
					adjust = 1
				}
				dpA[i][j] = madd(dpA[i][j], mmul(dpsubA[currentv][j-adjust], dpA[i-1][k]))
				dpB[i][j] = madd(dpB[i][j], mmul(dpsubB[currentv][j-adjust], dpB[i-1][k]))
				dpC[i][j] = madd(dpC[i][j], mmul(dpsubC[currentv][j-adjust], dpC[i-1][k]))
				dpD[i][j] = madd(dpD[i][j], mmul(dpsubD[currentv][j-adjust], dpD[i-1][k]))
				dpE[i][j] = madd(dpE[i][j], mmul(dpsubE[currentv][j-adjust], dpE[i-1][k]))
			}
		}
	}

	o := mdiv(madd(madd(madd(madd(dpA[nll-1][0]*8, dpB[nll-1][0]*6), dpC[nll-1][0]*6), dpD[nll-1][0]*3), dpE[nll-1][0]), 24)

	out(o)
	_ = wtr.Flush()
}

func init() {
	sc.Split(bufio.ScanWords)
}

const MOD = 998244353

func madd(a, b int) int {
	a += b
	if a > MOD {
		a -= MOD
	}
	return a
}

func mmul(a, b int) int {
	return a * b % MOD
}

func mdiv(a, b int) int {
	a %= MOD
	return a * modinv(b) % MOD
}

func modinv(a int) int {
	res := 1
	n := MOD - 2
	for n > 0 {
		if n&1 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		n >>= 1
	}
	return res
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

func out(v interface{}) {
	fmt.Fprintln(wtr, v)
}
