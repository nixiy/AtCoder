package main

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextStr() string {
	sc.Scan()
	return sc.Text()
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func min(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func max(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
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

func init() {
	sc.Split(bufio.ScanWords)
}

// 等差数列の和
// ただし初項a、項数n、末項lがわかっている場合
func sumArithmeticProgression_l(n, a, l int) int {
	fmt.Println("n, a, l: ", n, a, l)
	return n * (a + l) / 2
}

// 等差数列の和
// ただし初項a、項数n、公差dがわかっている場合
func sumArithmeticProgression_d(n, a, d int) int {
	return (n / 2) * (2*a + (n-1)*d)
}

func mint(mod, n int64) int64 {
	mint := big.NewInt(int64(n))
	return mint.Mod(mint, big.NewInt(mod)).Int64()
}

// TODO テストケース通らずのコード
func main() {
	const MOD = 998244353
	// x 以下の正整数で、 x と桁数が同じものの数
	// n桁に着目すると
	// 1桁: 総数9個
	// 2桁: 総数90個
	// 3桁: 総数900個
	// 4桁: 総数9000個
	// つまりn桁の場合の総数は、9*10^(n-1)個と表現できる。
	// また、1桁以下の総数:9、2桁以下の総数: 99、3桁以下の総数: 999 個である事は自明である
	// → すなわちn桁の数値の総数は(10^n)-1である
	// また、とある数値xの個数は以下で求めることができる
	// x-10^(len(x)) + len(x)-1

	n := nextInt()
	ans := 0
	inv2 := invMod(2, MOD) // 2の逆元を求める
	for i := 1; i <= n; i *= 10 {
		x := min(n, i*10-1)
		y := x - i + 1
		ans += (y + 1) % MOD * (y % MOD) % MOD * inv2 % MOD
	}
	ans = (ans%MOD + MOD) % MOD
	fmt.Println(ans)
}

// mod M におけるaの逆元を求める
func invMod(a, M int) int {
	p, x, u := M, 1, 0
	for p != 0 {
		t := a / p
		a, p = p, a-t*p
		x, u = u, x-t*u
	}
	if x < 0 {
		x += M
	}
	return x
}
