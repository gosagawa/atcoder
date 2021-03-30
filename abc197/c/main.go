package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var wtr = bufio.NewWriter(os.Stdout)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &nums[i])
	}
	if n == 1 {
		fmt.Fprintln(wtr, nums[0])
		_ = wtr.Flush()
		return
	}
	result := -1
	patern := int(math.Pow(2, float64(n)-1))
	for i := 1; i < patern; i++ {

		numGroupResult := make([]int, n)

		groupIndex := 0
		for index, v := range nums {
			if index == 0 {
				numGroupResult[groupIndex] = v
				continue
			}
			if bitOn(i, index) {
				groupIndex++
				numGroupResult[groupIndex] = v
			} else {
				numGroupResult[groupIndex] = numGroupResult[groupIndex] | v
			}
		}
		tmp := 0
		for _, v := range numGroupResult {
			tmp = tmp ^ v
		}

		if result == -1 || tmp < result {
			result = tmp
		}

	}

	fmt.Fprintln(wtr, result)
	_ = wtr.Flush()
}

func bitOn(n, pos int) bool {

	pos = pos - 1

	return (n>>pos)&1 == 1
}
