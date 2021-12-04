package main

import (
	"bufio"
	"container/heap"
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
	//defer profile.Start(profile.ProfilePath(".")).Stop()
	defer flush()

	o := 0

	n, k := ni2()
	if k == 1 {
		out(fib(n + 1))
		return
	}
	if k <= 6 && n <= 6 {
		var rec func(sl []int)
		rec = func(sl []int) {
			l := len(sl)
			t := inf
			for i := 0; i < l; i++ {
				t = min(t, sl[l-i-1])
				if t*(i+1) > k {
					return
				}
			}
			if l == n {
				o++
				return
			}
			sl = append(sl, 0)
			for i := 0; i <= k; i++ {
				sl[l] = i
				rec(sl)
			}
			sl = sl[:l]
		}
		rec([]int{})
		out(o)
		return
	}
	if k <= 100 && n <= 100 {
		dp := make([][][]int, k+1)
		for h := 0; h <= k; h++ {
			dp[h] = make([][]int, n+2)
			for l := 0; l <= n+1; l++ {
				dp[h][l] = make([]int, n+2)
				for r := 0; r <= n+1; r++ {
					if r == l-1 {
						dp[h][l][r] = 1
					}
					if h == k && l == r {
						dp[h][l][r] = 1
					}
				}
			}
		}
		for h := k - 1; h >= 0; h-- {
			for l := 1; l <= n; l++ {
				for r := l; r <= n; r++ {
					if h*(r-l+1) > k {
						continue
					}

					dp[h][l][r] = madd(dp[h][l][r], dp[h+1][l][r])
					for i := l; i <= r; i++ {
						dp[h][l][r] = madd(dp[h][l][r], mmul(dp[h][l][i-1], dp[h+1][i+1][r]))
					}
				}
			}
		}
		o = dp[0][1][n]
		out(o)
		return
	}
	if k <= 10000 && n <= 10000 {

		dp := make([][]int, k+1)
		for h := 0; h <= k; h++ {
			dp[h] = make([]int, n+1)
			for m := 0; m <= n+1; m++ {
				dp[h][0] = 1
				if h == k {
					dp[h][1] = 1
				}
			}
		}
		for h := k - 1; h >= 0; h-- {
			for m := 1; m <= n; m++ {
				if h*m > k {
					continue
				}
				for i := 1; i <= m+1; i++ {
					if m-i == -1 {
						dp[h][m] = madd(dp[h][m], dp[h+1][i-1])
					} else {
						dp[h][m] = madd(dp[h][m], mmul(dp[h+1][i-1], dp[h][m-i]))
					}
				}
			}
		}
		o = dp[0][n]
		out(o)
		return
	}

	c := NewConvolution(mod, 3)

	polynomialInverse := func(csl []int, l int) []int {
		n := len(csl)
		a := []int{1, 0}
		level := 0
		for {
			if (1 << level) >= l {
				break
			}
			cs := min(2<<level, n)
			p := c.Convolve(a, csl[:cs])
			q := make([]int, 2<<level)
			q[0] = 1
			for j := (1 << level); j < (2 << level); j++ {
				if j >= len(p) {
					break
				}
				q[j] = madd(q[j], -p[j])
			}
			a = c.Convolve(a, q)
			a = a[:min(len(a)-1, 4<<level)]
			level++
		}
		a = a[:l]
		return a
	}
	dp := make([][]int, k+1)
	dp[k] = []int{1, 1, 1}
	for i := k - 1; i >= 1; i-- {
		limit := min(k/i, n)
		csl := make([]int, len(dp[i+1]))
		csl[0] = 1
		for j := 1; j < len(dp[i+1]); j++ {
			csl[j] = madd(csl[j], -dp[i+1][j])
		}
		dp[i] = polynomialInverse(csl, limit+2)
	}

	s := min(k, n)
	track := []int{n + s + 1}
	for {
		if track[len(track)-1] < s+1 {
			break
		}
		track = append(track, track[len(track)-1]/2)
	}
	reversei(track)
	cl := make([]int, s+2)
	cl[0] = 1
	for i := 1; i < s+2; i++ {
		cl[i] = madd(cl[i], -dp[1][i])
	}
	gl := polynomialInverse(cl, s+2)
	reversei(cl)
	poly := make([]int, s+1)
	poly[track[0]] = 1
	for i := 1; i < len(track); i++ {
		poly = c.Convolve(poly, poly)
		if track[i]%2 == 1 {
			poly = append([]int{0}, poly...)
		} else {
			poly = append(poly, 0)
		}
		p1 := poly[s+1:]
		reversei(p1)
		p2 := c.Convolve(p1, gl)
		p2 = p2[:s+1]
		reversei(p2)
		p3 := c.Convolve(p2, cl)
		for j := 0; j < 2*s+1; j++ {
			poly[j] = madd(poly[j], -p3[j])
		}
		poly = poly[:s+1]

	}
	out(poly[s])
}

func fib(n int) int {
	b := make([][2][2]int, 65)
	b[0] = [2][2]int{
		[2]int{0, 1},
		[2]int{1, 1},
	}
	for i := 0; i < 64; i++ {
		b[i+1] = mulq(b[i], b[i])
	}
	r := [][2][2]int{}
	for i := 0; i < 64; i++ {
		if n&(1<<i) != 0 {
			r = append(r, b[i])
		}
	}
	a := r[0]
	for i := 1; i < len(r); i++ {
		a = mulq(a, r[i])
	}

	return a[1][1]
}

func mulq(a, b [2][2]int) [2][2]int {

	n00 := madd(mmul(a[0][0], b[0][0]), mmul(a[0][1], b[1][0]))
	n10 := madd(mmul(a[1][0], b[0][0]), mmul(a[1][1], b[1][0]))
	n01 := madd(mmul(a[0][0], b[0][1]), mmul(a[0][1], b[1][1]))
	n11 := madd(mmul(a[1][0], b[0][1]), mmul(a[1][1], b[1][1]))

	return [2][2]int{
		[2]int{n00, n01},
		[2]int{n10, n11},
	}
}

// ==================================================
// init
// ==================================================

const inf = math.MaxInt64
const mod = 998244353

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

func outYN(v bool) {
	if v {
		out("Yes")
	} else {
		out("No")
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

func itoa(i int) string {
	return strconv.Itoa(i)
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

func madd32(a, b int32) int32 {
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
func mmul32(a, b int32) int32 {
	return int32(int(a) * int(b) % mod)
}

func mdiv(a, b int) int {
	a %= mod
	return a * minvfermat(b, mod) % mod
}

func mpow(a, n, m int) int {
	if m == 1 {
		return 0
	}
	r := 1
	for n > 0 {
		if n&1 == 1 {
			r = r * a % m
		}
		a, n = a*a%m, n>>1
	}
	return r
}

func mpow32(a, n, m int32) int32 {
	if m == 1 {
		return 0
	}
	var r int32
	r = 1
	for n > 0 {
		if n&1 == 1 {
			r = r * a % m
		}
		a, n = a*a%m, n>>1
	}
	return r
}

func minv(a, m int) int {
	p, x, u := m, 1, 0
	for p != 0 {
		t := a / p
		a, p = p, a-t*p
		x, u = u, x-t*u
	}
	x %= m
	if x < 0 {
		x += m
	}
	return x
}

// m only allow prime number
func minvfermat(a, m int) int {
	return mpow(a, m-2, mod)
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

func isl(l int, def int) []int {
	sl := make([]int, l)
	for i := 0; i < l; i++ {
		sl[i] = def
	}
	return sl
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
func reversei32(sl []int32) {
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

func addIdx(pos, v int, sl []int) []int {
	sl = append(sl[:pos+1], sl[pos:]...)
	sl[pos] = v
	return sl
}

func delIdx(pos int, sl []int) []int {
	return append(sl[:pos], sl[pos+1:]...)
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

// ==================================================
// queue
// ==================================================

/*
	q := list.New()
	q.PushBack(val)
	e := q.Front()
	for e != nil {
		t := e.Value.(int)

		// Do something

		e = e.Next()
    }
*/

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

/*
	cusum2d := newCusum2d(n, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			cusum2d.set(i, j, 1)
		}
	}
	for i := 0; i < n-k+1; i++ {
		for j := 0; j < n-k+1; j++ {
			t:=cusum2d.get(i, j, i+k, j+k)
		}
	}
*/

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

// x1 <= x <= x2, y1 <= y <= y2
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
	u.par[x] += u.par[y]
	u.par[y] = x
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

// l <= x < r
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
result1 := s.query(l,r)
result2 := s.findrightest(l,r,x)
result3 := s.findlefttest(l,r,x)
*/
type stree struct {
	n    int
	b    []int
	def  int
	cmp  func(i, j int) int
	culc func(i, j int) int
}

func newstree(n int, minmax streeminmmax, ctype streeculctype) *stree {
	tn := 1
	for tn < n {
		tn *= 2
	}
	s := &stree{
		n: tn,
		b: make([]int, 2*tn-1),
	}
	switch minmax {
	case stmin:
		s.def = inf
		for i := 0; i < 2*tn-1; i++ {
			s.b[i] = s.def
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

func (s *stree) query(a, b int) int {
	return s.querysub(a, b, 0, 0, s.n)
}

func (s *stree) querysub(a, b, k, l, r int) int {
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
}

/*
s := newlazystree(n,stmin|stmax,stset|stadd)
s.set(i,x)
s.add(i,x)
s.rc(l,r,x)
result1 := s.query(l,r)
result2 := s.findrightest(l,r,x)
result3 := s.findlefttest(l,r,x)
*/
type lazystree struct {
	on   int
	n    int
	b    []int
	lazy []int
	def  int
	cmp  func(i, j int) int
	culc func(i, j int) int
}

func newlazystree(n int, minmax streeminmmax, ctype streeculctype) lazystree {
	tn := 1
	for tn < n {
		tn *= 2
	}
	s := lazystree{
		on:   n,
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
			if i == s.def {
				return j
			}
			if i == s.def {
				return i
			}
			return i + j
		}
	case stset:
		s.culc = func(i, j int) int {
			return j
		}
	}
	return s
}

func (s lazystree) eval(k int) {
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

func (s lazystree) add(i, x int) {
	i += s.n - 1
	s.b[i] += x

	for i > 0 {
		i = (i - 1) / 2
		s.b[i] = s.cmp(s.b[i*2+1], s.b[i*2+2])
	}
}

func (s lazystree) set(i, x int) {
	i += s.n - 1
	s.b[i] = x

	for i > 0 {
		i = (i - 1) / 2
		s.b[i] = s.cmp(s.b[i*2+1], s.b[i*2+2])
	}
}

// range culc
func (s lazystree) rc(a, b, x int) {
	s.rcsub(a, b, x, 0, 0, s.n)
}

func (s lazystree) rcsub(a, b, x, k, l, r int) {
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

func (s lazystree) get(a int) int {
	return s.query(a, a+1)
}

func (s lazystree) query(a, b int) int {
	return s.querysub(a, b, 0, 0, s.n)
}

func (s lazystree) querysub(a, b, k, l, r int) int {
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

func (s lazystree) findrightest(a, b, x int) int {
	return s.findrightestsub(a, b, x, 0, 0, s.n)
}

func (s lazystree) findrightestsub(a, b, x, k, l, r int) int {
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

func (s lazystree) findleftest(a, b, x int) int {
	return s.findleftestsub(a, b, x, 0, 0, s.n)
}

func (s lazystree) findleftestsub(a, b, x, k, l, r int) int {
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

func (s lazystree) debug() {
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

func (s lazystree) debug2() {
	l := make([]string, s.n)
	for i := 0; i < s.on; i++ {
		l[i] = strconv.Itoa(s.get(i))
	}
	out(strings.Join(l, " "))
}

// ==================================================
// tree
// ==================================================

type tree struct {
	size       int
	root       int
	edges      [][]edge
	parentsize int
	parent     [][]int
	depth      []int
	orderidx   int
	order      []int
}

/*
	n := ni()
	edges := make([][]edge, n)
	for i := 0; i < n-1; i++ {
		s, t := ni2()
		s--
		t--
		edges[s] = append(edges[s], edge{to: t})
		edges[t] = append(edges[t], edge{to: s})
	}
	tree := newtree(n, 0, edges)
	tree.init()
*/
func newtree(size int, root int, edges [][]edge) *tree {
	parentsize := int(math.Log2(float64(size))) + 1
	parent := make([][]int, parentsize)
	for i := 0; i < parentsize; i++ {
		parent[i] = make([]int, size)
	}
	depth := make([]int, size)
	order := make([]int, size)
	return &tree{
		size:       size,
		root:       root,
		edges:      edges,
		parentsize: parentsize,
		parent:     parent,
		depth:      depth,
		order:      order,
	}
}

func (t *tree) init() {
	t.dfs(t.root, -1, 0)
	for i := 0; i+1 < t.parentsize; i++ {
		for j := 0; j < t.size; j++ {
			if t.parent[i][j] < 0 {
				t.parent[i+1][j] = -1
			} else {
				t.parent[i+1][j] = t.parent[i][t.parent[i][j]]
			}
		}
	}
}

func (t *tree) dfs(v, p, d int) {
	t.order[v] = t.orderidx
	t.orderidx++
	t.parent[0][v] = p
	t.depth[v] = d
	for _, nv := range t.edges[v] {
		if nv.to != p {
			t.dfs(nv.to, v, d+1)
		}
	}
}

func (t *tree) lca(u, v int) int {
	if t.depth[u] > t.depth[v] {
		u, v = v, u
	}
	for i := 0; i < t.parentsize; i++ {
		if (t.depth[v]-t.depth[u])>>i&1 == 1 {
			v = t.parent[i][v]
		}
	}
	if u == v {
		return u
	}
	for i := t.parentsize - 1; i >= 0; i-- {
		if t.parent[i][u] != t.parent[i][v] {
			u = t.parent[i][u]
			v = t.parent[i][v]
		}
	}
	return t.parent[0][u]
}

func (t *tree) dist(u, v int) int {
	return t.depth[u] + t.depth[v] - t.depth[t.lca(u, v)]*2
}

func (t *tree) auxiliarytree(sl []int) []int {
	sort.Slice(sl, func(i, j int) bool { return t.order[sl[i]] < t.order[sl[j]] })
	return sl
}

// ==================================================
// graph
// ==================================================

type edge struct {
	from int
	to   int
	cost int
	rev  int
}

func setDualEdge(edges [][]edge, s, t, c int) {
	edges[s] = append(edges[s], edge{to: t, cost: c, rev: len(edges[t])})
	edges[t] = append(edges[t], edge{to: s, cost: 0, rev: len(edges[s]) - 1})
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
	level        []int
	iter         []int
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

func (g *graph) dinic() int {
	f := 0
	for {
		g.dinicbfs(0)
		if g.level[g.size-1] < 0 {
			break
		}
		g.iter = make([]int, g.size)
		for {
			t := g.dinicdfs(0, g.size-1, inf)
			if t <= 0 {
				break
			}
			f += t
		}
	}
	return f
}

func (g *graph) dinicbfs(s int) {
	g.level = make([]int, g.size)
	for i := 0; i < g.size; i++ {
		g.level[i] = -1
	}
	g.level[s] = 0

	q := list.New()
	q.PushBack(s)
	e := q.Front()
	for e != nil {
		t := e.Value.(int)

		for _, e := range g.edges[t] {
			if e.cost > 0 && g.level[e.to] < 0 {
				g.level[e.to] = g.level[t] + 1
				q.PushBack(e.to)
			}
		}

		e = e.Next()
	}
}

func (g *graph) dinicdfs(v, t, f int) int {
	if v == t {
		return f
	}
	for i, e := range g.edges[v] {
		if e.cost > 0 && g.level[v] < g.level[e.to] {
			d := g.dinicdfs(e.to, t, min(f, e.cost))
			if d > 0 {
				g.edges[v][i].cost -= d
				g.edges[e.to][e.rev].cost += d
				return d
			}
		}
	}
	return 0
}

// ==================================================
// fft
// ==================================================

func convolutionMod(a, b []int32, mod int) []int32 {
	n1, n2 := len(a), len(b)
	n := n1 + n2 - 1
	if n1 == 0 || n2 == 0 {
		return []int32{}
	}

	z := 1 << ceilPow2(n)
	aa, bb := make([]int32, z), make([]int32, z)
	copy(aa, a)
	copy(bb, b)
	a, b = aa, bb

	butterfly(a, mod)
	butterfly(b, mod)
	for i := 0; i < z; i++ {
		a[i] = int32(int(a[i]) * int(b[i]) % mod)
	}
	butterflyInv(a, mod)
	a = a[:n]
	iz := minv(z, mod)
	for i := 0; i < n; i++ {
		a[i] = int32(int(a[i]) * iz % mod)
		if a[i] < 0 {
			a[i] += int32(mod)
		}
	}

	return a
}

func primitiveRoot(m int) int {
	if m == 2 {
		return 1
	}
	if m == 167772161 || m == 469762049 || m == 998244353 {
		return 3
	}
	if m == 754974721 {
		return 11
	}
	divs := make([]int, 20)
	divs[0] = 2
	cnt := 1
	x := (m - 1) / 2
	for x%2 == 0 {
		x /= 2
	}
	for i := 3; i*i <= x; i += 2 {
		if x%i == 0 {
			divs[cnt] = i
			cnt++
			for x%i == 0 {
				x /= i
			}
		}
	}
	if x > 1 {
		divs[cnt] = x
		cnt++
	}
	for g := 2; ; g++ {
		ok := true
		for i := 0; i < cnt; i++ {
			if mpow(g, (m-1)/divs[i], m) == 1 {
				ok = false
				break
			}
		}
		if ok {
			return g
		}
	}
}

func ceilPow2(n int) int {
	return bits.Len(uint(n - 1))
}

func bsf(n int) int {
	return bits.TrailingZeros(uint(n))
}

func butterfly(a []int32, M int) {
	g := primitiveRoot(M)
	se := make([]int32, 30)
	if M == 998244353 {
		se = []int32{
			911660635, 509520358, 369330050, 332049552, 983190778, 123842337, 238493703, 975955924, 603855026, 856644456, 131300601, 842657263, 730768835, 942482514, 806263778, 151565301, 510815449, 503497456, 743006876, 741047443, 56250497, 867605899, 0, 0, 0, 0, 0, 0, 0, 0}
	} else {

		es, ies := make([]int32, 30), make([]int32, 30)
		cnt2 := bsf(M - 1)
		e := int32(mpow(g, (M-1)>>cnt2, M))
		ie := int32(minv(int(e), M))
		for i := cnt2; i >= 2; i-- {
			es[i-2] = e
			ies[i-2] = ie
			e = int32(int(e) * int(e) % M)
			ie = int32(int(ie) * int(ie) % M)
		}
		var now int32
		now = 1
		for i := 0; i <= cnt2-2; i++ {
			se[i] = int32(int(es[i]) * int(now) % M)
			now = int32(int(now) * int(ies[i]) % M)
		}
	}

	mm := int32(M)
	n := len(a)
	h := ceilPow2(n)
	for ph := 1; ph <= h; ph++ {
		w := 1 << (ph - 1)
		p := 1 << (h - ph)
		var now int32
		now = 1
		for s := 0; s < w; s++ {
			offset := s << (h - ph + 1)
			for i := 0; i < p; i++ {
				l := a[i+offset]
				r := int32(int(a[i+offset+p]) * int(now) % M)
				a[i+offset] = (l + r) % mm
				a[i+offset+p] = (mm + l - r) % mm
			}
			now = int32(int(now) * int(se[bsf(^s)]) % M)
		}
	}
}

func butterflyInv(a []int32, M int) {
	g := primitiveRoot(M)
	sie := make([]int32, 30)
	if M == 998244353 {
		sie = []int32{86583718, 372528824, 373294451, 645684063, 112220581, 692852209, 155456985, 797128860, 90816748, 860285882, 927414960, 354738543, 109331171, 293255632, 535113200, 308540755, 121186627, 608385704, 438932459, 359477183, 824071951, 103369235, 0, 0, 0, 0, 0, 0, 0, 0}
	} else {

		es, ies := make([]int32, 30), make([]int32, 30)
		cnt2 := bsf(M - 1)
		e := int32(mpow(g, (M-1)>>cnt2, M))
		ie := int32(minv(int(e), M))
		for i := cnt2; i >= 2; i-- {
			es[i-2] = e
			ies[i-2] = ie
			e = int32(int(e) * int(e) % M)
			ie = int32(int(ie) * int(ie) % M)
		}
		var now int32
		now = 1
		for i := 0; i <= cnt2-2; i++ {
			sie[i] = int32(int(ies[i]) * int(now) % M)
			now = int32(int(now) * int(es[i]) % M)
		}
		out(sie)
	}

	mm := int32(M)
	n := len(a)
	h := ceilPow2(n)
	for ph := h; ph >= 1; ph-- {
		w := 1 << (ph - 1)
		p := 1 << (h - ph)
		var now int32
		now = 1
		for s := 0; s < w; s++ {
			offset := s << (h - ph + 1)
			for i := 0; i < p; i++ {
				l := int32(a[i+offset])
				r := int32(a[i+offset+p])
				a[i+offset] = (l + r) % mm
				a[i+offset+p] = int32(int(mm+l-r) * int(now) % M)
			}
			now = mmul32(now, sie[bsf(^s)])
		}
	}
}

func ConolutionPowMod(a, e, mod int) int {
	res, m := 1, a
	for e > 0 {
		if e&1 != 0 {
			res = res * m % mod
		}
		m = m * m % mod
		e >>= 1
	}
	return res
}

type Conolution struct {
	mod, primroot, rank2                      int
	root, iroot, rate2, irate2, rate3, irate3 []int
}

func NewConvolution(mod, primroot int) *Conolution {
	rank2 := bits.TrailingZeros(uint(mod - 1))
	if rank2 < 3 {
		panic("Hard wired to work for a significantly large power of 2 in the modulus")
	}
	root := make([]int, rank2+1)
	iroot := make([]int, rank2+1)
	rate2 := make([]int, rank2-2+1)
	irate2 := make([]int, rank2-2+1)
	rate3 := make([]int, rank2-3+1)
	irate3 := make([]int, rank2-3+1)
	root[rank2] = ConolutionPowMod(primroot, (mod-1)>>rank2, mod)
	iroot[rank2] = ConolutionPowMod(root[rank2], mod-2, mod)
	for i := rank2 - 1; i >= 0; i-- {
		root[i] = root[i+1] * root[i+1] % mod
		iroot[i] = iroot[i+1] * iroot[i+1] % mod
	}
	prod, iprod := 1, 1
	for i := 0; i <= rank2-2; i++ {
		rate2[i] = root[i+2] * prod % mod
		irate2[i] = iroot[i+2] * iprod % mod
		prod = prod * iroot[i+2] % mod
		iprod = iprod * root[i+2] % mod
	}
	prod, iprod = 1, 1
	for i := 0; i <= rank2-3; i++ {
		rate3[i] = root[i+3] * prod % mod
		irate3[i] = iroot[i+3] * iprod % mod
		prod = prod * iroot[i+3] % mod
		iprod = iprod * root[i+3] % mod
	}
	return &Conolution{mod, primroot, rank2, root, iroot, rate2, irate2, rate3, irate3}
}
func (q *Conolution) butterfly(a []int) {
	mod := q.mod
	n := len(a)
	h := 0
	for (1 << h) < n {
		h++
	}
	ll := 0
	for ll < h {
		if h-ll == 1 {
			p := 1 << (h - ll - 1)
			rot := 1
			for s := 0; s < (1 << ll); s++ {
				offset := s << (h - ll)
				for i := 0; i < p; i++ {
					l := a[i+offset]
					r := a[i+offset+p] * rot % mod
					u := l + r
					if u >= mod {
						u -= mod
					}
					v := l - r
					if v < 0 {
						v += mod
					}
					a[i+offset] = u
					a[i+offset+p] = v
				}
				if s+1 != (1 << ll) {
					rot = rot * q.rate2[bits.TrailingZeros(^uint(s))] % mod
				}
			}
			ll++
		} else {
			p := 1 << (h - ll - 2)
			rot := 1
			imag := q.root[2]
			for s := 0; s < (1 << ll); s++ {
				rot2 := rot * rot % mod
				rot3 := rot2 * rot % mod
				offset := s << (h - ll)
				for i := 0; i < p; i++ {
					mod2 := mod * mod
					a0 := a[i+offset]
					a1 := a[i+offset+p] * rot
					a2 := a[i+offset+2*p] * rot2
					a3 := a[i+offset+3*p] * rot3
					a1na3imag := (a1 + mod2 - a3) % mod * imag
					na2 := mod2 - a2
					a[i+offset] = (a0 + a2 + a1 + a3) % mod
					a[i+offset+p] = (a0 + a2 + (2*mod2 - a1 - a3)) % mod
					a[i+offset+2*p] = (a0 + na2 + a1na3imag) % mod
					a[i+offset+3*p] = (a0 + na2 + (mod2 - a1na3imag)) % mod
				}
				if s+1 != (1 << ll) {
					rot = rot * q.rate3[bits.TrailingZeros(^uint(s))] % mod
				}
			}
			ll += 2
		}
	}
}
func (q *Conolution) butterflyInv(a []int) {
	mod := q.mod
	n := len(a)
	h := 0
	for (1 << h) < n {
		h++
	}
	ll := h
	for ll > 0 {
		if ll == 1 {
			p := 1 << (h - ll)
			irot := 1
			for s := 0; s < (1 << (ll - 1)); s++ {
				offset := s << (h - ll + 1)
				for i := 0; i < p; i++ {
					l := a[i+offset]
					r := a[i+offset+p]
					u := l + r
					if u >= mod {
						u -= mod
					}
					v := (mod + l - r) * irot % mod
					a[i+offset] = u
					a[i+offset+p] = v
				}
				if s+1 != (1 << (ll - 1)) {
					irot = irot * q.irate2[bits.TrailingZeros(^uint(s))] % mod
				}
			}
			ll--
		} else {
			p := 1 << (h - ll)
			irot := 1
			iimag := q.iroot[2]
			for s := 0; s < (1 << (ll - 2)); s++ {
				irot2 := irot * irot % mod
				irot3 := irot2 * irot % mod
				offset := s << (h - ll + 2)
				for i := 0; i < p; i++ {
					a0 := a[i+offset]
					a1 := a[i+offset+p]
					a2 := a[i+offset+2*p]
					a3 := a[i+offset+3*p]
					a2na3iimag := (mod + a2 - a3) * iimag % mod
					a[i+offset] = (a0 + a1 + a2 + a3) % mod
					a[i+offset+p] = (a0 + (mod - a1) + a2na3iimag) * irot % mod
					a[i+offset+2*p] = (a0 + a1 + (mod - a2) + (mod - a3)) * irot2 % mod
					a[i+offset+3*p] = (a0 + (mod - a1) + (mod - a2na3iimag)) * irot3 % mod
				}
				if s+1 != (1 << (ll - 2)) {
					irot = irot * q.irate3[bits.TrailingZeros(^uint(s))] % mod
				}
			}
			ll -= 2
		}
	}
	iz := ConolutionPowMod(n, mod-2, mod)
	for i := 0; i < n; i++ {
		a[i] = a[i] * iz % mod
	}
}
func (q *Conolution) convolveFFT(a []int, b []int) []int {
	mod := q.mod
	finalsz := len(a) + len(b) - 1
	z := 1
	for z < finalsz {
		z *= 2
	}
	lena, lenb := len(a), len(b)
	la := make([]int, z)
	lb := make([]int, z)
	for i := 0; i < lena; i++ {
		la[i] = a[i]
	}
	for i := 0; i < lenb; i++ {
		lb[i] = b[i]
	}
	q.butterfly(la)
	q.butterfly(lb)
	for i := 0; i < z; i++ {
		la[i] *= lb[i]
		la[i] %= mod
	}
	q.butterflyInv(la)
	return la[:finalsz]
}
func (q *Conolution) ConolutionNaive(a []int, b []int) []int {
	mod := q.mod
	finalsz := len(a) + len(b) - 1
	ans := make([]int, finalsz)
	for i, a := range a {
		for j, b := range b {
			ans[i+j] += a * b
			ans[i+j] %= mod
		}
	}
	return ans
}
func (q *Conolution) Convolve(a []int, b []int) []int {
	lmin := len(a)
	if len(b) < lmin {
		lmin = len(b)
	}
	if lmin <= 60 {
		return q.ConolutionNaive(a, b)
	} else {
		return q.convolveFFT(a, b)
	}
}
