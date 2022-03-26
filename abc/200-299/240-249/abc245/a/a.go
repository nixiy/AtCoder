package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// true taka
// false aoki
func solve(takaH, takaM, aokiH, aokiM int) bool {
	takaMin := takaH*60 + takaM
	aokiMin := aokiH*60 + aokiM

	if takaMin <= aokiMin {
		return true
	} else {
		return false
	}
}

func main() {
	takaH, takaM, aokiH, aokiM := ni(), ni(), ni(), ni()

	if solve(takaH, takaM, aokiH, aokiM) {
		fmt.Println("Takahashi")
	} else {
		fmt.Println("Aoki")
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
