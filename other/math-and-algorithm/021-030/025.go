package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := ni()
	// 1,2
	exp1_2 := 0.0
	for i := 0; i < N; i++ {
		exp1_2 += float64(ni())
	}
	exp1_2 *= 1.0 / 3.0
	// 3,4,5,6
	exp3_4_5_6 := 0.0
	for i := 0; i < N; i++ {
		exp3_4_5_6 += float64(ni())
	}
	exp3_4_5_6 *= 2.0 / 3.0

	fmt.Println(exp1_2 + exp3_4_5_6)
}

var sc = bufio.NewScanner(os.Stdin)

func ni() int           { sc.Scan(); return atoi(sc.Text()) }
func atoi(a string) int { i, _ := strconv.Atoi(a); return i }

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
