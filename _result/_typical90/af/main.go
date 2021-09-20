package main

import (
	"bufio"
	"container/heap"
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

	o := inf

	n := ni()
	as := make([][]int, n)
	ms := make([][]bool, n)
	cs := make([]int, n)
	for i := 0; i < n; i++ {
		cs[i] = i
		as[i] = make([]int, n)
		ms[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			as[i][j] = ni()
		}
	}
	m := ni()
	for i := 0; i < m; i++ {
		x, y := ni2()
		x--
		y--
		ms[x][y] = true
		ms[y][x] = true
	}
	for {
		r := 0
		for i := 0; i < n; i++ {
			if i != n-1 && ms[cs[i]][cs[i+1]] {
				r = inf
				break
			}
			r += as[cs[i]][i]
		}
		o = min(o, r)

		if !nextPermutation(sort.IntSlice(cs)) {
			break
		}
	}
	if o == inf {
		out(-1)
		return
	}
	out(o)

}

// ==================================================
// init
// ==================================================

const inf = math.MaxInt64
const mod = 1000000007

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

func nftoi(decimalLen int) int {
	sc.Scan()
	s := sc.Text()

	r := 0
	minus := strings.Split(s, "-")
	isMinus := false
	if len(minus) == 2 {
		s = minus[1]
		isMinus = true
	}

	t := strings.Split(s, ".")
	i := atoi(t[0])
	r += i * pow(10, decimalLen)
	if len(t) > 1 {
		i = atoi(t[1])
		i *= pow(10, decimalLen-len(t[1]))
		r += i
	}
	if isMinus {
		return -r
	}
	return r
}

func atoi(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return i
}

func btoi(b byte) int {
	return atoi(string(b))
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

func getAngle(x, y float64) float64 {
	return math.Atan2(y, x) * 180 / math.Pi
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

func modcombination(n int, k int) int {
	if k > n || k <= 0 {
		panic(fmt.Sprintf("invalid param n:%v k:%v", n, k))
	}
	if n-k < k {
		k = n - k
	}
	v := 1
	for i := 0; i < k; i++ {
		v = mmul(v, n-i)
		v = mdiv(v, i+1)
	}
	return v
}

func factorial(n int) int {
	return permutation(n, n-1)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func divisor(n int) []int {
	sqrtn := int(math.Sqrt(float64(n)))
	c := 2
	divisor := []int{}
	for {
		if n%2 != 0 {
			break
		}
		divisor = append(divisor, 2)
		n /= 2
	}
	c = 3
	for {
		if n%c == 0 {
			divisor = append(divisor, c)
			n /= c
		} else {
			c += 2
			if c > sqrtn {
				break
			}
		}
	}
	if n != 1 {
		divisor = append(divisor, n)
	}
	return divisor
}

// ==================================================
// mod
// ==================================================

func madd(a, b int) int {
	a += b
	if a < 0 {
		a += mod
	} else if a >= mod {
		a -= mod
	}
	return a
}

func mmul(a, b int) int {
	return a * b % mod
}

func mdiv(a, b int) int {
	a %= mod
	return a * modinvfermat(b, mod) % mod
}

func modinv(a, m int) int {
	b := m
	u := 1
	v := 0
	for b != 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		u -= t * v
		u, v = v, u
	}
	u %= m
	if u < 0 {
		u += m
	}
	return u
}

// m only allow prime number
func modinvfermat(a, m int) int {
	res := 1
	n := m - 2
	for n > 0 {
		if n&1 == 1 {
			res = res * a % m
		}
		a = a * a % m
		n >>= 1
	}
	return res
}

func mpow(m, n int) int {
	res := 1
	for n != 0 {
		if n&1 == 1 {
			res = mmul(res, m)
		}
		m, n = mmul(m, m), n>>1
	}
	return res
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
// sort
// ==================================================

func sorti(sl []int) {
	sort.Sort(sort.IntSlice(sl))
}

func sortir(sl []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(sl)))
}

func sorts(sl []string) {
	sort.Slice(sl, func(i, j int) bool {
		return sl[i] < sl[j]
	})
}

func sort2ar(sl [][2]int, key1, key2 int) {
	sort.Slice(sl, func(i, j int) bool {
		if sl[i][key1] == sl[j][key1] {
			return sl[i][key2] < sl[j][key2]
		}
		return sl[i][key1] < sl[j][key1]
	})
}

// ==================================================
// slice
// ==================================================

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

// ==================================================
// heap
// ==================================================

/*
  h := &int2dHeap{&int2d{dist[r], r}}
  heap.Init(h)
  v := heap.Pop(h).(*int2d)
  heap.Push(h, &int2d{x, y})
*/

type int2d [2]int

type int2dHeap []*int2d

func (h int2dHeap) Len() int           { return len(h) }
func (h int2dHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h int2dHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *int2dHeap) Push(x interface{}) {
	*h = append(*h, x.(*int2d))
}

func (h *int2dHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	old[n-1] = nil
	*h = old[0 : n-1]
	return x
}

func (h *int2dHeap) IsEmpty() bool {
	return h.Len() == 0
}

type pq struct {
	arr   []interface{}
	comps []compFunc
}

/*
	graph.comps = []compFunc{
		func(p, q interface{}) int {
			if p.(state).score < q.(state).score {
				return -1
			} else if p.(state).score == q.(state).score {
				return 0
			}
			return 1
		},
	}
*/
type compFunc func(p, q interface{}) int

func newpq(comps []compFunc) *pq {
	return &pq{
		comps: comps,
	}
}

func (pq pq) Len() int {
	return len(pq.arr)
}

func (pq pq) Swap(i, j int) {
	pq.arr[i], pq.arr[j] = pq.arr[j], pq.arr[i]
}

func (pq pq) Less(i, j int) bool {
	for _, comp := range pq.comps {
		result := comp(pq.arr[i], pq.arr[j])
		switch result {
		case -1:
			return true
		case 1:
			return false
		case 0:
			continue
		}
	}
	return true
}

func (pq *pq) Push(x interface{}) {
	pq.arr = append(pq.arr, x)
}

func (pq *pq) Pop() interface{} {
	n := pq.Len()
	item := pq.arr[n-1]
	pq.arr = pq.arr[:n-1]
	return item
}

func (pq *pq) Top() interface{} {
	n := pq.Len()
	return pq.arr[n-1]
}

func (pq *pq) IsEmpty() bool {
	return pq.Len() == 0
}

// ==================================================
// cusum2d
// ==================================================

type cusum2d struct {
	s [][]int
}

func newCusum2d(n, m int) *cusum2d {
	c := &cusum2d{}
	c.s = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		c.s[i] = make([]int, m+1)
	}
	return c
}
func (c *cusum2d) set(x, y, add int) {
	c.s[x+1][y+1] = c.s[x+1][y] + c.s[x][y+1] - c.s[x][y]
	c.s[x+1][y+1] += add
}

func (c *cusum2d) get(x1, y1, x2, y2 int) int {
	return c.s[x2][y2] + c.s[x1][y1] - c.s[x1][y2] - c.s[x2][y1]
}

// ==================================================
// union find
// ==================================================

type unionFind struct {
	par []int
}

func newUnionFind(n int) *unionFind {
	u := &unionFind{
		par: make([]int, n),
	}
	for i := range u.par {
		u.par[i] = -1
	}
	return u
}

func (u *unionFind) root(x int) int {
	if u.par[x] < 0 {
		return x
	}
	u.par[x] = u.root(u.par[x])
	return u.par[x]
}

func (u *unionFind) unite(x, y int) {
	x = u.root(x)
	y = u.root(y)
	if x == y {
		return
	}
	if u.size(x) < u.size(y) {
		x, y = y, x
	}
	u.par[y] += u.par[x]
	u.par[x] = y
}

func (u *unionFind) issame(x, y int) bool {
	if u.root(x) == u.root(y) {
		return true
	}
	return false
}

func (u *unionFind) size(x int) int {
	return -u.par[u.root(x)]
}

// ==================================================
// bit
// ==================================================

type bit struct {
	n int
	b []int
}

func newbit(n int) *bit {
	return &bit{
		n: n + 1,
		b: make([]int, n+1),
	}
}

func (b *bit) add(i, x int) {
	for i++; i < b.n && i > 0; i += i & -i {
		b.b[i] += x
	}
}

func (b *bit) sum(i int) int {
	ret := 0
	for i++; i > 0; i -= i & -i {
		ret += b.b[i]
	}
	return ret
}

func (b *bit) rangesum(l, r int) int {
	return b.sum(r-1) - b.sum(l-1)
}

func (b *bit) lowerBound(x int) int {
	idx, k := 0, 1
	for k < b.n {
		k <<= 1
	}
	for k >>= 1; k > 0; k >>= 1 {
		if idx+k < b.n && b.b[idx+k] < x {
			x -= b.b[idx+k]
			idx += k
		}
	}
	return idx
}

// ==================================================
// segment tree
// ==================================================

type stree struct {
	n    int
	b    []int
	lazy []int
	def  int
	cmp  func(i, j int) int
	culc func(i, j int) int
}
type streeculctype int

var stadd streeculctype = 1
var stset streeculctype = 2

type streeminmmax int

var stmin streeminmmax = 1
var stmax streeminmmax = 2

/*
s := newstree(n,stmin|stmax,stset|stadd)
s.set(i,x)
s.add(i,x)
s.rc(l,r,x)
result1 := s.query(l,r)
result2 := s.findrightest(l,r,x)
result3 := s.findlefttest(l,r,x)
*/
func newstree(n int, minmax streeminmmax, ctype streeculctype) *stree {
	tn := 1
	for tn < n {
		tn *= 2
	}
	s := &stree{
		n:    tn,
		b:    make([]int, 2*tn-1),
		lazy: make([]int, 2*tn-1),
	}
	switch minmax {
	case stmin:
		s.def = inf
		for i := 0; i < 2*tn-1; i++ {
			s.b[i] = s.def
			s.lazy[i] = s.def
		}
		s.cmp = func(i, j int) int {
			return min(i, j)
		}
	case stmax:
		s.cmp = func(i, j int) int {
			return max(i, j)
		}
	}
	switch ctype {
	case stadd:
		s.culc = func(i, j int) int {
			return i + j
		}
	case stset:
		s.culc = func(i, j int) int {
			return j
		}
	}
	return s
}

func (s *stree) eval(k int) {
	if s.lazy[k] == s.def {
		return
	}
	if k < s.n-1 {
		s.lazy[k*2+1] = s.culc(s.lazy[k*2+1], s.lazy[k])
		s.lazy[k*2+2] = s.culc(s.lazy[k*2+2], s.lazy[k])
	}
	s.b[k] = s.culc(s.b[k], s.lazy[k])
	s.lazy[k] = s.def
}

func (s *stree) add(i, x int) {
	i += s.n - 1
	s.b[i] += x

	for i > 0 {
		i = (i - 1) / 2
		s.b[i] = s.cmp(s.b[i*2+1], s.b[i*2+2])
	}
}

func (s *stree) set(i, x int) {
	i += s.n - 1
	s.b[i] = x

	for i > 0 {
		i = (i - 1) / 2
		s.b[i] = s.cmp(s.b[i*2+1], s.b[i*2+2])
	}
}

// range culc
func (s *stree) rc(a, b, x int) {
	s.rcsub(a, b, x, 0, 0, s.n)
}

func (s *stree) rcsub(a, b, x, k, l, r int) {
	s.eval(k)
	if a <= l && r <= b {
		s.lazy[k] = s.culc(s.lazy[k], x)
		s.eval(k)
	} else if l < b && a < r {
		s.rcsub(a, b, x, k*2+1, l, (l+r)/2)
		s.rcsub(a, b, x, k*2+2, (l+r)/2, r)
		s.b[k] = s.cmp(s.b[k*2+1], s.b[k*2+2])
	}
}

func (s *stree) query(a, b int) int {
	return s.querysub(a, b, 0, 0, s.n)
}

func (s *stree) querysub(a, b, k, l, r int) int {
	s.eval(k)
	if r <= a || b <= l {
		return s.def
	}
	if a <= l && r <= b {
		return s.b[k]
	}
	return s.cmp(
		s.querysub(a, b, k*2+1, l, (l+r)/2),
		s.querysub(a, b, k*2+2, (l+r)/2, r),
	)
}

func (s *stree) findrightest(a, b, x int) int {
	return s.findrightestsub(a, b, x, 0, 0, s.n)
}

func (s *stree) findrightestsub(a, b, x, k, l, r int) int {
	if s.b[k] > x || r <= a || b <= l {
		return a - 1
	} else if k >= s.n-1 {
		return k - s.n + 1
	}
	vr := s.findrightestsub(a, b, x, 2*k+2, (l+r)/2, r)
	if vr != a-1 {
		return vr
	}
	return s.findrightestsub(a, b, x, 2*k+1, l, (l+r)/2)
}

func (s *stree) findleftest(a, b, x int) int {
	return s.findleftestsub(a, b, x, 0, 0, s.n)
}

func (s *stree) findleftestsub(a, b, x, k, l, r int) int {
	if s.b[k] > x || r <= a || b <= l {
		return b
	} else if k >= s.n-1 {
		return k - s.n + 1
	}
	vl := s.findleftestsub(a, b, x, 2*k+1, l, (l+r)/2)
	if vl != b {
		return vl
	}
	return s.findleftestsub(a, b, x, 2*k+2, (l+r)/2, r)
}

func (s *stree) debug() {
	l := []string{}
	t := 2
	out("data")
	for i := 0; i < 2*s.n-1; i++ {
		if i+1 == t {
			t *= 2
			out(strings.Join(l, " "))
			l = []string{}
		}
		if s.b[i] == inf {
			l = append(l, "∞")
		} else {
			l = append(l, strconv.Itoa(s.b[i]))
		}
	}
	out(strings.Join(l, " "))
	out("lazy")
	l = []string{}
	t = 2
	for i := 0; i < 2*s.n-1; i++ {
		if i+1 == t {
			t *= 2
			out(strings.Join(l, " "))
			l = []string{}
		}
		if s.lazy[i] == inf {
			l = append(l, "∞")
		} else {
			l = append(l, strconv.Itoa(s.lazy[i]))
		}
	}
	out(strings.Join(l, " "))
}

// ==================================================
// graph
// ==================================================

type edge struct {
	from int
	to   int
	cost int
}

type state struct {
	score int
	node  int
}

type graph struct {
	size         int
	edges        [][]edge
	starts       []state
	comps        []compFunc
	defaultScore int
}

func newgraph(size int, edges [][]edge) *graph {
	graph := &graph{
		size:  size,
		edges: edges,
	}

	graph.defaultScore = inf
	graph.comps = []compFunc{
		func(p, q interface{}) int {
			if p.(state).score < q.(state).score {
				return -1
			} else if p.(state).score == q.(state).score {
				return 0
			}
			return 1
		},
	}
	return graph
}

/*
	v, e := ni2()
	edges := make([][]edge, v)

	for i := 0; i < e; i++ {
		s, t, c := ni3()
		s--
		t--
		edges[s] = append(edges[s], edge{to: t, cost: c})
		edges[t] = append(edges[t], edge{to: s, cost: c})
	}

	graph := newgraph(v, edges)
	graph.setStart(0)
	dist := graph.dijkstra(0)
*/

func (g *graph) dijkstra(start int) []int {

	g.starts = []state{{node: start}}

	score := make([]int, g.size)
	for i := 0; i < g.size; i++ {
		score[i] = g.defaultScore
	}
	que := newpq(g.comps)
	for _, start := range g.starts {
		score[start.node] = start.score
		heap.Push(que, start)
	}

	for !que.IsEmpty() {
		st := heap.Pop(que).(state)
		if st.score > score[st.node] {
			continue
		}
		for _, edge := range g.edges[st.node] {
			newScore := st.score + edge.cost
			if score[edge.to] > newScore {
				score[edge.to] = newScore
				heap.Push(que, state{score: newScore, node: edge.to})
			}
		}
	}
	return score
}

func (g *graph) floydWarshall() ([][]int, bool) {

	score := make([][]int, g.size)
	for i := 0; i < g.size; i++ {
		score[i] = make([]int, g.size)
		for j := 0; j < g.size; j++ {
			if i == j {
				score[i][j] = 0
			} else {
				score[i][j] = g.defaultScore
			}
		}
		for _, edge := range g.edges[i] {
			score[i][edge.to] = edge.cost
		}
	}
	for k := 0; k < g.size; k++ {
		for i := 0; i < g.size; i++ {
			for j := 0; j < g.size; j++ {
				if score[i][k] == g.defaultScore || score[k][j] == g.defaultScore {
					continue
				}
				score[i][j] = min(score[i][j], score[i][k]+score[k][j])
			}
		}
	}

	for k := 0; k < g.size; k++ {
		if score[k][k] < 0 {
			return nil, true
		}
	}

	return score, false
}

/*
	v, e := ni2()
	edges := make([][]edge, 1)
	edges[0] = make([]edge, e)

	for i := 0; i < e; i++ {
		s, t, d := ni3()
		edges[0][i] = edge{from: s, to: t, cost: d}
	}

	graph := newgraph(v, edges)

	o = graph.kruskal()
*/
func (g *graph) kruskal() int {

	sort.Slice(g.edges[0], func(i, j int) bool { return g.edges[0][i].cost < g.edges[0][j].cost })

	e := len(g.edges[0])

	uf := newUnionFind(g.size)
	r := 0
	for i := 0; i < e; i++ {
		edge := g.edges[0][i]
		if uf.issame(edge.from, edge.to) {
			continue
		}
		r += edge.cost
		uf.unite(edge.from, edge.to)
	}

	return r
}
