package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"math/bits"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var wtr = bufio.NewWriter(os.Stdout)

func main() {

	defer flush()

	o := 0
	f := func(s string) int {
		r := 0
		minus := strings.Split(s, "-")
		isMinus := false
		if len(minus) == 2 {
			s = minus[1]
			isMinus = true
		}

		t := strings.Split(s, ".")
		i, e := strconv.Atoi(t[0])
		if e != nil {
			panic(e)
		}
		r += i * 10000
		if len(t) > 1 {
			i, e := strconv.Atoi(t[1])
			if e != nil {
				panic(e)
			}
			if len(t[1]) == 1 {
				i *= 1000
			} else if len(t[1]) == 2 {
				i *= 100
			} else if len(t[1]) == 3 {
				i *= 10
			}
			r += i
		}
		if isMinus {
			return -r
		}
		return r
	}
	x := f(ns())
	y := f(ns())
	r := f(ns())
	r2 := r * r
	//out(x, y, r)
	minx := modgt(x-r, 10000)
	maxx := modlt(x+r, 10000)
	miny := modgt(y-r, 10000)
	maxy := modlt(y+r, 10000)
	//out(minx, maxx, miny, maxy)

	minceny := modlt(y, 10000)
	maxceny := modgt(y, 10000)
	//out(minceny, maxceny)

	var tymin, tymax int
	if (minx-x)*(minx-x)+(minceny-y)*(minceny-y) <= r2 {
		tymax = minceny
		tymin = minceny
	} else if (minx-x)*(minx-x)+(maxceny-y)*(maxceny-y) <= r2 {
		tymax = maxceny
		tymin = maxceny
	} else {
		minx += 10000
		if (minx-x)*(minx-x)+(minceny-y)*(minceny-y) <= r2 {
			tymax = minceny
			tymin = minceny
		} else if (minx-x)*(minx-x)+(maxceny-y)*(maxceny-y) <= r2 {
			tymax = maxceny
			tymin = maxceny
		} else {
			out(0)
			return
		}
	}

	for cx := minx; cx <= maxx; cx += 10000 {
		hasTy := false
		if cx <= x {
			for cy := tymax; cy <= maxy; cy += 10000 {
				if (cx-x)*(cx-x)+(cy-y)*(cy-y) <= r2 {
					hasTy = true
					tymax = cy
				} else {
					break
				}
			}
			for cy := tymin; cy >= miny; cy -= 10000 {
				if (cx-x)*(cx-x)+(cy-y)*(cy-y) <= r2 {
					hasTy = true
					tymin = cy
				} else {
					break
				}
			}
		} else {
			for cy := tymax; cy >= miny; cy -= 10000 {
				if (cx-x)*(cx-x)+(cy-y)*(cy-y) <= r2 {
					hasTy = true
					tymax = cy
					break
				}
			}
			for cy := tymin; cy <= maxy; cy += 10000 {
				if (cx-x)*(cx-x)+(cy-y)*(cy-y) <= r2 {
					hasTy = true
					tymin = cy
					break
				}
			}
		}
		//out(cx, tymax, tymin)
		if hasTy {
			o += (tymax-tymin)/10000 + 1
		}

	}

	out(o)
}

func modgt(i, m int) int {
	if i%m == 0 {
		return i
	}
	if i >= 0 {
		return i - i%m + m
	}
	return i - i%m
}

func modlt(i, m int) int {
	if i%m == 0 {
		return i
	}
	if i >= 0 {
		return i - i%m
	}
	return i - i%m - m
}

// ==================================================
// init
// ==================================================

const inf = math.MaxInt64

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
}

// ==================================================
// io
// ==================================================

func ni() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func ni2() (int, int) {
	return ni(), ni()
}

func ni3() (int, int, int) {
	return ni(), ni(), ni()
}

func ni4() (int, int, int, int) {
	return ni(), ni(), ni(), ni()
}

func nis(n int) sort.IntSlice {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = ni()
	}
	return sort.IntSlice(a)
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

// ==================================================
// num
// ==================================================

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func pow2(a int) int {
	return int(math.Pow(2, float64(a)))
}

func ch(cond bool, ok, ng int) int {
	if cond {
		return ok
	}
	return ng
}

func mul(a, b int) (int, int) {
	if a < 0 {
		a, b = -a, -b
	}
	if a == 0 || b == 0 {
		return 0, 0
	} else if a > 0 && b > 0 && a > math.MaxInt64/b {
		return 0, +1
	} else if a < math.MinInt64/b {
		return 0, -1
	}
	return a * b, 0
}

func permutation(n int, k int) int {
	if k > n || k <= 0 {
		panic(fmt.Sprintf("invalid param n:%v k:%v", n, k))
	}
	v := 1
	for i := 0; i < k; i++ {
		v *= (n - i)
	}
	return v
}

/*
	for {

		// Do something

		if !nextPermutation(sort.IntSlice(x)) {
			break
		}
	}
*/
func nextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}

func combination(n int, k int) int {
	if n-k < k {
		k = n - k
	}
	v := 1
	for i := 0; i < k; i++ {
		v *= (n - i)
		v /= (i + 1)
	}
	return v
}

func factorial(n int) int {
	return permutation(n, n-1)
}

// ==================================================
// binarysearch
// ==================================================
/*
	f := func(c int) bool {
		return false
	}
*/
func bs(ok, ng int, f func(int) bool) int {
	for abs(ok-ng) > 1 {
		mid := (ok + ng) / 2

		if f(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}

	return ok
}

// ==================================================
// bit
// ==================================================

func hasbit(a int, n int) bool {
	return (a>>uint(n))&1 == 1
}

func nthbit(a int, n int) int {
	return int((a >> uint(n)) & 1)
}

func popcount(a int) int {
	return bits.OnesCount(uint(a))
}

func xor(a, b bool) bool { return a != b }

// ==================================================
// string
// ==================================================

func toLowerCase(s string) string {
	return strings.ToLower(s)
}

func toUpperCase(s string) string {
	return strings.ToUpper(s)
}

func isLower(b byte) bool {
	return 'a' <= b && b <= 'z'
}

func isUpper(b byte) bool {
	return 'A' <= b && b <= 'Z'
}

// ==================================================
// slice
// ==================================================

func sorti(sl []int) {
	sort.Sort(sort.IntSlice(sl))
}

func sortir(sl []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(sl)))
}

func reverse(sl []interface{}) {
	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}
}

func reversei(sl []int) {
	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}
}

func uniquei(sl []int) []int {
	hist := map[int]bool{}
	j := 0
	for i := 0; i < len(sl); i++ {
		if hist[sl[i]] {
			continue
		}

		a := sl[i]
		sl[j] = a
		hist[a] = true
		j++
	}
	return sl[:j]
}

func delIdx(idx int, L []int) []int {
	r := []int{}
	r = append(r, L[:idx]...)
	r = append(r, L[idx+1:]...)
	return r
}

// ==================================================
// point
// ==================================================

type point struct {
	x int
	y int
}

type pointf struct {
	x float64
	y float64
}

func (p point) isValid(x, y int) bool {
	return 0 <= p.x && p.x < x && 0 <= p.y && p.y < y
}

func pointAdd(a, b point) point {
	return point{a.x + b.x, a.y + b.y}
}

func pointDist(a, b point) float64 {
	return math.Sqrt(float64((a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y)))
}

func pointfDist(a, b pointf) float64 {
	return math.Sqrt((a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y))
}

type pointQueue struct {
	pt []point
}

func (q *pointQueue) push(pt point) {
	q.pt = append(q.pt, pt)
}

func (q *pointQueue) pop() (pt point) {
	pt = q.pt[0]
	q.pt = q.pt[1:]
	return
}

func (q *pointQueue) len() int {
	return len(q.pt)
}
