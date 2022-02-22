package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	x := nextInt()
	// float64は符号ビット1、指数部11、仮数部52bitからなる
	// 2^52 → 4.5035996e+15 となり、せいぜい15桁程度までしか値が保証されない。
	// 987654321987654321 は17桁あり、float64にキャストして除算するともれなく値がずれる。
	// 単純に整数同士で除算した場合は、切り捨ての挙動になる
	// 負数の除算の場合、小数部を単に切り捨てられるので、-2.4 → -2 にされてしまう。
	// 切り捨て除算 → 除算値を超えない最大の整数 -2.4 → -3
	// よって負数の除算の場合 あまりが負の場合-1を引くことで実現する
	if x%10 < 0 {
		fmt.Println(x/10 - 1)
	} else {
		fmt.Println(x / 10)
	}
}
