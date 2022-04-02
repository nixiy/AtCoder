package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func solve(N, Coupon, X int, A []int) int {
	// コスパが良くなるようにまずはクーポンを使う。
	ans := 0
	for i := 0; i < N; i++ {
		bestC := A[i] / X
		if bestC <= Coupon {
			A[i] = max(A[i]-(bestC*X), 0)
			Coupon -= bestC
			if Coupon == 0 {
				break
			}
		} else {
			A[i] = max(A[i]-(Coupon*X), 0)
			Coupon = 0
			break
		}

	}

	if Coupon > 0 {
		// 残りCouponで最も率の良いものから利用する。
		// そのため一旦降順にソートする
		sort.Slice(A, func(i, j int) bool { return A[i] > A[j] })

		for i := 0; i < N; i++ {
			A[i] = 0
			Coupon--
			if Coupon == 0 {
				break
			}
		}
	}

	for i := 0; i < N; i++ {
		ans += A[i]
	}

	return ans
}

func main() {
	N, Coupon, X := ni(), ni(), ni()
	A := nis(N)

	ans := solve(N, Coupon, X, A)
	fmt.Println(ans)
}

var sc = bufio.NewScanner(os.Stdin)

func ni() int           { sc.Scan(); return atoi(sc.Text()) }
func atoi(a string) int { i, _ := strconv.Atoi(a); return i }

func nis(n int) (a []int) {
	a = make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = ni()
	}
	return a
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
