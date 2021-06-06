package main

import (
	"fmt"
)

func main() {
	h := newHeap([]int{1, 2, 3, 4, 5, 6, 7})
	h.push(9)
	fmt.Println(h)
	p := h.pop()
	fmt.Println(p)
	fmt.Println(h)
}

type myheap struct {
	sl   []int
	size int
}

func newHeap(sl []int) *myheap {

	h := &myheap{
		sl:   sl,
		size: len(sl),
	}
	for i := h.size/2 - 1; i >= 0; i-- {
		h.heapify(i)
	}

	return h
}

func (h *myheap) heapify(i int) {
	child1 := i*2 + 1
	if child1 > h.size {
		return
	}

	if child1+1 < h.size && h.sl[child1+1] > h.sl[child1] {
		child1++
	}

	if h.sl[child1] <= h.sl[i] {
		return
	}

	h.sl[i], h.sl[child1] = h.sl[child1], h.sl[i]

	h.heapify(child1)
}

func (h *myheap) push(x int) {
	h.sl = append(h.sl, x)
	h.size++
	i := h.size - 1
	for {
		if i == 0 {
			break
		}
		p := (i - 1) / 2
		if h.sl[p] >= x {
			break
		}
		h.sl[i] = h.sl[p]
		i = p
	}
	h.sl[i] = x
}

func (h *myheap) top() int {
	if h.size > 0 {
		return h.sl[0]
	}
	return -1
}

func (h *myheap) pop() int {
	if h.size == 0 {
		return -1
	}
	r := h.sl[0]
	x := h.sl[h.size-1]
	h.sl = h.sl[:h.size-1]
	h.size--
	i := 0
	for {
		if i*2+1 >= h.size {
			break
		}
		child1 := i*2 + 1
		child2 := i*2 + 2
		if child2 < h.size && h.sl[child2] > h.sl[child1] {
			child1 = child2
		}
		if h.sl[child1] <= x {
			break
		}
		h.sl[i] = h.sl[child1]
		i = child1
	}
	h.sl[i] = x
	return r
}
