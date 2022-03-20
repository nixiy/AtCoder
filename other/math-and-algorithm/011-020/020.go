package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := ni()
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = ni()
	}

	// カードを5枚選ぶため、O(N^5) = 100^5 = 10^10 になり、一見計算が間に合わないように見える
	// しかし、100枚のカードから5枚を選ぶ方法は 100C5 = 75287520通りとなり、これは余裕で計算が完了する
	// よって5重ループで間に合う。
	// このようにcombinationを計算量の見積もりに利用する事も可能である
	cnt := 0
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			for k := j + 1; k < N; k++ {
				for l := k + 1; l < N; l++ {
					for m := l + 1; m < N; m++ {
						if A[i]+A[j]+A[k]+A[l]+A[m] == 1000 {
							cnt++
						}
					}
				}
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
