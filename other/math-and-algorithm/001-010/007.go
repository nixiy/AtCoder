package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N, X, Y := ni(), ni(), ni()
	cmCnt := 0
	for i := 1; i <= N; i++ {
		if i%X == 0 || i%Y == 0 {
			cmCnt++
		}
	}
	fmt.Println(cmCnt)
}

var sc = bufio.NewScanner(os.Stdin)

func ni() int           { sc.Scan(); return atoi(sc.Text()) }
func atoi(a string) int { i, _ := strconv.Atoi(a); return i }

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
