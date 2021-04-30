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

	if len(os.Args) > 1 && os.Args[1] == "i" {
		b, e := ioutil.ReadFile("./input")
		if e != nil {
			panic(e)
		}
		sc = bufio.NewScanner(strings.NewReader(strings.Replace(string(b), " ", "\n", -1)))
	}

	s1 := ns()
	s2 := ns()
	s3 := ns()
	s1len := len(s1)
	s2len := len(s2)
	s3len := len(s3)

	charMap := make(map[string]struct{})
	noZeroMap := make(map[string]struct{})
	var ns1, ns2, ns3 []string
	for i := 0; i < s3len-s1len; i++ {
		ns1 = append(ns1, "")
	}
	for i := 0; i < s3len-s2len; i++ {
		ns2 = append(ns2, "")
	}

	for i, v := range s1 {
		s := string(v)
		if i == 0 {
			noZeroMap[s] = struct{}{}
		}
		charMap[s] = struct{}{}
		ns1 = append(ns1, s)
	}
	for i, v := range s2 {
		s := string(v)
		if i == 0 {
			noZeroMap[s] = struct{}{}
		}
		charMap[s] = struct{}{}
		ns2 = append(ns2, s)
	}
	for i, v := range s3 {
		s := string(v)
		if i == 0 {
			noZeroMap[s] = struct{}{}
		}
		charMap[s] = struct{}{}
		ns3 = append(ns3, s)
	}
	if len(charMap) > 10 {
		fmt.Println("UNSOLVABLE")
		return
	}
	powMap := make([]int, 11, 11)
	for i := 0; i < 11; i++ {
		powMap[i] = int(math.Pow10(i))
	}

	tmpMap := make(map[string]int)
	usedNum := make(map[int]struct{})
	numMap, solvable := check(tmpMap, usedNum, noZeroMap, ns1, ns2, ns3, 0)
	if solvable {
		fmt.Fprintln(wtr, convertNum(ns1, numMap, powMap))
		fmt.Fprintln(wtr, convertNum(ns2, numMap, powMap))
		fmt.Fprintln(wtr, convertNum(ns3, numMap, powMap))
	} else {

		fmt.Fprintln(wtr, "UNSOLVABLE")
	}

	_ = wtr.Flush()
}

func check(numMap map[string]int, usedMap map[int]struct{}, noZeroMap map[string]struct{}, s1, s2, s3 []string, add int) (map[string]int, bool) {

	idx := len(s1) - 1

	ss1 := s1[idx]
	ss2 := s2[idx]
	ss3 := s3[idx]
	num1s := 0
	num1e := 9
	num1d := false
	if ss1 == "" {
		num1e = 0
		num1d = true
	} else if v, ok := numMap[ss1]; ok {
		num1s = v
		num1e = v
		num1d = true
	}
	for i := num1s; i <= num1e; i++ {
		if _, ok := noZeroMap[ss1]; i == 0 && ok {
			continue
		}
		if _, ok := usedMap[i]; !num1d && ok {
			continue
		}

		newUsedMap := make(map[int]struct{})
		for n := range usedMap {
			newUsedMap[n] = struct{}{}
		}
		if !num1d {
			newUsedMap[i] = struct{}{}
		}
		newNumMap := make(map[string]int)
		for s, n := range numMap {
			newNumMap[s] = n
		}
		newNumMap[ss1] = i

		num2s := 0
		num2e := 9
		num2d := false

		if ss2 == "" {
			num2e = 0
			num2d = true
		} else if v, ok := newNumMap[ss2]; ok {
			num2s = v
			num2e = v
			num2d = true
		}
		for j := num2s; j <= num2e; j++ {
			if _, ok := noZeroMap[ss2]; j == 0 && ok {
				continue
			}
			if _, ok := newUsedMap[j]; !num2d && ok {
				continue
			}
			newnewUsedMap := make(map[int]struct{})
			for n := range newUsedMap {
				newnewUsedMap[n] = struct{}{}
			}
			newnewUsedMap[j] = struct{}{}
			newnewNumMap := make(map[string]int)
			for s, n := range newNumMap {
				newnewNumMap[s] = n
			}
			if !num2d {
				newnewNumMap[ss2] = j
			}

			num3s := 0
			num3e := 9
			num3d := false
			if ss3 == "" {
				num3e = 0
				num3d = true
			} else if v, ok := newnewNumMap[ss3]; ok {
				num3s = v
				num3e = v
				num3d = true
			}
			for k := num3s; k <= num3e; k++ {
				if _, ok := noZeroMap[ss3]; k == 0 && ok {
					continue
				}
				if _, ok := newnewUsedMap[k]; !num3d && ok {
					continue
				}
				if k == (i+j+add)%10 {
					newnewnewUsedMap := make(map[int]struct{})
					for n := range newnewUsedMap {
						newnewnewUsedMap[n] = struct{}{}
					}
					newnewnewUsedMap[k] = struct{}{}
					newnewnewNumMap := make(map[string]int)
					for s, n := range newnewNumMap {
						newnewnewNumMap[s] = n
					}
					if !num3d {
						newnewnewNumMap[ss3] = k
					}

					newAdd := 0
					if i+j+add >= 10 {
						newAdd = 1
					}
					if idx == 0 {
						if newAdd == 0 {
							return newnewnewNumMap, true
						}
					} else {

						numMap, solvable := check(newnewnewNumMap, newnewnewUsedMap, noZeroMap, s1[:idx], s2[:idx], s3[:idx], newAdd)
						if solvable {
							return numMap, true
						}
					}
				}
			}
		}
	}
	return numMap, false
}

func convertNum(s []string, numMap map[string]int, powMap []int) int {
	l := len(s)
	o := 0
	for i := l - 1; i >= 0; i-- {
		o += numMap[s[i]] * powMap[l-1-i]
	}
	return o
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
