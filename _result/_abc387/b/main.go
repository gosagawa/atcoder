package main

import (
	"bufio"
	"container/heap"
	"container/list"
	"fmt"
	"io/ioutil"
	"math"
	"math/bits"
	"math/rand"
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
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if i*j != n {
				o += i * j
			}
		}
	}
	out(o)
}

// ==================================================
// init
// ==================================================

const inf = math.MaxInt64
const mod1000000007 = 1000000007
const mod998244353 = 998244353
const mod = mod998244353
const baseRune = 'a'
const maxlogn = 62

var mpowcache map[[3]int]int
var debugFlg bool

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
	if len(os.Args) > 1 && os.Args[1] == "i" {
		b, e := ioutil.ReadFile("./input")
		if e != nil {
			panic(e)
		}
		sc = bufio.NewScanner(strings.NewReader(strings.Replace(string(b), " ", "\n", -1)))
		debugFlg = true
	}
	mpowcache = make(map[[3]int]int)
}

// ==================================================
// io
// ==================================================

func ni() int {
	sc.Scan()
	return atoi(sc.Text())
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

func ni2d(n, m int) [][]int {
	a := i2s(n, m, 0)
	for i := 0; i < n; i++ {
		a[i] = nis(m)
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
	return stois(s, baseRune)
}

func nsi2s(n int) [][]int {
	mp := make([][]int, n)
	for i := 0; i < n; i++ {
		mp[i] = nsis()
	}
	return mp
}

// mp := convidxi2s(nsi2s(n), map[string]int{".": 0, "#": 1})
func convidxi2s(sl [][]int, conv map[string]int) [][]int {
	imap := make(map[int]int)
	for s, v := range conv {
		imap[ctoi(s)] = v
	}
	for i, sl2 := range sl {
		for j, v := range sl2 {
			sl[i][j] = imap[v]
		}
	}
	return sl
}

// mp := convidxis(nsis(), map[string]int{".": 0, "#": 1})
func convidxis(sl []int, conv map[string]int) []int {
	imap := make(map[int]int)
	for s, v := range conv {
		imap[ctoi(s)] = v
	}
	for i, v := range sl {
		sl[i] = imap[v]
	}
	return sl
}

func ctoi(c string) int {
	return int(rune(c[0]) - baseRune)
}

func nsiis() []int {
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

func dbg(v ...interface{}) {
	if !debugFlg {
		return
	}
	out(v...)
}

func dbgi2s(sl [][]int) {
	if !debugFlg {
		return
	}
	for _, v := range sl {
		outis(v)
	}
	out("")
}

func dbglst(n int, lst *lazysegtree) {
	if !debugFlg {
		return
	}
	sl := make([]segstruct, n)
	for i := 0; i < n; i++ {
		sl[i] = lst.get(i)
	}
	out(sl)
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

func outisnr(sl []int) {
	for _, v := range sl {
		out(v)
	}
}

func out2d(i, j int) {
	outf("%v %v", i, j)
}

func outsj(sl []string) {
	out(sj(sl))
}

func outsjsp(sl []string) {
	out(sjsp(sl))
}

func outfl(v float64) {
	outf("%.15f", v)
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

func bytoi(b byte) int {
	return atoi(string(b))
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
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

func maxf(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func maxsf(a *float64, b float64) {
	if *a < b {
		*a = b
	}
}

func minf(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func minsf(a *float64, b float64) {
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

var pow2cache [64]int

func pow2(i int) int {
	if pow2cache[i] == 0 {
		pow2cache[i] = int(math.Pow(2, float64(i)))
	}
	return pow2cache[i]
}

var pow10cache [20]int

func pow10(i int) int {
	if pow10cache[i] == 0 {
		pow10cache[i] = int(math.Pow(10, float64(i)))
	}
	return pow10cache[i]
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

func addAngle(a1, a2 float64) float64 {
	r := a1 + a2
	if r >= 360 {
		r -= 360
	}
	return r
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

func gcm(a, b int) int {
	t := gcd(a, b)
	return a * b / t
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

func alldivisor(n int) []int {
	sqrtn := int(math.Sqrt(float64(n)))
	divisor := []int{}
	for i := 1; i <= sqrtn; i++ {
		if n%i != 0 {
			continue
		}
		divisor = append(divisor, i)
		if n/i != i {
			divisor = append(divisor, n/i)
		}
	}
	return divisor
}

func mmod(a, m int) int {
	return (a%m + m) % m
}

func extGcd(a, b int) (int, int, int) {
	if b == 0 {
		return 1, 0, a
	}
	q, p, d := extGcdSub(b, a%b, 0, 0)
	q -= a / b * p
	return p, q, d
}

func extGcdSub(a, b, p, q int) (int, int, int) {
	if b == 0 {
		return 1, 0, a
	}
	q, p, d := extGcdSub(b, a%b, q, p)
	q -= a / b * p
	return p, q, d
}

func chineseRem(b1, m1, b2, m2 int) (bool, int, int) {
	p, _, d := extGcd(m1, m2)
	if (b2-b1)%d != 0 {
		return false, 0, 0
	}
	m := m1 * (m2 / d)
	tmp := (b2 - b1) / d * p % (m2 / d)
	r := mmod(b1+m1*tmp, m)
	return true, r, m
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

func randomint(i int) int {
	return rand.Int() % i
}

// ==================================================
// mod
// ==================================================

func madd(a, b int) int {
	a %= mod
	b %= mod
	a += b
	if a >= mod || a <= -mod {
		a %= mod
	}
	if a < 0 {
		a += mod
	}
	return a
}

func madds(a *int, b int) {
	*a = madd(*a, b)
}

func mmul(a, b int) int {
	a %= mod
	b %= mod
	return a * b % mod
}

func mmuls(a *int, b int) {
	*a = mmul(*a, b)
}

func mdiv(a, b int) int {
	a %= mod
	if b <= 0 || b >= mod {
		panic("invalid division")
	}
	return a * minvfermat(b, mod) % mod
}

func mdivs(a *int, b int) {
	*a = mdiv(*a, b)
}

func mpow(a, n, m int) int {
	if v, ok := mpowcache[[3]int{a, n, m}]; ok {
		return v
	}
	fa := a
	fn := n
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
	mpowcache[[3]int{fa, fn, m}] = r
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

func bs3i(low, high int, f func(int) int) (int, int) {

	for high-low > 2 {
		c1 := (low*2 + high) / 3
		c2 := (low + high*2) / 3
		fc1 := f(c1)
		fc2 := f(c2)

		if fc1 > fc2 {
			low = c1
		} else {
			high = c2
		}
	}
	if high-low == 2 {
		fc1 := f(low)
		fc2 := f(high)
		//	out(high, low, c1, c2, fc1, fc2)

		if fc1 > fc2 {
			low++
		} else {
			high--
		}
	}
	ri := 0
	rv := 0
	if high-low == 1 {
		fc1 := f(low)
		fc2 := f(high)
		//	out(high, low, c1, c2, fc1, fc2)

		if fc1 > fc2 {
			ri = high
			rv = fc2
		} else {
			ri = low
			rv = fc1
		}
	}
	return ri, rv
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

func ss(l int) []string {
	return make([]string, l)
}

func sj(sl []string) string {
	return strings.Join(sl, "")
}

func sjsp(sl []string) string {
	return strings.Join(sl, " ")
}

// out(stois("abcde", 'a'))
// out(stois("abcde", 'a'-1))
// out(stois("12345", '0'))
func stois(s string, baseRune rune) []int {
	r := make([]int, len(s))
	for i, v := range s {
		r[i] = int(v - baseRune)
	}
	return r
}

func istos(s []int, baseRune rune) string {
	r := make([]byte, len(s))
	for i, v := range s {
		r[i] = byte(v) + byte(baseRune)
	}
	return string(r)
}

func issum(sl []int) int {
	r := 0
	for _, v := range sl {
		r += v
	}
	return r
}

func issummod(sl []int) int {
	r := 0
	for _, v := range sl {
		r += v
		r %= mod
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

func rotate(sl [][]int) [][]int {
	n := len(sl)
	r := i2s(n, n, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			r[i][j] = sl[n-1-j][i]
		}
	}
	return r
}

// ==================================================
// matrix
// ==================================================

func matrixmul(a, b [][]int) [][]int {
	ac := len(a)
	ar := len(a[0])
	bc := len(b)
	br := len(b[0])
	if ar != bc {
		panic(fmt.Sprintf("invalid matrix mul ar:%v bc:%v", ar, bc))
	}

	r := i2s(ac, br, 0)
	for i := 0; i < ac; i++ {
		for j := 0; j < br; j++ {
			for k := 0; k < ar; k++ {
				r[i][j] += mmul(a[i][k], b[k][j])
				r[i][j] %= mod
			}
		}
	}
	return r
}

func slmatrixmul(a []int, b [][]int) []int {
	ar := len(a)
	bc := len(b)
	br := len(b[0])
	if ar != bc {
		panic(fmt.Sprintf("invalid matrix mul ar:%v bc:%v", ar, bc))
	}
	r := is(br, 0)
	for i := 0; i < br; i++ {
		for j := 0; j < ar; j++ {
			r[i] += mmul(a[j], b[j][i])
			r[i] %= mod
		}
	}
	return r
}

func matrixpow(n int, matrix [][]int) [][]int {

	size := len(matrix)
	base := make([][][]int, maxlogn)
	base[0] = matrix
	for i := 0; i < maxlogn-1; i++ {
		base[i+1] = matrixmul(base[i], base[i])
	}
	r := i2s(size, size, 0)
	for i := 0; i < size; i++ {
		r[i][i] = 1
	}

	for i := 0; i < maxlogn; i++ {
		if hasbit(n, i) {
			r = matrixmul(r, base[i])
		}
	}
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

func (p point) isValid(h, w int) bool {
	return 0 <= p.x && p.x < h && 0 <= p.y && p.y < w
}

func (p point) dist(to point) float64 {
	return pointDist(p, to)
}

func (p point) getAngle(f point) float64 {
	return math.Atan2(float64(p.y-f.y), float64(p.x-f.x))*180/math.Pi + 180
}

func pointAdd(a, b point) point {
	return point{x: a.x + b.x, y: a.y + b.y}
}

func pointSub(a, b point) point {
	return point{x: a.x - b.x, y: a.y - b.y}
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

func pointInnerProduct(a, b point) int {
	return (a.x * b.y) - (b.x * a.y)
}

type line struct {
	x int
	y int
	c int
}

func newline(x, y, c int) line {
	return line{x, y, c}
}

func isSame(l1, l2 line) bool {
	return l1.x == l2.x && l1.y == l2.y && l1.c == l2.c
}

func isParallel(l1, l2 line) bool {
	return l1.x == l2.x && l1.y == l2.y && l1.c != l2.c
}

func NewlineByTwoPoint(x1, y1, x2, y2 int) line {
	l := line{}
	if x1 == x2 {
		l = line{1, 0, -x1}
	} else if y1 == y2 {
		l = line{0, 1, -y1}
	} else {
		l = line{y1 - y2, x2 - x1, -x2*y1 + x1*y1 + y2*x1 - y1*x1}
	}
	g := gcd(gcd(abs(l.x), abs(l.y)), abs(l.c))
	l.x /= g
	l.y /= g
	l.c /= g
	if l.x < 0 {
		l.x = -l.x
		l.y = -l.y
		l.c = -l.c
	}
	return l
}
func crossPoint(l1, l2 line) (int, int, bool) {
	x := l1.y*l2.c - l2.y*l1.c
	y := l2.x*l1.c - l1.x*l2.c
	d := l1.x*l2.y - l2.x*l1.y
	if d == 0 {
		return 0, 0, false
	}
	if d == 0 || x%d != 0 || y%d != 0 {
		return 0, 0, false
	}
	return x / d, y / d, true
}

// ==================================================
// bfs / dfs
// ==================================================

/*
snippet dfs "dfs"
var dfs func(v, p int)
dfs = func(v, p int) {
	for _, nv := range edges[v] {
		if nv.to == p {
			continue
		}
		dfs(nv.to, v)
	}
}
endsnippet

snippet bfs "bfs"
	q := list.New()
	q.PushBack(val)
	e := q.Front()
	for e != nil {
		t := e.Value.(int)

		// Do something

		e = e.Next()
	}
endsnippet

snippet bfsgrid "bfsgrid"
	q := list.New()
	q.PushBack(point{0, 0})
	e := q.Front()
	dx := []int{1, 0, -1, 0, 1, 1, -1, -1}
	dy := []int{0, 1, 0, -1, 1, -1, 1, -1}
	dist := i2s(h, w, inf)
	dist[0][0] = 0
	for e != nil {
		t := e.Value.(point)

		for k := 0; k < 4; k++ {
			np := point{t.x + dx[k], t.y + dy[k]}
			if !np.isValid(h, w) || dist[np.x][np.y] != inf {
				continue
			}
			dist[np.x][np.y] = dist[t.x][t.y] + 1
			q.PushBack(np)
		}

		e = e.Next()
	}
endsnippet
*/

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

type IntQueue struct {
	sum   int
	queue []int
	size  int
}

func newIntQueue() *IntQueue {
	return &IntQueue{}
}

func (iq *IntQueue) push(v int) {
	iq.queue = append(iq.queue, v)
	iq.sum += v
	//iq.sum = madd(iq.sum, v)
	iq.size++
}

func (iq *IntQueue) pop() int {
	v := iq.queue[0]
	iq.queue = iq.queue[1:]
	iq.sum -= v
	//iq.sum = madd(iq.sum, -v)
	iq.size--
	return v
}

func (iq *IntQueue) shrink(l int) {
	for {
		if iq.size <= l {
			break
		}
		iq.pop()
	}
}

func (iq *IntQueue) popBack() int {
	v := iq.queue[len(iq.queue)-1]
	iq.queue = iq.queue[:len(iq.queue)-1]
	iq.sum -= v
	//iq.sum = madd(iq.sum, -v)
	iq.size--
	return v
}

func (iq *IntQueue) isEmpty() bool {
	return iq.size == 0
}

// ==================================================
// heap
// ==================================================

/*
ih := newIntHeap(asc)
ih.Push(v)

	for !ih.IsEmpty() {
		v := ih.Pop()
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

// get sum f <= i && i < t
func (c *cusum) get(f, t int) int {
	if f > t || f >= c.l {
		return 0
	}
	return c.s[t] - c.s[f]
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
	x2++
	y2++
	return c.s[x2][y2] + c.s[x1][y1] - c.s[x1][y2] - c.s[x2][y1]
}

// ==================================================
// union find
// ==================================================

type unionFind struct {
	par     []int
	weights []int
}

func newUnionFind(n int) *unionFind {
	u := &unionFind{
		par:     make([]int, n),
		weights: make([]int, n),
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
	px := u.par[x]
	u.par[x] = u.root(px)
	u.weights[x] += u.weights[px]
	return u.par[x]
}

func (u *unionFind) unite(x, y int, arg ...int) {

	w := 0
	if len(arg) == 1 {
		w = arg[0]
	}
	w += u.weight(x)
	w -= u.weight(y)
	x = u.root(x)
	y = u.root(y)
	if x == y {
		return
	}
	if u.size(x) < u.size(y) {
		x, y = y, x
		w = -w
	}
	u.par[x] += u.par[y]
	u.par[y] = x
	u.weights[y] = w
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

func (u *unionFind) weight(x int) int {
	u.root(x)
	return u.weights[x]
}

func (u *unionFind) diff(x, y int) int {
	return u.weight(y) - u.weight(x)
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
type segstruct struct {
	v    int
	size int
}
*/

type segstruct int

type segfstruct int

type lazysegtree struct {
	n           int
	size        int
	log         int
	d           []segstruct
	lz          []segfstruct
	op          func(segstruct, segstruct) segstruct
	e           func() segstruct
	mapping     func(segfstruct, segstruct) segstruct
	composition func(segfstruct, segfstruct) segfstruct
	id          func() segfstruct
}

/*
	// 区間加算・区間和取得
	op := func(a segstruct, b segstruct) segstruct {
		return segstruct{a.v + b.v, a.size + b.size}
	}
	e := func() segstruct {
		return segstruct{0, 0}
	}
	id := func() segfstruct {
		return segfstruct(inf)
	}
	mapping := func(f segfstruct, x segstruct) segstruct {
		if f == id() {
			return x
		}
		return segstruct{x.v + int(f) * x.size, x.size}
	}
	compostion := func(f segfstruct, g segfstruct) segfstruct {
		if f == id() {
			return g
		}
		if g == id() {
			return f
		}
		return segfstruct(int(f) + int(g))
	}

	// 区間変更・区間最小値取得
	op := func(a segstruct, b segstruct) segstruct {
		return segstruct(min(int(a), int(b)))
	}
	e := func() segstruct {
		return segstruct(inf)
	}
	id := func() segfstruct {
		return segfstruct(inf)
	}
	mapping := func(f segfstruct, x segstruct) segstruct {
		if f == id() {
			return x
		}
		return segstruct(int(f))
	}
	compostion := func(f segfstruct, g segfstruct) segfstruct {
		if f == id() {
			return g
		}
		return f
	}

	lst := newlazysegtree(
		n,
		base,
		op,
		e,
		mapping,
		compostion,
		id,
	)
	lst.applyrange(f, t, segfstruct(v))
	iv := int(lst.get(i))
	rv := int(lst.prod(l,r)
	av := int(t.allprod()))
*/

func newlazysegtree(
	n int,
	v []segstruct,
	op func(segstruct, segstruct) segstruct,
	e func() segstruct,
	mapping func(segfstruct, segstruct) segstruct,
	composition func(segfstruct, segfstruct) segfstruct,
	id func() segfstruct,
) *lazysegtree {

	l := &lazysegtree{
		n:           n,
		op:          op,
		e:           e,
		mapping:     mapping,
		composition: composition,
		id:          id,
	}
	l.size = pow2(bitlen(n))
	l.log = bitlen(n)
	l.d = make([]segstruct, l.size*2)
	for i := range l.d {
		l.d[i] = e()
	}
	l.lz = make([]segfstruct, l.size)
	for i := range l.lz {
		l.lz[i] = id()
	}
	if len(v) > 0 {
		if len(v) != n {
			panic("invalid v value")
		}
		for i := 0; i < l.n; i++ {
			l.d[l.size+i] = v[i]
		}
		for i := l.size - 1; i >= 1; i-- {
			l.update(i)
		}
	}
	return l

}

func (l *lazysegtree) set(p int, x segstruct) {
	if p < 0 || p > l.n {
		panic(fmt.Sprintf("invalid p value n=%v p=%v", l.n, p))
	}
	p += l.size
	for i := l.log; i >= 1; i-- {
		l.push(p >> i)
	}
	l.d[p] = x
	for i := 1; i <= l.log; i++ {
		l.update(p >> i)
	}
}

func (l *lazysegtree) get(p int) segstruct {
	if p < 0 || p > l.n {
		panic(fmt.Sprintf("invalid p value n=%v p=%v", l.n, p))
	}
	p += l.size
	for i := l.log; i >= 1; i-- {
		l.push(p >> i)
	}
	return l.d[p]
}

func (l *lazysegtree) prod(le, ri int) segstruct {
	if le < 0 || le > l.n {
		panic(fmt.Sprintf("invalid le value n=%v le=%v", l.n, le))
	}
	if ri < 0 || ri > l.n {
		panic(fmt.Sprintf("invalid ri value n=%v ri=%v", l.n, ri))
	}
	if ri < le {
		panic(fmt.Sprintf("invalid ri value le=%v ri=%v", le, ri))
	}
	if le == ri {
		return l.e()
	}

	le += l.size
	ri += l.size

	for i := l.log; i >= 1; i-- {
		if ((le >> i) << i) != le {
			l.push(le >> i)
		}
		if ((ri >> i) << i) != ri {
			l.push((ri - 1) >> i)
		}
	}

	sml := l.e()
	smr := l.e()
	for {
		if le >= ri {
			break
		}
		if le&1 == 1 {
			sml = l.op(sml, l.d[le])
			le++
		}
		if ri&1 == 1 {
			ri--
			smr = l.op(l.d[ri], smr)
		}
		le >>= 1
		ri >>= 1

	}

	return l.op(sml, smr)
}

func (l *lazysegtree) allprod() segstruct {
	return l.d[1]
}

func (l *lazysegtree) apply(p int, f segfstruct) {
	if p < 0 || p > l.n {
		panic(fmt.Sprintf("invalid p value n=%v p=%v", l.n, p))
	}
	p += l.size
	for i := l.log; i >= 1; i-- {
		l.push(p >> i)
	}
	l.d[p] = l.mapping(f, l.d[p])
	for i := 1; i <= l.log; i++ {
		l.update(p >> i)
	}
}

func (l *lazysegtree) applyrange(le, ri int, f segfstruct) {
	if le < 0 || le > l.n {
		panic(fmt.Sprintf("invalid le value n=%v le=%v", l.n, le))
	}
	if ri < 0 || ri > l.n {
		panic(fmt.Sprintf("invalid ri value n=%v ri=%v", l.n, ri))
	}
	if ri < le {
		panic(fmt.Sprintf("invalid ri value le=%v ri=%v", le, ri))
	}

	if le == ri {
		return
	}

	le += l.size
	ri += l.size

	for i := l.log; i >= 1; i-- {
		if ((le >> i) << i) != le {
			l.push(le >> i)
		}
		if ((ri >> i) << i) != ri {
			l.push((ri - 1) >> i)
		}
	}

	{
		le2 := le
		ri2 := ri
		for {
			if le >= ri {
				break
			}
			if le&1 == 1 {
				l.allApply(le, f)
				le++
			}
			if ri&1 == 1 {
				ri--
				l.allApply(ri, f)
			}
			le >>= 1
			ri >>= 1
		}
		le = le2
		ri = ri2
	}

	for i := 1; i <= l.log; i++ {
		if ((le >> i) << i) != le {
			l.update(le >> i)
		}
		if ((ri >> i) << i) != ri {
			l.update((ri - 1) >> i)
		}
	}
}

func (l *lazysegtree) maxright(le int, g func(segstruct) bool) int {

	if le < 0 || le > l.n {
		panic(fmt.Sprintf("invalid le value n=%v ri=%v", l.n, le))
	}
	if !g(l.e()) {
		panic("invalid g func")
	}
	if le == l.n {
		return l.n
	}
	le += l.size
	for i := l.log; i >= 1; i-- {
		l.push(le >> i)
	}
	sm := l.e()
	for {
		for {
			if le%2 != 0 {
				break
			}
			le >>= 1
		}
		if !g(l.op(sm, l.d[le])) {
			for {

				if le >= l.size {
					break
				}
				l.push(le)
				le = (2 * le)
				if g(l.op(sm, l.d[le])) {
					sm = l.op(sm, l.d[le])
					le++
				}
			}
			return le - l.size
		}
		sm = l.op(sm, l.d[le])
		le++
		if (le & -le) == le {
			break
		}
	}
	return l.n
}

func (l *lazysegtree) maxleft(ri int, g func(segstruct) bool) int {
	if ri < 0 || ri > l.n {
		panic("invalid ri value")
	}
	if !g(l.e()) {
		panic("invalid g func")
	}

	if ri == 0 {
		return 0
	}
	ri += l.size
	for i := l.log; i >= 1; i-- {
		l.push((ri - 1) >> i)
	}
	sm := l.e()
	for {
		ri--
		for {
			if ri > 1 && (ri%2 == 1) {
			} else {
				break
			}
			ri >>= 1
		}
		if !g(l.op(l.d[ri], sm)) {
			for {
				if ri >= l.size {
					break
				}
				l.push(ri)
				ri = (2*ri + 1)
				if g(l.op(l.d[ri], sm)) {
					sm = l.op(l.d[ri], sm)
					ri--
				}
			}
			return ri + 1 - l.size
		}
		sm = l.op(l.d[ri], sm)
		if (ri & -ri) == ri {
			break
		}
	}
	return 0
}

func (l *lazysegtree) update(k int) {
	l.d[k] = l.op(l.d[2*k], l.d[2*k+1])
}

func (l *lazysegtree) allApply(k int, f segfstruct) {
	l.d[k] = l.mapping(f, l.d[k])
	if k < l.size {
		l.lz[k] = l.composition(f, l.lz[k])
	}
}

func (l *lazysegtree) push(k int) {
	l.allApply(2*k, l.lz[k])
	l.allApply(2*k+1, l.lz[k])
	l.lz[k] = l.id()
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

func reverseEdgeCost(edges [][]edge, from, i int) {
	redge := edges[from][i]
	t := edges[redge.to][redge.rev].cost
	edges[redge.to][redge.rev].cost = redge.cost
	edges[redge.from][i].cost = t
}

func eraseEdgeCost(edges [][]edge, from, i int) {
	redge := edges[from][i]
	edges[redge.to][redge.rev].cost = 0
	edges[redge.from][i].cost = 0
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

	q := []int{}
	q = append(q, s)
	ti := 0
	for {
		if ti >= len(q) {
			break
		}
		t := q[ti]

		for _, e := range g.edges[t] {
			if e.cost > 0 && g.level[e.to] < 0 {
				g.level[e.to] = g.level[t] + 1
				q = append(q, e.to)
			}
		}

		ti++
	}
}

func (g *graph) dinicdfs(v, t, f int) int {
	if v == t {
		return f
	}
	for i := g.iter[v]; i < len(g.edges[v]); i++ {
		e := g.edges[v][i]
		g.iter[v] = i

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
