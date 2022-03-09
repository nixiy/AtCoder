package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(S, T, X int) bool {
	// 日付をまたがない
	if S < T {
		if S <= X && X < T {
			return true
		}
	} else { // 日付をまたぐ
		if S <= X && X < 24 || 0 <= X && X < T {
			return true
		}
	}
	return false
}
func main() {
	S, T, X := ni(), ni(), ni()
	if solve(S, T, X) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
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
