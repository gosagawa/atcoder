package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"math/bits"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var wtr = bufio.NewWriter(os.Stdout)

func main() {

	f, err := os.Create("cpu.pprof")
	if err != nil {
		panic(err)
	}

	if err := pprof.StartCPUProfile(f); err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()

	if len(os.Args) > 1 && os.Args[1] == "i" {
		b, e := ioutil.ReadFile("./input")
		if e != nil {
			panic(e)
		}
		sc = bufio.NewScanner(strings.NewReader(strings.Replace(string(b), " ", "\n", -1)))
	}
	sc.Buffer(make([]byte, 10000), 2000000)

	s := nb()
	t := nb()
	slen := len(s)
	tlen := len(t)

	tuintlistlen := (tlen-1)/64 + 1
	tuintlist := make([]uint64, tuintlistlen, tuintlistlen)
	tuintindex := -1
	for i, c := range t {
		if i%64 == 0 {
			tuintindex++
		}

		d := c - '0'
		tuintlist[tuintindex] *= uint64(2)
		n1 := tuintlist[tuintindex] + uint64(d)
		tuintlist[tuintindex] = n1
	}
	alllen := slen - tlen + 1
	suintalllist := make([][]uint64, alllen, alllen)
	suintalllistindex := make([]int, alllen, alllen)
	suintalllistcounter := make([]int, alllen, alllen)
	for ai := 0; ai < alllen; ai++ {
		suintalllist[ai] = make([]uint64, tuintlistlen, tuintlistlen)
		suintalllistindex[ai] = -1
		suintalllistcounter[ai] = 0
	}
	bktree := &BKTree{}
	for i, c := range s {
		for ai := 0; ai < alllen; ai++ {
			if 0 <= i-ai && i-ai < tlen {
				if suintalllistcounter[ai]%64 == 0 {
					suintalllistindex[ai]++
				}
				suintindex := suintalllistindex[ai]
				d := c - '0'
				suintalllist[ai][suintindex] *= uint64(2)
				n1 := suintalllist[ai][suintindex] + uint64(d)
				suintalllist[ai][suintindex] = n1
				suintalllistcounter[ai]++
				if i-ai == tlen-1 {
					bktree.Add(suintalllist[ai])
				}
			}
		}
	}

	fmt.Println(bktree.Search(tuintlist))
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

func nb() []byte {
	sc.Scan()
	return sc.Bytes()
}

func distance(a, b uint64) int {

	d := 0
	var k uint64 = 1
	for i := 0; i < 64; i++ {
		if a&k != b&k {
			d++
		}
		k <<= 1
	}
	return d
}

type BKTree struct {
	root *node
}

type node struct {
	entry    hashValue
	children []struct {
		distance int
		node     *node
	}
}

func (n *node) addChild(e hashValue) {
	newnode := &node{entry: e}
loop:
	d := n.entry.Distance(e)
	for _, c := range n.children {
		if c.distance == d {
			n = c.node
			goto loop
		}
	}
	n.children = append(n.children, struct {
		distance int
		node     *node
	}{d, newnode})
}

func (bk *BKTree) Add(entry hashValue) {
	if bk.root == nil {
		bk.root = &node{
			entry: entry,
		}
		return
	}
	bk.root.addChild(entry)
}

type hashValue []uint64

func (h hashValue) Distance(e hashValue) int {
	a := []uint64(h)
	b := []uint64(e)

	d := 0
	for i, v := range a {
		d += bits.OnesCount64(v ^ b[i])
	}
	return d
}

func (bk *BKTree) Search(needle hashValue) int {

	result := 1000000
	for i := 0; i <= 100000; i++ {
		tolerance := int(math.Pow(2, float64(i)))
		candidates := []*node{bk.root}
		for len(candidates) != 0 {
			c := candidates[len(candidates)-1]
			candidates = candidates[:len(candidates)-1]
			d := c.entry.Distance(needle)
			if d <= tolerance && d < result {
				result = d
			}

			low, high := d-tolerance, d+tolerance
			for _, c := range c.children {
				if low <= c.distance && c.distance <= high {
					candidates = append(candidates, c.node)
				}
			}
		}
		if result != 1000000 {
			break
		}
	}
	return result
}
