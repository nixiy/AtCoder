package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N, S := ni(), ni()
	cnt := 0
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			if i+j <= S {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}

var sc = bufio.NewScanner(os.Stdin)

func ni() int           { sc.Scan(); return atoi(sc.Text()) }
func atoi(a string) int { i, _ := strconv.Atoi(a); return i }

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
