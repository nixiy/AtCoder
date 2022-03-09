package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func solve(N_nin, K_joui int, point [][3]int) (ans []bool) {
	p := make([]int, N_nin)     // 1-3日目の得点を合計した配列を持つ
	copyp := make([]int, N_nin) // K番目の得点を把握する用
	for i := 0; i < N_nin; i++ {
		p[i] = point[i][0] + point[i][1] + point[i][2]
	}

	// K位のポイントをメモ deepCopyしないと片方の変更がもう片方に波及してしまう(同一メモリを参照しているため)
	copy(copyp, p)
	sort.Slice(copyp, func(i, j int) bool {
		return copyp[i] > copyp[j]
	})
	kPT := copyp[K_joui-1]

	// 各得点に対してK位になれるかを判定
	for i := 0; i < N_nin; i++ {
		// 自身の得点+300が、K位の人より大きいか(同点でもK位になれるため=を含む
		if p[i]+300 >= kPT {
			ans = append(ans, true)
		} else {
			ans = append(ans, false)
		}
	}

	return ans
}

func main() {
	N_nin, K_joui := ni(), ni()
	point := make([][3]int, N_nin)
	for i := 0; i < N_nin; i++ {
		point[i][0] = ni()
		point[i][1] = ni()
		point[i][2] = ni()
	}

	ans := solve(N_nin, K_joui, point)
	for _, an := range ans {
		if an {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

var sc = bufio.NewScanner(os.Stdin)

func ni() int { sc.Scan(); return atoi(sc.Text()) }

func atoi(a string) int { i, _ := strconv.Atoi(a); return i }

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
