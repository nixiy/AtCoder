package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func lower_bound(a []int, x int) int {
	return sort.SearchInts(a, x)
}

func solve(N, K int, A []int) int {
	i := lower_bound(A, K)
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

func init() {
	sc.Split(bufio.ScanWords)
}
