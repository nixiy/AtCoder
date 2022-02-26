package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

const SearchMasuMax = 6

func main() {
	N := nextInt()
	S := make([][]string, N) // 盤面の2次元配列
	for i := 0; i < N; i++ {
		S[i] = strings.Split(nextStr(), "")
	}

	for tate := 0; tate < N; tate++ {
		for yoko := 0; yoko < N; yoko++ {
			// 横の走査
			if yoko+5 < N {
				sharpCheck(S, tate, yoko, 0, 1)
			}

			// 横の走査
			if tate+5 < N {
				sharpCheck(S, tate, yoko, 1, 0)
			}

			// 右下がりの走査
			if 0 <= tate-5 && yoko+5 < N {
				sharpCheck(S, tate, yoko, -1, 1)
			}

			// 右上がりの走査
			if tate+5 < N && yoko+5 < N {
				sharpCheck(S, tate, yoko, 1, 1)
			}
		}
	}

	// 全て引っかからなかったらno
	no()
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextStr() string {
	sc.Scan()
	return sc.Text()
}

func yes() {
	fmt.Println("Yes")
}

func no() {
	fmt.Println("No")
}

func init() {
	sc.Split(bufio.ScanWords)
}

func cntCheck(cnt int) {
	if cnt >= 4 {
		yes()
		os.Exit(0)
	}
}

func sharpCheck(S [][]string, tate, yoko, tateK, yokoK int) {
	cnt := 0
	// 6マス分数える
	for k := 0; k < SearchMasuMax; k++ {
		if S[tate+(k*tateK)][yoko+(k*yokoK)] == "#" {
			cnt++
		}
	}
	cntCheck(cnt)
}
