package main

import (
	"bufio"
	"fmt"
	"strconv"
)

var sc = bufio.newscanner(os.stdin)
var wtr = bufio.newwriter(os.stdout)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {

	var n, x, c, cMax int

	sc.Split(bufio.ScanWords)
	n = nextInt()

	cMapMax := make(map[int]int)
	cMapMin := make(map[int]int)
	for i := 0; i < n; i++ {
		x = nextInt()
		c = nextInt()

		if _, ok := cMapMax[c]; !ok {
			cMapMax[c] = x
			cMapMin[c] = x
		}
		if c > cMax {
			cMax = c
		}
		if cMapMax[c] < x {
			cMapMax[c] = x
		}
		if cMapMin[c] > x {
			cMapMin[c] = x
		}
	}

	var resultA, resultB, positionA, positionB int

	for i := 1; i <= cMax; i++ {
		if _, ok := cMapMax[i]; !ok {
			continue
		}
		min := cMapMin[i]
		max := cMapMax[i]

		resultAA := resultA
		resultAB := resultA
		resultBA := resultB
		resultBB := resultB

		if max > positionA {
			resultAA += max - positionA
		} else {
			resultAA += positionA - max
		}
		if min > positionA {
			resultAB += min - positionA
		} else {
			resultAB += positionA - min
		}
		if max > positionB {
			resultBA += max - positionB
		} else {
			resultBA += positionB - max
		}
		if min > positionB {
			resultBB += min - positionB
		} else {
			resultBB += positionB - min
		}

		resultA = resultAA
		if resultAA > resultBA {
			resultA = resultBA
		}
		resultB = resultAB
		if resultAB > resultBB {
			resultB = resultBB
		}

		resultA += max - min
		resultB += max - min

		positionA = min
		positionB = max

		/*
			println("------------------------------------")
			println(i)
			println(resultAA, resultBA, resultAB, resultBB)
			println(resultA, resultB)
			println(positionA, positionB)
		*/
	}

	if cMapMax[cMax] > 0 {
		resultB += cMapMax[cMax]
	} else {
		resultB -= cMapMax[cMax]
	}
	if cMapMin[cMax] > 0 {
		resultA += cMapMin[cMax]
	} else {
		resultA -= cMapMin[cMax]
	}

	if resultA < resultB {
		fmt.Fprintln(wtr, resultA)
	} else {
		fmt.Fprintln(wtr, resultB)
	}

	_ = wtr.Flush()
}
