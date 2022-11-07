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

	defer flush()

	o := 0
	n := ni()
	ns := nis(n)
	out(o)
}

// ==================================================
// init
// ==================================================

const inf = math.MaxInt64
const mod1000000007 = 1000000007
const mod998244353 = 998244353
const mod = mod1000000007

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

func nis(arg ...int) []int {
	n := arg[0]
	t := 0
	if len(arg) == 2 {
		t = arg[1]
	}

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = ni() - t
	}
	return a
}

func ni2s(n int) ([]int, []int) {
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], b[i] = ni2()
	}
	return a, b
}

func ni3s(n int) ([]int, []int, []int) {
	a := make([]int, n)
	b := make([]int, n)
	c := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], b[i], c[i] = ni3()
	}
	return a, b, c
}

func ni4s(n int) ([]int, []int, []int, []int) {
	a := make([]int, n)
	b := make([]int, n)
	c := make([]int, n)
	d := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], b[i], c[i], d[i] = ni4()
	}
	return a, b, c, d
}

func ni2a(n int) [][2]int {
	a := make([][2]int, n)
	for i := 0; i < n; i++ {
		a[i][0], a[i][1] = ni2()
	}
	return a
}

func ni3a(n int) [][3]int {
	a := make([][3]int, n)
	for i := 0; i < n; i++ {
		a[i][0], a[i][1], a[i][2] = ni3()
	}
	return a
}

func ni4a(n int) [][4]int {
	a := make([][4]int, n)
	for i := 0; i < n; i++ {
		a[i][0], a[i][1], a[i][2], a[i][3] = ni4()
	}
	return a
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

func nsis() []int {
	sc.Scan()
	s := sc.Text()
	return stois(s, '0')
}

func scani() int {
	var i int
	fmt.Scanf("%i", &i)
	return i
}

func scans() string {
	var s string
	fmt.Scanf("%s", &s)
	return s
}

func out(v ...interface{}) {
	_, e := fmt.Fprintln(wtr, v...)
	if e != nil {
		panic(e)
	}
}

func outf(f string, v ...interface{}) {
	out(fmt.Sprintf(f, v...))
}

func outwoln(v ...interface{}) {
	_, e := fmt.Fprint(wtr, v...)
	if e != nil {
		panic(e)
	}
}

func outis(sl []int) {
	r := make([]string, len(sl))
	for i, v := range sl {
		r[i] = itoa(v)
	}
	out(strings.Join(r, " "))
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

func maxs(a *int, b int) {
	if *a < b {
		*a = b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func mins(a *int, b int) {
	if *a > b {
		*a = b
	}
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

func pow10(a int) int {
	return int(math.Pow(10, float64(a)))
}

func sqrt(i int) int {
	return int(math.Sqrt(float64(i)))
}

func sqrtf(i int) float64 {
	return math.Sqrt(float64(i))
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

type combFactorial struct {
	fac    []int
	facinv []int
}

func newcombFactorial(n int) *combFactorial {

	fac := make([]int, n)
	facinv := make([]int, n)
	fac[0] = 1
	facinv[0] = minvfermat(1, mod)

	for i := 1; i < n; i++ {
		fac[i] = mmul(i, fac[i-1])
		facinv[i] = minvfermat(fac[i], mod)
	}

	return &combFactorial{
		fac:    fac,
		facinv: facinv,
	}
}

func (c *combFactorial) factorial(n int) int {
	return c.fac[n]
}

func (c *combFactorial) combination(n, r int) int {
	if r > n {
		return 0
	}
	return mmul(mmul(c.fac[n], c.facinv[r]), c.facinv[n-r])
}

func (c *combFactorial) permutation(n, r int) int {
	if r > n {
		return 0
	}
	return mmul(c.fac[n], c.facinv[n-r])
}

func (c *combFactorial) homogeousProduct(n, r int) int {
	return c.combination(n-1+r, r)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func divisor(n int) ([]int, map[int]int) {
	sqrtn := int(math.Sqrt(float64(n)))
	c := 2
	divisor := []int{}
	divisorm := make(map[int]int)
	for {
		if n%2 != 0 {
			break
		}
		divisor = append(divisor, 2)
		divisorm[2]++
		n /= 2
	}
	c = 3
	for {
		if n%c == 0 {
			divisor = append(divisor, c)
			divisorm[c]++
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
		divisorm[n]++
	}
	return divisor, divisorm
}

type binom struct {
	fac  []int
	finv []int
	inv  []int
}

func newbinom(n int) *binom {
	b := &binom{
		fac:  make([]int, n),
		finv: make([]int, n),
		inv:  make([]int, n),
	}
	b.fac[0] = 1
	b.fac[1] = 1
	b.inv[1] = 1
	b.finv[0] = 1
	b.finv[1] = 1
	for i := 2; i < n; i++ {
		b.fac[i] = b.fac[i-1] * i % mod
		b.inv[i] = mod - mod/i*b.inv[mod%i]%mod
		b.finv[i] = b.finv[i-1] * b.inv[i] % mod
	}
	return b
}

func (b *binom) get(n, r int) int {
	if n < r || n < 0 || r < 0 {
		return 0
	}
	return b.fac[n] * b.finv[r] % mod * b.finv[n-r] % mod
}

func matPow(a [][]int, n int) [][]int {
	r := make([][]int, len(a))
	for i := 0; i < len(a); i++ {
		r[i] = is(len(a), 0)
		r[i][i] = 1
	}
	for n > 0 {
		if n&1 != 0 {
			r = matMul(a, r)
		}
		a = matMul(a, a)
		n = n >> 1
	}
	return r
}

func matMul(a, b [][]int) [][]int {
	r := make([][]int, len(a))
	for i := 0; i < len(a); i++ {
		r[i] = is(len(b[0]), 0)
	}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b[0]); j++ {
			for k := 0; k < len(b); k++ {
				r[i][j] = madd(r[i][j], mmul(a[i][k], b[k][j]))
			}
		}
	}
	return r
}

// ==================================================
// mod
// ==================================================

func madd(a, b int) int {
	a += b
	if a >= mod || a <= -mod {
		a %= mod
	}
	if a < 0 {
		a += mod
	}
	return a
}

func mmul(a, b int) int {
	return a * b % mod
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
	o = bs(0, len(sl)-1, func(c int) bool {
		return true
	})
*/
func bs(ok, ng int, f func(int) bool) int {
	if !f(ok) {
		return -1
	}
	if f(ng) {
		return ng
	}
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

/*
	o = bsfl(0.0, 100.0, 100, func(c float64) bool {
		return true
	})
*/
func bsfl(ok, ng float64, c int, f func(float64) bool) float64 {
	for i := 0; i < c; i++ {

		mid := (ok + ng) / 2

		if f(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}

	return ok
}

func bs3fl(low, high float64, c int, f func(float64) float64) float64 {

	for i := 0; i < c; i++ {
		c1 := (low*2 + high) / 3
		c2 := (low + high*2) / 3

		if f(c1) > f(c2) {
			low = c1
		} else {
			high = c2
		}
	}
	return low
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

func bitlen(a int) int {
	return bits.Len(uint(a))
}

func xor(a, b bool) bool { return a != b }

func debugbit(n int) string {
	r := ""
	for i := bitlen(n) - 1; i >= 0; i-- {
		if n&(1<<i) != 0 {
			r += "1"
		} else {
			r += "0"
		}
	}
	return r
}

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

type sortOrder int

const (
	asc sortOrder = iota
	desc
)

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

type Sort2ArOptions struct {
	keys   []int
	orders []sortOrder
}

type Sort2ArOption func(*Sort2ArOptions)

func opt2ar(key int, order sortOrder) Sort2ArOption {
	return func(args *Sort2ArOptions) {
		args.keys = append(args.keys, key)
		args.orders = append(args.orders, order)
	}
}

// sort2ar(sl,opt2ar(1,asc))
// sort2ar(sl,opt2ar(0,asc),opt2ar(1,asc))
func sort2ar(sl [][2]int, setters ...Sort2ArOption) {
	args := &Sort2ArOptions{}

	for _, setter := range setters {
		setter(args)
	}

	sort.Slice(sl, func(i, j int) bool {
		for idx, key := range args.keys {
			if sl[i][key] == sl[j][key] {
				continue
			}
			switch args.orders[idx] {
			case asc:
				return sl[i][key] < sl[j][key]
			case desc:
				return sl[i][key] > sl[j][key]
			}
		}
		return true
	})
}

// ==================================================
// slice
// ==================================================

func is(l int, def int) []int {
	sl := make([]int, l)
	for i := 0; i < l; i++ {
		sl[i] = def
	}
	return sl
}

func i2s(l, m int, def int) [][]int {
	sl := make([][]int, l)
	for i := 0; i < l; i++ {
		sl[i] = make([]int, m)
		for j := 0; j < m; j++ {
			sl[i][j] = def
		}
	}
	return sl
}

//	out(stois("abcde", 'a'))
//	out(stois("abcde", 'a'-1))
//	out(stois("12345", '0'))
func stois(s string, baseRune rune) []int {
	r := make([]int, len(s))
	for i, v := range s {
		r[i] = int(v - baseRune)
	}
	return r
}

func istos(s []int, baseRune rune) string {
	r := ""
	for _, v := range s {
		r += string(v + int(baseRune))
	}
	return r
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
	hist := make(map[int]struct{})
	j := 0
	rsl := make([]int, len(sl))
	for i := 0; i < len(sl); i++ {
		if _, ok := hist[sl[i]]; ok {
			continue
		}
		rsl[j] = sl[i]
		hist[sl[i]] = struct{}{}
		j++
	}
	return rsl[:j]
}

// coordinate compression
func cocom(sl []int) ([]int, map[int]int) {
	rsl := uniquei(sl)
	sorti(rsl)
	rm := make(map[int]int)
	for i := 0; i < len(rsl); i++ {
		rm[rsl[i]] = i
	}
	return rsl, rm
}

func popBack(sl []int) (int, []int) {
	return sl[len(sl)-1], sl[:len(sl)-1]
}

func addIdx(pos, v int, sl []int) []int {
	if len(sl) == pos {
		sl = append(sl, v)
		return sl
	}
	sl = append(sl[:pos+1], sl[pos:]...)
	sl[pos] = v
	return sl
}

func delIdx(pos int, sl []int) []int {
	return append(sl[:pos], sl[pos+1:]...)
}

// find x of sl[x] < v. return -1 if no lowerbound found
func lowerBound(v int, sl []int) int {
	if len(sl) == 0 {
		return -1
	}
	idx := bs(0, len(sl)-1, func(c int) bool {
		return sl[c] < v
	})
	return idx
}

// find x of v < sl[x]. return len(sl) if no upperbound found
func upperBound(v int, sl []int) int {
	if len(sl) == 0 {
		return 0
	}
	idx := bs(0, len(sl)-1, func(c int) bool {
		return sl[c] <= v
	})
	return idx + 1
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
	return point{x: a.x + b.x, y: a.y + b.y}
}

func pointDist(a, b point) float64 {
	return sqrtf((a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y))
}

func pointDistDouble(a, b point) int {
	return (a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y)
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
	ih := newIntHeap(asc)
	ih.Push(v)
	for !ih.IsEmpty() {
		v := ih.Pop(h)
	}
*/
type IntHeap struct {
	sum int
	pq  *pq
}

func newIntHeap(order sortOrder) *IntHeap {
	ih := &IntHeap{}
	ih.pq = newpq([]compFunc{func(p, q interface{}) int {
		if p.(int) == q.(int) {
			return 0
		}
		if order == asc {
			if p.(int) < q.(int) {
				return -1
			} else {
				return 1
			}
		} else {
			if p.(int) > q.(int) {
				return -1
			} else {
				return 1
			}
		}
	}})
	heap.Init(ih.pq)
	return ih
}
func (ih *IntHeap) Push(x int) {
	ih.sum += x
	heap.Push(ih.pq, x)
}

func (ih *IntHeap) Pop() int {
	v := heap.Pop(ih.pq).(int)
	ih.sum -= v
	return v
}

func (ih *IntHeap) Len() int {
	return ih.pq.Len()
}

func (ih *IntHeap) IsEmpty() bool {
	return ih.pq.Len() == 0
}

func (ih *IntHeap) GetRoot() int {
	return ih.pq.GetRoot().(int)
}

func (ih *IntHeap) GetSum() int {
	return ih.sum
}

/*
	h := &OrgIntHeap{}
	heap.Init(h)

	heap.Push(h, v)
	for !h.IsEmpty() {
		v = heap.Pop(h).(int)
	}
*/
type OrgIntHeap []int

func (h OrgIntHeap) Len() int { return len(h) }

// get from bigger
// func (h OrgIntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h OrgIntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h OrgIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *OrgIntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *OrgIntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *OrgIntHeap) IsEmpty() bool {
	return h.Len() == 0
}

// h.Min().(int)
func (h *OrgIntHeap) Min() interface{} {
	return (*h)[0]
}

/*
	type pqst struct {
		x int
		y int
	}

	pq := newpq([]compFunc{func(p, q interface{}) int {
		if p.(pqst).x != q.(pqst).x {
			// get from bigger
			// if p.(pqst).x > q.(pqst).x {
			if p.(pqst).x < q.(pqst).x {
				return -1
			} else {
				return 1
			}
		}
		if p.(pqst).y != q.(pqst).y {
			// get from bigger
			// if p.(pqst).y > q.(pqst).y {
			if p.(pqst).y < q.(pqst).y {
				return -1
			} else {
				return 1
			}
		}
		return 0
	}})
	heap.Init(pq)
	heap.Push(pq, pqst{x: 1, y: 1})
	for !pq.IsEmpty() {
		v := heap.Pop(pq).(pqst)
	}
*/

type pq struct {
	arr   []interface{}
	comps []compFunc
}

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

func (pq *pq) IsEmpty() bool {
	return pq.Len() == 0
}

// pq.GetRoot().(edge)
func (pq *pq) GetRoot() interface{} {
	return pq.arr[0]
}

// ==================================================
// cusum
// ==================================================

type cusum struct {
	l int
	s []int
}

func newcusum(sl []int) *cusum {
	c := &cusum{}
	c.l = len(sl)
	c.s = make([]int, len(sl)+1)
	for i, v := range sl {
		c.s[i+1] = c.s[i] + v
	}
	return c
}

// get sum f <= i && i <= t
func (c *cusum) getRange(f, t int) int {
	if f > t || f >= c.l {
		return 0
	}
	return c.s[t+1] - c.s[f]
}

// get sum 0 to i
func (c *cusum) get(i int) int {
	return c.s[i+1]
}

func (c *cusum) upperBound(i int) int {
	return upperBound(i, c.s)
}

func (c *cusum) lowerBound(i int) int {
	return lowerBound(i, c.s)
}

/*
	mp := make([][]int, n)
	for i := 0; i < k; i++ {
		mp[i] = make([]int, m)
	}
	cusum2d := newcusum2d(sl)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			t:=cusum2d.get(0, 0, i, j)
		}
	}
*/

type cusum2d struct {
	s [][]int
}

func newcusum2d(sl [][]int) *cusum2d {
	c := &cusum2d{}
	n := len(sl)
	m := len(sl[0])
	c.s = make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		c.s[i] = make([]int, m+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			c.s[i+1][j+1] = c.s[i+1][j] + c.s[i][j+1] - c.s[i][j]
			c.s[i+1][j+1] += sl[i][j]
		}
	}
	return c
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

func (b *bit) culc(i, j int) int {
	return i + j
	//return madd(i, j)
}

func (b *bit) add(i, x int) {
	for i++; i < b.n && i > 0; i += i & -i {
		b.b[i] = b.culc(b.b[i], x)
	}
}

func (b *bit) sum(i int) int {
	ret := 0
	for i++; i > 0; i -= i & -i {
		ret = b.culc(ret, b.b[i])
	}
	return ret
}

// l <= x < r
func (b *bit) rangesum(l, r int) int {
	return b.culc(b.sum(r-1), -b.sum(l-1))
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

const (
	stadd streeculctype = iota
	stmadd
	stset
	stmin
	stmax
)

/*
s := newstree(n,stmin|stmax|stsum|stmsum)
s.set(i,x)
s.add(i,x)
result1 := s.query(l,r)
result2 := s.findrightest(l,r,x)
result3 := s.findlefttest(l,r,x)
*/
type stree struct {
	n   int
	b   []int
	def int
	cmp func(i, j int) int
}

func newstree(n int, minmax streeculctype) *stree {
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
	case stadd:
		s.cmp = func(i, j int) int {
			if i == s.def {
				return j
			}
			if j == s.def {
				return i
			}
			return i + j
		}
	case stmadd:
		s.cmp = func(i, j int) int {
			if i == s.def {
				return j
			}
			if j == s.def {
				return i
			}
			return madd(i, j)
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
s := newlazystree(n,stmin|stmax,stset|stadd|stmadd)
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

func newlazystree(n int, cmptype streeculctype, culctype streeculctype) lazystree {
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
	switch cmptype {
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
	switch culctype {
	case stadd:
		s.culc = func(i, j int) int {
			if i == s.def {
				return j
			}
			if j == s.def {
				return i
			}
			return i + j
		}
	case stmadd:
		s.culc = func(i, j int) int {
			if i == s.def {
				return j
			}
			if j == s.def {
				return i
			}
			return madd(i, j)
		}
	case stmax:
		s.culc = func(i, j int) int {
			return max(i, j)
		}
	case stmin:
		s.culc = func(i, j int) int {
			return min(i, j)
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

// range culc a <= n n < b
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
	l := make([]string, s.n+1)
	for i := 0; i <= s.on; i++ {
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
	deg := make([]int, v)
	for i := 0; i < e; i++ {
		s, t, c := ni3()
		s--
		t--
		edges[s] = append(edges[s], edge{to: t, cost: c})
		deg[t]++
	}
	graph := newgraph(v, edges)
	isdag, r := graph.topologicalSort(deg)
*/
func (g *graph) topologicalSort(deg []int) (bool, []int) {

	r := []int{}
	q := list.New()
	for i := 0; i < g.size; i++ {
		if deg[i] == 0 {
			q.PushBack(i)
		}
	}
	e := q.Front()
	for e != nil {
		t := e.Value.(int)
		r = append(r, t)
		for _, edge := range g.edges[t] {
			deg[edge.to]--
			if deg[edge.to] == 0 {
				q.PushBack(edge.to)
			}
		}

		e = e.Next()
	}
	for _, v := range deg {
		if v != 0 {
			return false, nil
		}
	}
	return true, r
}

/*
	v, e := ni2()
	edges := make([][]edge, v)
	edgers := make([][]edge, v)

	for i := 0; i < e; i++ {
		s, t := ni2()
		s--
		t--
		edges[s] = append(edges[s], edge{to: t})
		edgers[t] = append(edgers[t], edge{to: s})
	}

	scc := getScc(v, edges, edgers)
*/
func getScc(v int, edges, edgers [][]edge) [][]int {
	used := make([]bool, v)
	scc := [][]int{}
	vs := []int{}

	var dfs func(i int)
	dfs = func(i int) {
		used[i] = true
		for _, v := range edges[i] {
			if used[v.to] == false {
				dfs(v.to)
			}
		}
		vs = append(vs, i)
	}

	var rdfs func(i, k int)
	rdfs = func(i, k int) {
		used[i] = true
		scc[k] = append(scc[k], i)
		for _, v := range edgers[i] {
			if used[v.to] == false {
				rdfs(v.to, k)
			}
		}
	}

	for i := 0; i < v; i++ {
		if used[i] == false {
			dfs(i)
		}
	}
	used = make([]bool, v)
	k := 0
	for i := v - 1; i >= 0; i-- {
		if used[vs[i]] == false {
			scc = append(scc, []int{})
			rdfs(vs[i], k)
			k++
		}
	}
	return scc
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

/*
	v, e := ni2()
	edges := make([][]edge, v)
	for i := 0; i < e; i++ {
		s, t, c := ni3()
		s--
		t--
		setDualEdge(edges, s, t, c)
	}
	graph := newgraph(v, edges)
	o = graph.dinic()
*/
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

func convolution(a, b []int) []int {
	n1, n2 := len(a), len(b)
	n := n1 + n2 - 1
	if n1 == 0 || n2 == 0 {
		return []int{}
	}

	MOD1 := 754974721
	MOD2 := 167772161
	MOD3 := 469762049
	M2M3 := MOD2 * MOD3
	M1M3 := MOD1 * MOD3
	M1M2 := MOD1 * MOD2
	M1M2M3 := MOD1 * MOD2 * MOD3

	i1 := minv(M2M3, MOD1)
	i2 := minv(M1M3, MOD2)
	i3 := minv(M1M2, MOD3)

	c1 := convolutionMod(a, b, MOD1)
	c2 := convolutionMod(a, b, MOD2)
	c3 := convolutionMod(a, b, MOD3)

	c := make([]int, n)
	offset := []int{0, 0, M1M2M3, 2 * M1M2M3, 3 * M1M2M3}

	for i := 0; i < n; i++ {
		x := 0
		x += c1[i] * i1 % MOD1 * M2M3
		x += c2[i] * i2 % MOD2 * M1M3
		x += c3[i] * i3 % MOD3 * M1M2
		diff := c1[i] - x%MOD1
		if diff < 0 {
			diff += MOD1
		}
		x -= offset[diff%5]
		c[i] = x
	}

	return c
}

func convolutionMod(a, b []int, mod int) []int {
	n1, n2 := len(a), len(b)
	n := n1 + n2 - 1
	if n1 == 0 || n2 == 0 {
		return []int{}
	}

	z := 1 << ceilPow2(n)
	aa, bb := make([]int, z), make([]int, z)
	copy(aa, a)
	copy(bb, b)
	a, b = aa, bb

	butterfly(a, mod)
	butterfly(b, mod)
	for i := 0; i < z; i++ {
		a[i] = a[i] * b[i] % mod
	}
	butterflyInv(a, mod)
	a = a[:n]
	iz := minv(z, mod)
	for i := 0; i < n; i++ {
		a[i] = a[i] * iz % mod
		if a[i] < 0 {
			a[i] += mod
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
	x := 0
	for 1<<x < n {
		x++
	}
	return x
}

func bsf(n int) int {
	x := 0
	for n&(1<<x) == 0 {
		x++
	}
	return x
}

func butterfly(a []int, M int) {
	g := primitiveRoot(M)
	n := len(a)
	h := ceilPow2(n)

	se := make([]int, 30)
	es, ies := make([]int, 30), make([]int, 30)
	cnt2 := bsf(M - 1)
	e := mpow(g, (M-1)>>cnt2, M)
	ie := minv(e, M)
	for i := cnt2; i >= 2; i-- {
		es[i-2] = e
		ies[i-2] = ie
		e = e * e % M
		ie = ie * ie % M
	}
	now := 1
	for i := 0; i <= cnt2-2; i++ {
		se[i] = es[i] * now % M
		now = now * ies[i] % M
	}
	for ph := 1; ph <= h; ph++ {
		w := 1 << (ph - 1)
		p := 1 << (h - ph)
		now := 1
		for s := 0; s < w; s++ {
			offset := s << (h - ph + 1)
			for i := 0; i < p; i++ {
				l := a[i+offset]
				r := a[i+offset+p] * now % M
				a[i+offset] = (l + r) % M
				a[i+offset+p] = (M + l - r) % M
			}
			now = now * se[bsf(^s)] % M
		}
	}
}

func butterflyInv(a []int, M int) {
	g := primitiveRoot(M)
	n := len(a)
	h := ceilPow2(n)

	sie := make([]int, 30)
	es, ies := make([]int, 30), make([]int, 30)
	cnt2 := bsf(M - 1)
	e := mpow(g, (M-1)>>cnt2, M)
	ie := minv(e, M)
	for i := cnt2; i >= 2; i-- {
		es[i-2] = e
		ies[i-2] = ie
		e = e * e % M
		ie = ie * ie % M
	}
	now := 1
	for i := 0; i <= cnt2-2; i++ {
		sie[i] = ies[i] * now % M
		now = now * es[i] % M
	}
	for ph := h; ph >= 1; ph-- {
		w := 1 << (ph - 1)
		p := 1 << (h - ph)
		inow := 1
		for s := 0; s < w; s++ {
			offset := s << (h - ph + 1)
			for i := 0; i < p; i++ {
				l := a[i+offset]
				r := a[i+offset+p]
				a[i+offset] = (l + r) % M
				a[i+offset+p] = (M + l - r) * inow % M
			}
			inow = inow * sie[bsf(^s)] % M
		}
	}
}
