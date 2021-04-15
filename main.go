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
	for i := 0; i < n; i++ {
		cs[i] = ni()
	}
	m := make([][]int, n, n)
	for i := 0; i < n-1; i++ {
		a := ni()
		b := ni()
		m[a-1] = append(m[a-1], b)
		m[b-1] = append(m[b-1], a)
	}
	tree := NewTree(m, cs)

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
	root  *node
	Goods []bool
}

type node struct {
	key      int
	c        int
	cm       map[int]struct{}
	children []*node
}

func NewTree(m [][]int, cs []int) *Tree {
	cm := make(map[int]struct{})
	cm[cs[0]] = struct{}{}
	goods := make([]bool, len(cs), len(cs))
	goods[0] = true
	rootnode := &node{key: 1, c: cs[1], cm: cm}
	for _, v := range m[0] {
		rootnode.children = append(rootnode.children, newNode(m, cs, cm, 1, v, goods))
	}
	tree := &Tree{
		root:  rootnode,
		Goods: goods,
	}
	return tree
}

func newNode(m [][]int, cs []int, cm map[int]struct{}, parentKey, key int, goods []bool) *node {
	/*
		out(key)
		out(cm)
		out(cs[key-1])
	*/
	if _, ok := cm[cs[key-1]]; !ok {
		goods[key-1] = true
	}
	newCm := make(map[int]struct{})
	newCm[cs[key-1]] = struct{}{}
	for v := range cm {
		newCm[v] = struct{}{}
	}
	node := &node{key: key}
	for _, v := range m[key-1] {
		if v == parentKey {
			continue
		}
		node.children = append(node.children, newNode(m, cs, newCm, key, v, goods))
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
