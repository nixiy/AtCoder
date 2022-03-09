package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	peopleN, cardK, startA := ni(), ni(), ni()

	ans := startA - 1
	for i := 0; i < cardK; i++ {
		ans++
		if ans > peopleN {
			ans = 1
		}
	}
	fmt.Println(ans)
}

var sc = bufio.NewScanner(os.Stdin)

func ni() int           { sc.Scan(); return atoi(sc.Text()) }
func atoi(a string) int { i, _ := strconv.Atoi(a); return i }

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
