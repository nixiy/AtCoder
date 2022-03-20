package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s [3]string
	var t [3]string

	s[0], s[1], s[2] = ns(), ns(), ns()
	t[0], t[1], t[2] = ns(), ns(), ns()

	matchCnt := 0
	for i := 0; i < 3; i++ {
		if s[i] == t[i] {
			matchCnt++
		}
	}

	// シミュレーションしたところ、一致個数で振り分け可能。
	// 3個: true
	// 2個: - (2個合っていたら3個合っているため存在しないパターン)
	// 1個: false (奇数回の入れ替えが必要)
	// 0個: true (6回入れ替えれば可能)
	// → 1個の場合以外Yes。
	if matchCnt == 1 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}

var sc = bufio.NewScanner(os.Stdin)

func ns() string { sc.Scan(); return sc.Text() }
func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
