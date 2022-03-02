package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func lower_bound(a []int, x int) int {
	return sort.Search(len(a), func(i int) bool { return x <= a[i] })
}

func solve(N, K int, A []int) int {
	i := lower_bound(A, K)
	fmt.Println(i)
	if i < N {
		return i
	} else {
		return -1
	}
}

func main() {
	N, K := ni(), ni()
	A := make([]int, N)

	fmt.Println(solve(N, K, A))
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

func yes() {
	fmt.Println("Yes")
}

func no() {
	fmt.Println("No")
}

// 頻出するYes No出力用
func printYesNo(b bool) {
	if b {
		yes()
	} else {
		no()
	}
}

func init() {
	sc.Split(bufio.ScanWords)
}
