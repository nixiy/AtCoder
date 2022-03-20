package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(n int, a []int) (cnt int) {
	cMap := make(map[int]int)
	for i := 0; i < n; i++ {
		cMap[a[i]]++
	}

	// 愚直に計算すると200000C2= 19999900000となり時間内に終わらない。
	// しかし、積の法則を用いて{1,99999},{2,99998}...{50000}の場合の結果を合計する事で計算が可能になる。
	for cNum := range cMap {
		if cNum == 100000-cNum { // 数値が同一であれば、その数値の中から2枚選ぶ通りを加算する
			cnt += comb(cMap[cNum], 2)
		} else { // 数値が異なれば、それぞれの通りを乗算する
			cnt += cMap[cNum] * cMap[100000-cNum]
			// 既に施行の終わったカードを使用済みにする
			cMap[cNum] = 0
			cMap[100000-cNum] = 0
		}
	}

	return cnt
}

func main() {
	N := ni()
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = ni()
	}
	fmt.Println(solve(N, A))
}

// 組み合わせの数nCr
func comb(n, r int) int {
	if r == 0 {
		return 1
	} else {
		return comb(n, r-1) * (n - r + 1) / r
	}
}

var sc = bufio.NewScanner(os.Stdin)

func ni() int           { sc.Scan(); return atoi(sc.Text()) }
func atoi(a string) int { i, _ := strconv.Atoi(a); return i }

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
