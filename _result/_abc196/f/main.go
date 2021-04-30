package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"math/cmplx"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"

	"github.com/mjibson/go-dsp/fft"
)

var sc = bufio.NewScanner(os.Stdin)
var wtr = bufio.NewWriter(os.Stdout)
var o int

func main() {

	if len(os.Args) > 1 && os.Args[1] == "i" {
		f, err := os.Create("cpu.pprof")
		if err != nil {
			panic(err)
		}

		if err := pprof.StartCPUProfile(f); err != nil {
			panic(err)
		}
		defer pprof.StopCPUProfile()

		b, e := ioutil.ReadFile("./input")
		if e != nil {
			panic(e)
		}
		sc = bufio.NewScanner(strings.NewReader(strings.Replace(string(b), " ", "\n", -1)))
	}
	sc.Buffer(make([]byte, 10000), 2000000)

	s := nb()
	t := nb()
	sil1, sil2 := convertSlice(s, false)
	til1, til2 := convertSlice(t, true)
	slen := len(s)
	tlen := len(t)
	o := tlen

	res1 := convolve(sil1, til2)
	res2 := convolve(sil2, til1)
	for i := tlen - 1; i < slen; i++ {
		res := res1[i] + res2[i]
		if res < o {
			o = res
		}
	}

	fmt.Println(o)
	_ = wtr.Flush()
}

func convolve(f, g []complex128) []int {
	sz := 1
	lenf := len(f)
	leng := len(g)
	alllen := lenf + leng
	res := make([]int, alllen)
	if lenf+leng < 60 {
		for i := 0; i < lenf; i++ {
			for j := 0; j < leng; j++ {
				res[i+j] += int(real(f[i])) * int(real(g[j]))
			}
		}
		return res
	}

	for sz < lenf+leng {
		sz *= 2
	}
	nf := make([]complex128, sz, sz)
	for i, v := range f {
		nf[i] = v
	}
	ng := make([]complex128, sz, sz)
	for i, v := range g {
		ng[i] = v
	}
	nf = fft.FFT(nf)
	ng = fft.FFT(ng)

	for i := 0; i < sz; i++ {
		nf[i] *= ng[i]
	}
	nf = fft.IFFT(nf)

	for i := 0; i < alllen; i++ {
		res[i] = int(real(nf[i]))
	}
	return res
}

func dft(f []complex128, inverse int) {
	sz := len(f)
	if sz == 1 {
		return
	}
	hsz := sz / 2

	veca := make([]complex128, hsz, hsz)
	vecb := make([]complex128, hsz, hsz)
	for i := 0; i < hsz; i++ {
		veca[i] = f[2*i]
		vecb[i] = f[2*i+1]
	}
	dft(veca, inverse)
	dft(vecb, inverse)

	var now complex128 = 1
	var zeta complex128 = cmplx.Rect(1.0, float64(inverse)*2*math.Pi/float64(sz))

	for i := 0; i < hsz; i++ {
		f[i] = veca[i%hsz] + now*vecb[i%hsz]
		now *= zeta
	}
}

func convertSlice(b []byte, r bool) ([]complex128, []complex128) {
	bl := len(b)
	res1 := make([]complex128, bl, bl)
	res2 := make([]complex128, bl, bl)

	for i, c := range b {
		index := i
		if r {
			index = bl - i - 1
		}
		if c == '0' {
			res1[index] = complex(1.0, 0)
		} else {
			res2[index] = complex(1.0, 0)
		}
	}
	return res1, res2
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
