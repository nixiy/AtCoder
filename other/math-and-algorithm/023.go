package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := ni()

	// 和の期待値 ⇔ 期待値の和

	// 青サイコロ
	blueAns := 0.0
	for i := 0; i < N; i++ {
		blueAns += float64(ni())
	}
	blueAns /= float64(N)
	// 赤サイコロ
	redAns := 0.0
	for i := 0; i < N; i++ {
		redAns += float64(ni())
	}
	redAns /= float64(N)

	fmt.Println(blueAns + redAns)

}

var sc = bufio.NewScanner(os.Stdin)

func ni() int           { sc.Scan(); return atoi(sc.Text()) }
func atoi(a string) int { i, _ := strconv.Atoi(a); return i }

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
