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
var sevenmod map[int]int

func main() {

	defer flush()

	n := ni()
	s := ns()
	x := ns()
	r := 1

	if game(r, n, s, x, make(map[int]struct{})) {
		out("Takahashi")
	} else {
		out("Aoki")
	}
}
func game(r, n int, s string, x string, mod map[int]struct{}) bool {
	var rs int
	var rx string
	if len(s) > 0 {
		var e error
		rs, e = strconv.Atoi(string(s[len(s)-1]))
		if e != nil {
			panic(e)
		}
		rx = string(x[len(x)-1])
	}
	rsmod := sevenmod[rs]
	newmod := make(map[int]struct{})

	//out(r, n, rs, rx, rsmod, mod)

	nr := r + 1
	if len(s) > 0 {
		s = s[:len(s)-1]
	}
	if len(x) > 0 {
		x = x[:len(x)-1]
	}

	if r == 1 && r == n {
		switch rx {
		case "T":
			return true
		case "A":
			if rs%7 != 0 {
				return false
			}
			return true
		}
	} else if r == 1 {
		switch rx {
		case "T":
			newmod[0] = struct{}{}
			newmod[(14-rs)%7] = struct{}{}
		case "A":
			if rs%7 != 0 {
				return false
			}
			newmod[0] = struct{}{}
		}
		return game(nr, n, s, x, newmod)

	} else if r == n {
		a := rsmod
		b := 0
		switch rx {
		case "T":

			if _, oka := mod[a]; oka {
				return true
			}
			if _, okb := mod[b]; okb {
				return true
			}

		case "A":

			if _, oka := mod[a]; oka {
				if _, okb := mod[b]; okb {
					return true
				}
			}

		}
		return false
	} else {
		switch rx {
		case "T":
			for i := 0; i < 7; i++ {

				for v := range mod {
					if sevenmod[i+rs] == v {
						newmod[i] = struct{}{}
					}
					if sevenmod[i] == v {
						newmod[i] = struct{}{}
					}
				}

			}
		case "A":
			for i := 0; i < 7; i++ {
				t := false

				a := sevenmod[i+rs]
				b := sevenmod[i]
				if _, oka := mod[a]; oka {
					if _, okb := mod[b]; okb {
						t = true
					}
				}

				if t {
					newmod[i] = struct{}{}
				}
			}
			if len(newmod) == 0 {
				return false
			}
		}
		return game(nr, n, s, x, newmod)
	}
	return false
}

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
	if len(os.Args) > 1 && os.Args[1] == "i" {
		b, e := ioutil.ReadFile("./input")
		if e != nil {
			panic(e)
		}
		sc = bufio.NewScanner(strings.NewReader(strings.Replace(string(b), " ", "\n", -1)))
	}
	sevenmod = make(map[int]int)
	sevenmod[0] = 0
	sevenmod[1] = 3
	sevenmod[2] = 6
	sevenmod[3] = 2
	sevenmod[4] = 5
	sevenmod[5] = 1
	sevenmod[6] = 4
	sevenmod[7] = 0
	sevenmod[8] = 3
	sevenmod[9] = 6
	sevenmod[10] = 2
	sevenmod[11] = 5
	sevenmod[12] = 1
	sevenmod[13] = 4
	sevenmod[14] = 0
	sevenmod[15] = 3
	sevenmod[16] = 6
	sevenmod[17] = 2
	sevenmod[18] = 5
	sevenmod[19] = 1
	sevenmod[20] = 4
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

func out(v ...interface{}) {
	_, e := fmt.Fprintln(wtr, v...)
	if e != nil {
		panic(e)
	}
}

func flush() {
	e := wtr.Flush()
	if e != nil {
		panic(e)
	}
}
