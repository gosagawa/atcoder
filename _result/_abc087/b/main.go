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

func main() {

	if len(os.Args) > 1 && os.Args[1] == "i" {
		b, e := ioutil.ReadFile("./input")
		if e != nil {
			panic(e)
		}
		sc = bufio.NewScanner(strings.NewReader(strings.Replace(string(b), " ", "\n", -1)))
	}

	// defaut
	a := ni()
	b := ni()
	c := ni()
	x := ni()
	o := 0
	for i := 0; i <= a; i++ {
		subtotali := 0
		subtotali += i * 500
		if subtotali == x {
			o++
			continue
		} else if subtotali > x {
			break
		}
		for j := 0; j <= b; j++ {
			subtotalj := subtotali
			subtotalj += j * 100
			if subtotalj == x {
				o++
				continue
			} else if subtotalj > x {
				break
			}
			for k := 0; k <= c; k++ {
				subtotalk := subtotalj
				subtotalk += k * 50
				if subtotalk == x {
					o++
					continue
				} else if subtotalk > x {
					break
				}
			}
		}
	}

	// output
	fmt.Fprintln(wtr, o)
	_ = wtr.Flush()
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
