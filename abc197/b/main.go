package main

import (
	"bufio"
	"fmt"
	"os"
)

var wtr = bufio.NewWriter(os.Stdout)

func main() {
	var h, w, x, y int
	var s string
	fmt.Scanf("%d %d %d %d", &h, &w, &x, &y)
	grid := make(map[int]map[int]bool)

	for i := 0; i < h; i++ {
		fmt.Scanf("%s", &s)
		grid[i] = make(map[int]bool)
		for iv, v := range s {
			grid[i][iv] = false
			if string(v) == "#" {
				grid[i][iv] = true
			}
		}
	}
	count := 1
	if x > 1 {
		for cx := x - 2; cx >= 0; cx-- {
			if grid[cx][y-1] {
				break
			}
			count++
		}
	}
	if x < h {
		for cx := x; cx < h; cx++ {
			if grid[cx][y-1] {
				break
			}
			count++
		}
	}
	if y > 1 {
		for cy := y - 2; cy >= 0; cy-- {
			if grid[x-1][cy] {
				break
			}
			count++
		}
	}
	if y < w {
		for cy := y; cy < w; cy++ {
			if grid[x-1][cy] {
				break
			}
			count++
		}
	}

	fmt.Fprintln(wtr, count)
	_ = wtr.Flush()
}
