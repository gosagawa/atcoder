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
var tuintlist []uint64
var o int

func main() {

	if len(os.Args) > 1 && os.Args[1] == "i" {
		f, err := os.Create("cpu.pprof")
		if err != nil {
			panic(err)
		}

		if err := pprof.StartCPUProfile(f); err != nil {
			panic(err)
		}
		defer pprof.StopCPUProfile()

		b, e := ioutil.ReadFile("./input")
		if e != nil {
			panic(e)
		}
		sc = bufio.NewScanner(strings.NewReader(strings.Replace(string(b), " ", "\n", -1)))
	}
	sc.Buffer(make([]byte, 10000), 2000000)

	s := nb()
	t := nb()
	tlen := len(t)

	tuintlistlen := (tlen-1)/64 + 1

	tuintlist = convertUintList(t, tuintlistlen)
	baseuintlist := convertUintList(s[0:tlen], tuintlistlen)
	o = tlen

	um := NewUintMap()
	tmpum := NewUintMap()
	for i, v := range baseuintlist {
		if i == 0 {
			tmpum = um.AddChild(v)
			continue
		}
		tmpum = tmpum.AddChild(v)
	}
	o = tmpum.GetDistanse()

	slen := len(s)
	alllen := slen - tlen + 1
	for ai := 0; ai < alllen; ai++ {
		check(um, s[ai:ai+tlen])
	}

	fmt.Println(o)
	_ = wtr.Flush()
}

func check(um *UintMap, b []byte) {

	hasNext := len(b) > 64
	if hasNext {
		newum := um.AddChild(convertUint(b[:64]))
		if newum != nil {
			check(newum, b[64:])
		}
		return
	}
	newum := um.AddChild(convertUint(b))
	if newum != nil && newum.GetDistanse() < o {
		o = newum.GetDistanse()
	}

}
func convertUint(b []byte) uint64 {
	var ui uint64
	for _, c := range b {
		d := c - '0'
		ui *= uint64(2)
		n1 := ui + uint64(d)
		ui = n1
	}
	return ui
}

func convertUintList(b []byte, uintlen int) []uint64 {
	uintindex := -1
	uintlist := make([]uint64, uintlen, uintlen)
	for i, c := range b {

		if i%64 == 0 {
			uintindex++
		}

		d := c - '0'
		uintlist[uintindex] *= uint64(2)
		n1 := uintlist[uintindex] + uint64(d)
		uintlist[uintindex] = n1
	}
	return uintlist
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

type UintMap struct {
	distanse   int
	generation int
	childlen   map[int]*UintMap
}

func NewUintMap() *UintMap {
	return &UintMap{
		childlen: make(map[int]*UintMap),
	}
}

func (um *UintMap) AddChild(ui uint64) *UintMap {
	i := int(ui)
	if v, ok := um.childlen[i]; ok {
		return v
	}
	distanse := um.Distance(ui)
	if distanse > o {
		return nil
	}
	newum := NewUintMap()
	newum.generation = um.generation + 1
	newum.distanse = distanse
	um.childlen[i] = newum
	return newum
}

func (um *UintMap) GetDistanse() int {
	return um.distanse
}

func (um *UintMap) Distance(ui uint64) int {
	return bits.OnesCount64(ui^tuintlist[um.generation]) + um.distanse
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
