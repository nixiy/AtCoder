package main

import (
	"bufio"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
}

var sc = bufio.NewScanner(os.Stdin)

func ns() string        { sc.Scan(); return sc.Text() }
func ni() int           { sc.Scan(); return atoi(sc.Text()) }
func atoi(a string) int { i, _ := strconv.Atoi(a); return i }
func itoa(i int) string { return strconv.Itoa(i) }

func pow(x, y int) int { return int(math.Pow(float64(x), float64(y))) }

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

// 小さい方を返す
func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

// 絶対値で返す
func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func sqrt(x int) float64 { return math.Sqrt(float64(x)) }

// 変数の中身を入れ替える
func swap(a, b *int) { *b, *a = *a, *b }

// 最大公約数: greatest common divisor
func gcd(a, b int) int {
	if a%b == 0 {
		return b
	} else {
		return gcd(b, a%b)
	}
}

// 3つ以上の値の最大公約数
func multiGcd(target []int) (ans int) {
	if len(target) <= 1 {
		ans = -1
	} else {
		pastGcd := gcd(target[0], target[1])
		for i := 2; i < len(target); i++ {
			pastGcd = gcd(pastGcd, target[i])
		}
		ans = pastGcd
	}
	return ans
}

// 最小公倍数
func lcm(a, b int) int { return b / gcd(a, b) * a }

// 3つ以上の値の最小公倍数
func multiLcm(target []int) (ans int) {
	if len(target) <= 1 {
		ans = -1
	} else {
		pastLcm := lcm(target[0], target[1])
		for i := 2; i < len(target); i++ {
			pastLcm = lcm(pastLcm, target[i])
		}
		ans = pastLcm
	}
	return ans
}

// 数値配列をuniqして返す
func uniq(input []int) (uniq []int) {
	m := make(map[int]bool)
	for _, elm := range input {
		if !m[elm] {
			m[elm] = true
			uniq = append(uniq, elm)
		}
	}
	return uniq
}

// ex chmin(&p, v)
func chmin(p *int, v int) bool {
	if *p > v {
		*p = v
		return true
	}
	return false
}

// ex chmax(&p, v)
func chmax(p *int, v int) bool {
	if *p < v {
		*p = v
		return true
	}
	return false
}

// 等差数列の和
// ただし初項a、項数n、末項lがわかっている場合
func sumArithmeticProgression_l(n, first, last int) int { return n * (first + last) / 2 }

// 等差数列の和
// ただし初項a、項数n、公差dがわかっている場合
func sumArithmeticProgression_d(n, first, diff int) int { return (n / 2) * (2*first + (n-1)*diff) }

// string[i]のように取得するとbyteで取得できてしまう
// 中間処理でruneを使用して部分文字を取得する
func getRune(str string, index int) string {
	rs := []rune(str)
	if 0 <= index && index < len(str) {
		return string(rs[index])
	} else {
		return ""
	}
}

// sのリバースを返す
// ex abcd → dcba
// 10000文字を100000ループ: 7313ms
func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}

// 速度計測の結果、上記のreverseの10倍早かった
// 参考:
// 10000文字を100000ループ: 653ms
func reverseString(s string) string {
	reversed := make([]byte, len(s))
	for i := range reversed {
		reversed[i] = s[len(s)-1-i]
	}
	return string(reversed)
}

type intStack []int

func (stack *intStack) empty() bool { return len(*stack) == 0 }
func (stack *intStack) first() int  { return (*stack)[0] }
func (stack *intStack) last() int   { return (*stack)[len(*stack)-1] }
func (stack *intStack) push(i int)  { *stack = append(*stack, i) }

func (stack *intStack) pop() int {
	result := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return result
}

type intQueue []int

func (queue *intQueue) empty() bool   { return len(*queue) == 0 }
func (queue *intQueue) enqueue(i int) { *queue = append(*queue, i) }

func (queue *intQueue) dequeue() int {
	result := (*queue)[0]
	*queue = (*queue)[1:]
	return result
}

// 10進数xを文字列2進数で返す
func itob(x int) string {
	s := ""
	for x > 0 {
		s = strconv.Itoa(x%2) + s
		x /= 2
	}
	return s
}

// target以上を満たす最小の位置を0-basedの添字で返す
func lowerBound(a []int, target int) int { return sort.SearchInts(a, target) }

// targetより大きい最小の位置を0-basedの添字で返す
func upperBound(a []int, target int) int {
	return sort.Search(len(a), func(i int) bool { return a[i] > target })
}

func Permute(nums []int) [][]int {
	n := factorial(len(nums))
	ret := make([][]int, 0, n)
	permute(nums, &ret)
	return ret
}

func permute(nums []int, ret *[][]int) {
	*ret = append(*ret, makeCopy(nums))

	n := len(nums)
	p := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		p[i] = i
	}
	for i := 1; i < n; {
		p[i]--
		j := 0
		if i%2 == 1 {
			j = p[i]
		}

		nums[i], nums[j] = nums[j], nums[i]
		*ret = append(*ret, makeCopy(nums))
		for i = 1; p[i] == 0; i++ {
			p[i] = i
		}
	}
}

func factorial(n int) int {
	ret := 1
	for i := 2; i <= n; i++ {
		ret *= i
	}
	return ret
}

func makeCopy(nums []int) []int {
	return append([]int{}, nums...)
}

type UnionFind struct {
	parent []int // n個目の要素がどの親に属するか
	height []int // 木の高さ-1
	rSize  []int // 親配下の要素数(親自身を除く)
	uSize  int   // UnionFind全体の要素数
}

func NewUnionFind(size int) UnionFind {
	uf := UnionFind{}
	uf.uSize = size
	uf.parent = make([]int, size)
	uf.rSize = make([]int, size)
	uf.height = make([]int, size)
	for i := 0; i < uf.uSize; i++ {
		uf.parent[i] = i
	}
	return uf
}

// 木の根を求める
func (u *UnionFind) Root(x int) int {
	// speed 優先
	if u.parent[x] == x {
		return x
	}
	u.parent[x] = u.Root(u.parent[x]) // 経路圧縮
	return u.parent[x]
}

// xとyの属する集合をマージ
func (u *UnionFind) Merge(x, y int) {
	rootX, rootY := u.Root(x), u.Root(y)

	// 同じ親であれば既にマージ済み
	if rootX == rootY {
		return
	}

	heightX, heightY := u.height[rootX], u.height[rootY]
	// より大きい高さの木のroot配下に小さい木を接続する事で計算量を抑える: Union by Rank
	if heightX < heightY {
		swap(&rootX, &rootY)
	}
	u.parent[rootY] = rootX // 木のマージ
	u.rSize[rootX] = 1 + u.rSize[rootX] + u.rSize[rootY]

	// 同一の高さの木をマージした場合、付け加えられた側の高さが1高くなる
	if heightX == heightY {
		u.height[rootX]++
	}
}

// 自身が属する集合に何要素あるか(自身を含む)
func (u *UnionFind) Size(x int) int { return u.rSize[u.Root(x)] + 1 }

// aとbが同じ集合に属するか
func (u *UnionFind) IsSame(a, b int) bool { return u.Root(a) == u.Root(b) }

// どのrootがどの集合を持っているかを取得する
func (u *UnionFind) Groups() map[int][]int {
	hash := make(map[int][]int)
	for i := 0; i < u.uSize; i++ {
		r := u.Root(i)
		hash[r] = append(hash[r], i)
	}
	return hash
}

const ALPHABET = 26

// 対象の1文字をずらす
func byteShift(s byte, shift int) string {
	if shift < 0 {
		shift += ALPHABET
	}

	if 'a' <= s && s <= 'z' {
		return string(((s-'a')+byte(shift))%ALPHABET + 'a')
	} else if 'A' <= s && s <= 'Z' {
		return string(((s-'A')+byte(shift))%ALPHABET + 'A')
	} else {
		return string(s)
	}
}

// sをshift文字後ろにずらす O(len(s))
func strShift(s string, shift int) string {
	if shift < 0 {
		shift += ALPHABET
	}

	shiftedStr := ""
	for i := 0; i < len(s); i++ {
		shiftedStr += byteShift(s[i], shift)
	}
	return shiftedStr
}

// aからbまで何文字離れているか
func diff(a, b byte) int { return (int(b) + ALPHABET - int(a)) % ALPHABET }

// エラトステネスの篩を用いてn以下の素数を返す
func prime(N int) (p []int) {
	isPrimeSlice := make([]bool, N+1)
	for i := 2; i <= N; i++ {
		isPrimeSlice[i] = true
	}

	// エラトステネスの篩
	for i := 2; i*i <= N; i++ {
		if isPrimeSlice[i] {
			for x := 2 * i; x <= N; x += i {
				isPrimeSlice[x] = false
			}
		}
	}

	for i := 2; i <= N; i++ {
		if isPrimeSlice[i] {
			p = append(p, i)
		}
	}

	return p
}

// 高速な素数判定
func isPrime(N int) bool {
	if N < 2 {
		return false
	} else {
		// 2 - √N まで調べれば良い
		for i := 2; i*i <= N; i++ {
			if N%i == 0 {
				return false
			}
		}
		return true
	}
}

// 高速な約数列挙(未ソート)
func divisor(N int, isSort bool) (div []int) {
	for i := 1; i*i <= N; i++ {
		if N%i != 0 {
			continue
		}
		div = append(div, i)
		if i != N/i {
			div = append(div, N/i)
		}
	}
	if isSort {
		sort.Ints(div)
	}
	return div
}

// 高速な素因数分解
func factorization(N int) (f []int) {
	for i := 2; i*i <= N; i++ {
		for N%i == 0 {
			N /= i
			f = append(f, i)
		}
	}
	// 素数の最小値以上が残っていれば追加
	if N >= 2 {
		f = append(f, N)
	}
	return f
}

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
