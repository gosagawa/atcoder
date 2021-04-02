package main

import (
	"bufio"
	"fmt"
	"os"
)

var wtr = bufio.NewWriter(os.Stdout)

func main() {

	var n, o, x, c, cMax int
	fmt.Scanf("%d", &n)
	cMap := make(map[int][]int)
	cMapMax := make(map[int]int)
	cMapMin := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d", &x, &c)
		if _, ok := cMap[c]; !ok {
			cMap[c] = []int{}
			cMapMax[c] = x
			cMapMin[c] = x
		}
		cMap[c] = append(cMap[c], x)
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

	rMap := make(map[int]int)
	rMap[0] = 0

	for i := 1; i <= cMax-1; i++ {
		if len(cMap[i]) == 0 {
			continue
		}
		xList := cMap[i]
		nextRMap := make(map[int]int)
		for start, stackCost := range rMap {
			var minCost int
			if minCost != 0 && stackCost > minCost {
				continue
			}
			for _, next := range xList {
				cost := stackCost + getCost(start, next, cMapMin[i], cMapMax[i])
				//				println(start, next, cMapMin[i], cMapMax[i])
				//				println(cost)
				if _, ok := nextRMap[next]; !ok {
					nextRMap[next] = cost
				} else if nextRMap[next] > cost {
					nextRMap[next] = cost
				}
				if minCost == 0 || minCost > cost {
					minCost = cost
				}
			}
		}

		rMap = make(map[int]int)
		for start, stackCost := range nextRMap {
			rMap[start] = stackCost
		}
	}

	if cMapMin[cMax] > 0 {
		cMapMin[cMax] = 0
	}
	if cMapMax[cMax] < 0 {
		cMapMax[cMax] = 0
	}
	for start, stackCost := range rMap {
		cost := stackCost + getCost(start, 0, cMapMin[cMax], cMapMax[cMax])
		//		println(start, 0, cMapMax[cMax], cMapMin[cMax])
		//		println(cost)
		if o == 0 {
			o = cost
		} else if o > cost {
			o = cost
		}
	}

	fmt.Fprintln(wtr, o)
	_ = wtr.Flush()
}

func getCost(a, b, min, max int) int {
	if a > b {
		a, b = b, a
	}

	var result int
	if min < a {
		result += a - min
		a = min
	}
	if max > b {
		result += max - b
		b = max
	}
	result += b - a
	return result
}
