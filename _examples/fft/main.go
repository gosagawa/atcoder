package main

import (
	"fmt"
	"math"
	"math/cmplx"

	"github.com/mjibson/go-dsp/fft"
)

func main() {

	fmt.Println(convolve([]int{1, 2, 3, 4}, []int{1, 2, 4, 8}))
	fmt.Println(convolve([]int{1, 0, 1, 0, 1, 0, 1}, []int{0, 1, 0}))
}

func convolve(f, g []int) []float64 {
	sz := 1
	lenf := len(f)
	leng := len(g)
	for sz < lenf+leng {
		sz *= 2
	}
	nf := make([]complex128, sz, sz)
	for i, v := range f {
		nf[i] = complex(float64(v), 0)
	}
	ng := make([]complex128, sz, sz)
	for i, v := range g {
		ng[i] = complex(float64(v), 0)
	}
	nf = fft.FFT(nf)
	ng = fft.FFT(ng)

	for i := 0; i < sz; i++ {
		nf[i] *= ng[i]
	}
	nf = fft.IFFT(nf)

	res := make([]float64, sz, sz)
	for i := 0; i < sz; i++ {
		res[i] = real(nf[i])
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
