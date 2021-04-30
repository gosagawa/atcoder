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

	n := ni()
	cs := make([]int, n, n)
	cm := make(map[int]int)
	cCount := 0
	for i := 0; i < n; i++ {
		c := ni()
		if _, ok := cm[c]; !ok {
			cm[c] = cCount
			cCount++
		}
		cs[i] = cm[c]
	}
	m := make([][]int, n, n)
	for i := 0; i < n-1; i++ {
		a := ni()
		b := ni()
		m[a-1] = append(m[a-1], b)
		m[b-1] = append(m[b-1], a)
	}
	tree := NewTree(m, cs, n, cCount)

	for i, v := range tree.Goods {
		if v {
			out(i + 1)
		}
	}
	e := wtr.Flush()
	if e != nil {
		panic(e)
	}
}

type Tree struct {
	Goods []bool
}

type node struct {
	key int
	c   int
	cm  map[int]struct{}
}

func NewTree(m [][]int, cs []int, n int, cn int) *Tree {

	cm := make([]bool, cn, cn)
	cm[cs[0]] = true
	goods := make([]bool, n, n)
	goods[0] = true
	for _, v := range m[0] {
		newNode(m, cs, cm, 1, v, goods)
	}
	tree := &Tree{
		Goods: goods,
	}
	return tree
}

func newNode(m [][]int, cs []int, cm []bool, parentKey, key int, goods []bool) *node {
	/*
		out("---------------")
		out(key)
		out(cm)
		out(cs[key-1])
	*/

	if !cm[cs[key-1]] {
		goods[key-1] = true
		cm[cs[key-1]] = true
		defer func() {
			cm[cs[key-1]] = false
		}()
	}

	node := &node{key: key}
	for _, v := range m[key-1] {
		if v == parentKey {
			continue
		}
		newNode(m, cs, cm, key, v, goods)
	}
	return node
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

func out(v interface{}) {
	_, e := fmt.Fprintln(wtr, v)
	if e != nil {
		panic(e)
	}
}
