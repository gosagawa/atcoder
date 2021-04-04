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

func ni() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func ns() string {
	sc.Scan()
	return sc.Text()
}

func main() {

	sc.Split(bufio.ScanWords)

	if len(os.Args) > 1 && os.Args[1] == "i" {
		b, e := ioutil.ReadFile("./input")
		if e != nil {
			panic(e)
		}
		sc = bufio.NewScanner(strings.NewReader(strings.Replace(string(b), " ", "\n", -1)))
	}

	o := -1
	n := ni()
	m := ni()

	vMap := make([][][]int, n, n)
	searchedMap := make([][]bool, n)
	for i := 0; i < n; i++ {
		vMap[i] = make([][]int, 26, 26)
		searchedMap[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		a := ni() - 1
		b := ni() - 1
		c := int([]rune(ns())[0]) - 97

		vMap[a][c] = append(vMap[a][c], b)
		vMap[b][c] = append(vMap[b][c], a)
	}

	posL := [][2]int{
		{0, n - 1},
	}
	i := 0

Loop:
	for {
		newPosL := [][2]int{}
		for _, pos := range posL {
			for c, v := range vMap[pos[0]] {
				for _, sd := range v {
					for _, ed := range vMap[pos[1]][c] {
						if !searchedMap[sd][ed] {
							if sd == pos[1] && ed == pos[0] {
								o = i*2 + 1
								break Loop
							}
							if sd == ed {
								tmp := (i + 1) * 2
								if o == -1 {
									o = tmp
								}
							}
							newPos := [2]int{sd, ed}
							newPosL = append(newPosL, newPos)
							searchedMap[sd][ed] = true
						}
					}
				}
			}
			searchedMap[pos[0]][pos[1]] = true
		}
		if o != -1 {
			break
		}
		if len(newPosL) == 0 {
			break
		}
		posL = newPosL
		i++
	}

	fmt.Fprintln(wtr, o)
	_ = wtr.Flush()
}
