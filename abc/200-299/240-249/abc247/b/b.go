package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int, Sei, Name []string) bool {
	for i := 0; i < N; i++ {
		seiFailCnt := 0
		nameFailCnt := 0
		for j := 0; j < N; j++ {
			// 名字があだ名の場合
			if i != j {
				if Sei[i] != Sei[j] && Sei[i] != Name[j] {
				} else {
					seiFailCnt++
				}

				if Name[i] != Sei[j] && Name[i] != Name[j] {
				} else {
					nameFailCnt++
				}

			}
		}
		if seiFailCnt == 0 || nameFailCnt == 0 {
		} else {
			return false
		}
	}
	return true
}

func main() {
	N := ni()
	Sei := make([]string, N)
	Name := make([]string, N)

	for i := 0; i < N; i++ {
		Sei[i] = ns()
		Name[i] = ns()
	}

	isSuccess := solve(N, Sei, Name)

	if isSuccess {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

var sc = bufio.NewScanner(os.Stdin)

func ns() string        { sc.Scan(); return sc.Text() }
func ni() int           { sc.Scan(); return atoi(sc.Text()) }
func atoi(a string) int { i, _ := strconv.Atoi(a); return i }

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
