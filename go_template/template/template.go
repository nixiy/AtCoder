package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
}

var sc = bufio.NewScanner(os.Stdin)

func ni() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func ns() string {
	sc.Scan()
	return sc.Text()
}

// xのy乗
func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

// でかい方を返す
func max(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

// 小さい方を返す
func min(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func swap(a, b *int) {
	*b, *a = *a, *b
}

// 最大公約数: greatest common divisor
func gcd(a, b int) int {
	if a%b == 0 {
		return b
	} else {
		return gcd(b, a%b)
	}
}

// 最小公倍数
func lcm(a, b int) int {
	return a * b / gcd(a, b)
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
func chmin(p *int, v int) {
	if *p > v {
		*p = v
	}
}

// ex chmax(&p, v)
func chmax(p *int, v int) {
	if *p < v {
		*p = v
	}
}

// 等差数列の和
// ただし初項a、項数n、末項lがわかっている場合
func sumArithmeticProgression_l(n, first, last int) int {
	fmt.Println("n, first, last: ", n, first, last)
	return n * (first + last) / 2
}

// 等差数列の和
// ただし初項a、項数n、公差dがわかっている場合
func sumArithmeticProgression_d(n, first, diff int) int {
	return (n / 2) * (2*first + (n-1)*diff)
}

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
func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}

type intStack []int

func (stack *intStack) empty() bool {
	return len(*stack) == 0
}

func (stack *intStack) push(i int) {
	*stack = append(*stack, i)
}

func (stack *intStack) pop() int {
	result := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return result
}

type intQueue []int

func (queue *intQueue) empty() bool {
	return len(*queue) == 0
}

func (queue *intQueue) enqueue(i int) {
	*queue = append(*queue, i)
}

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
func lowerBound(a []int, target int) int {
	return sort.SearchInts(a, target)
}

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
	// 後続処理でXにYを付けるため、Yを小さくしておく
	if heightX < heightY {
		swap(&rootX, &rootY)
	}
	u.parent[rootY] = rootX
	u.rSize[rootX] = 1 + u.rSize[rootX] + u.rSize[rootY]
	if heightX == heightY {
		u.height[rootX]++
	}
}

func (u *UnionFind) Size(x int) int {
	return u.rSize[u.Root(x)] + 1
}

// aとbが同じ集合に属するか
func (u *UnionFind) IsSame(a, b int) bool {
	return u.Root(a) == u.Root(b)
}

func (u *UnionFind) Groups() map[int][]int {
	hash := make(map[int][]int)
	for i := 0; i < u.uSize; i++ {
		r := u.Root(i)
		hash[r] = append(hash[r], i)
	}
	return hash
}

func init() {
	sc.Split(bufio.ScanWords)
}
