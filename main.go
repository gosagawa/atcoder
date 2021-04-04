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

	vMap := make(map[int]map[string][]int)
	searchedMap := make(map[int]map[int]struct{})
	for i := 1; i <= n; i++ {
		vMap[i] = make(map[string][]int)
		searchedMap[i] = make(map[int]struct{})
	}
	for i := 0; i < m; i++ {
		a := ni()
		b := ni()
		c := ns()

		vMap[a][c] = append(vMap[a][c], b)
		vMap[b][c] = append(vMap[b][c], a)
	}

	posL := [][2]int{
		{1, n},
	}

Loop:
	for i := 0; i < 10000000; i++ {
		newPosL := [][2]int{}
		for _, pos := range posL {
			for c, v := range vMap[pos[0]] {
				for _, sd := range v {
					if ev, ok := vMap[pos[1]][c]; ok {
						for _, ed := range ev {
							if sd == ed {
								o = (i + 1) * 2
								break Loop
							}
							if sd == pos[1] && ed == pos[0] {
								o = i*2 + 1
								break Loop
							}
							if _, ok := searchedMap[sd][ed]; !ok {
								newPos := [2]int{sd, ed}
								newPosL = append(newPosL, newPos)
							}
						}
					}
				}
			}
			searchedMap[pos[0]][pos[1]] = struct{}{}
		}
		if len(newPosL) == 0 {
			break
		}
		posL = newPosL
	}

	fmt.Fprintln(wtr, o)
	_ = wtr.Flush()
}
