package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(X, Y int) int {
	// もともとはられている切手で十分
	// はられている切手だけでは足りないパターン
	if X >= Y {
		return 0
	} else {
		cnt := 0
		for i := Y - X; i > 0; i -= 10 {
			cnt++
		}
		return cnt
	}
}

func main() {
	X, Y := ni(), ni()
	fmt.Println(solve(X, Y))
}

var sc = bufio.NewScanner(os.Stdin)

func ni() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
