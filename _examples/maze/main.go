package main

import (
	"bufio"
	"container/list"
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
}
func asPoint() {
	defer flush()

	o := 0
	h, w := ni2()
	m := make([][]bool, h)
	dist := make([][]int, h)
	for i := 0; i < h; i++ {
		m[i] = make([]bool, w)
		dist[i] = make([]int, w)

		// get from string
		s := ns()
		for j, v := range s {
			dist[i][j] = -1
			if string(v) == "." {
				m[i][j] = true
			}
		}

		// get from int
		/*
			for j := 0; j < w; j++ {
				dist[i][j] = -1
				if ni() == 1 {
					m[i][j] = true
				}
			}
		*/
	}

	q := list.New()
	q.PushBack(point{0, 0})
	dist[0][0] = 0
	e := q.Front()
	for e != nil {
		p := e.Value.(point)
		dr := []point{{0, -1}, {-1, 0}, {1, 0}, {0, 1}}
		for _, d := range dr {
			np := pointAdd(p, d)
			if !np.isValid(h, w) || !m[np.x][np.y] || dist[np.x][np.y] != -1 {
				continue
			}
			dist[np.x][np.y] = dist[p.x][p.y] + 1
			q.PushBack(np)
		}
		e = e.Next()
	}

	out(o)
}

func asGraph() {

	defer flush()

	o := 0
	w, h := ni2()
	ns := make([][]bool, h)
	for i := 0; i < h; i++ {
		// get from int
		ns[i] = make([]bool, w)
		for j := 0; j < w; j++ {
			if ni() == 1 {
				ns[i][j] = true
			}
		}

		// get from string
		/*
			s := ns()
			for j, v := range s {
				if string(v) == "." {
					ns[i][j] = true
				}
			}
		*/
	}

	gr := make([][]int, h*w)
	dist := make([]int, h*w)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if !ns[i][j] {
				continue
			}
			b := point2val(i, j, w)
			dist[b] = -1
			if i > 0 {
				if ns[i-1][j] {
					gr[b] = append(gr[b], b-w)
				}
			}
			if i < h-1 {
				if ns[i+1][j] {
					gr[b] = append(gr[b], b+w)
				}
			}
			if j > 0 {
				if ns[i][j-1] {
					gr[b] = append(gr[b], b-1)
				}
			}
			if j < w-1 {
				if ns[i][j+1] {
					gr[b] = append(gr[b], b+1)
				}
			}
		}
	}

	//debug
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			b := point2val(i, j, w)
			fmt.Printf("%02d ", dist[b])
		}
		println()
	}

	//dfs
	var dfs func(i, c int, cs []bool)
	cs := make([]bool, h*w)
	dfs = func(i, c int, cs []bool) {
		if cs[i] {
			return
		}
		cs[i] = true
		o = max(o, c)
		c++
		for _, v := range gr[i] {
			dfs(v, c, cs)
		}
		cs[i] = false
	}
	dfs(0, 0, cs)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			b := point2val(i, j, w)
			dfs(b, 0, cs)
		}
	}

	//bfs
	q := list.New()
	var s, g int
	q.PushBack(s)
	dist[s] = 0
	e := q.Front()
	for e != nil {
		t := e.Value.(int)
		for _, v := range gr[t] {
			if dist[v] != -1 {
				continue
			}
			q.PushBack(v)
			dist[v] = dist[t] + 1
		}
		e = e.Next()
	}
	out(dist[g])

	out(o)
}

func point2val(x, y, w int) int {
	return y%w + x*w
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
	return permutation(n, k) / factorial(k)
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
