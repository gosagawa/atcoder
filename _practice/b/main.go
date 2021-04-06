package main

import (
	"fmt"
)

func main() {

	var N, Q int
	fmt.Scanf("%d %d", &N, &Q)

	allChar := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	s := []rune(allChar[:N])

	c := new(DefaultComparison)

	if N == 5 {
		s = SortForFive(c, s)
	} else {
		s = mergeSort(c, s)
	}

	fmt.Printf("! %v\n", string(s))
}

func mergeSort(c Comparison, s []rune) []rune {
	var num = len(s)

	if num == 1 {
		return s
	}

	middle := int(num / 2)
	var (
		left  = make([]rune, middle)
		right = make([]rune, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = s[i]
		} else {
			right[i-middle] = s[i]
		}
	}

	return merge(c, mergeSort(c, left), mergeSort(c, right))
}

func merge(c Comparison, left, right []rune) (result []rune) {
	result = make([]rune, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {

		if !c.Gt(left[0], right[0]) {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}

type Comparison interface {
	Gt(a, b rune) bool
}

type DefaultComparison struct {
}

func (c *DefaultComparison) Gt(a, b rune) bool {
	var ans string
	fmt.Printf("? %v %v\n", string(a), string(b))
	fmt.Scanf("%s", &ans)
	if ans == ">" {
		return true
	}
	return false
}

func SortForFive(c Comparison, a []rune) []rune {

	if c.Gt(a[0], a[1]) {
		a = swap(a, 0, 1)
	}

	if c.Gt(a[2], a[3]) {
		a = swap(a, 2, 3)
	}

	if c.Gt(a[0], a[2]) {
		a = swap(a, 0, 2)
		a = swap(a, 1, 3)
	}

	if c.Gt(a[2], a[4]) {
		if c.Gt(a[1], a[2]) {
			a = swap(a, 1, 4)
			if c.Gt(a[0], a[1]) {
				a = swap(a, 0, 1)
			}
			if c.Gt(a[3], a[4]) {
				a = swap(a, 3, 4)
			}
		} else {
			a = swap(a, 3, 4)
			a = swap(a, 2, 3)
			if c.Gt(a[1], a[2]) {
				a = swap(a, 1, 2)
			}
			if c.Gt(a[0], a[1]) {
				a = swap(a, 0, 1)
			}

		}
	} else {
		if c.Gt(a[3], a[4]) {
			a = swap(a, 3, 4)
		}
		if c.Gt(a[1], a[3]) {
			a = swap(a, 1, 2)
			a = swap(a, 2, 3)
			if c.Gt(a[3], a[4]) {
				a = swap(a, 3, 4)
			}
		} else {
			if c.Gt(a[1], a[2]) {
				a = swap(a, 1, 2)
			}
		}
	}

	return a
}

func swap(runes []rune, a, b int) []rune {
	tmp := runes[a]
	runes[a] = runes[b]
	runes[b] = tmp
	return runes
}
