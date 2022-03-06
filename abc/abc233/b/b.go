package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(L, R int, S string) (retStr string) {
	// prefix 反転する文字列 suffix に分かれる
	retStr += S[:L-1]
	retStr += reverseString(S[L-1 : R])
	retStr += S[R:]
	return retStr
}

func main() {
	L, R := ni(), ni()
	S := ns()
	fmt.Println(solve(L, R, S))
}

// 速度計測の結果、上記のreverseの10倍早かった
// 参考:
// 10000文字を100000ループ: 653ms
func reverseString(s string) string {
	reversed := make([]byte, len(s))
	for i := range reversed {
		reversed[i] = s[len(s)-1-i]
	}
	return string(reversed)
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

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
