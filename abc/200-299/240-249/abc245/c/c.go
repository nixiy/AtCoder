package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N, K := ni(), ni()
	A := nis(N)
	B := nis(N)

	dp := make([][2]bool, N) // dp[Xn][Ai, Bi]
	iA, iB := 0, 1
	dp[0][iA] = true
	dp[0][iB] = true

	// DPテーブルの作成
	for i := 1; i < N; i++ {
		// 1個前までのAについて成り立っていれば
		if dp[i-1][iA] {
			if abs(A[i-1]-A[i]) <= K {
				dp[i][iA] = true
			}
			if abs(A[i-1]-B[i]) <= K {
				dp[i][iB] = true
			}
		}

		// 1個前までのBについて成り立っていれば
		if dp[i-1][iB] {
			if abs(B[i-1]-A[i]) <= K {
				dp[i][iA] = true
			}
			if abs(B[i-1]-B[i]) <= K {
				dp[i][iB] = true
			}
		}
	}

	if dp[N-1][iA] || dp[N-1][iB] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
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

// 絶対値で返す
func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
