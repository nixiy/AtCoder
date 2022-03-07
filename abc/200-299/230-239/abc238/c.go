package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MOD = 998244353
const INV2 = 499122177

func triangularNumber(x int) int {
	x %= MOD
	res := x
	res *= x + 1
	res %= MOD
	res *= INV2
	res %= MOD
	return res
}

func main() {
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

	n := ni()

	res := 0
	p10 := 10
	for dg := 1; dg <= 18; dg++ {
		l := p10 / 10
		r := min(n, p10-1)
		if l <= r {
			res += triangularNumber(r - l + 1)
			res %= MOD
		}
		p10 *= 10
	}

	fmt.Println(res)
}

var sc = bufio.NewScanner(os.Stdin)

func ni() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

// 小さい方を返す
func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func init() {
	sc.Split(bufio.ScanWords)
}
