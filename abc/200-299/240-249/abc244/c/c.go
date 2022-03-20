package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := ni()
	numMap := make(map[int]bool) // 既に使われた数字の管理

	// 2N+1だと最後に0を受け取るループが足りなかったため+2にしている
	for i := 1; i <= 2*N+2; i++ {
		if i%2 == 1 { // 高橋くん
			num := 0
			// 未使用数値を1個選ぶ
			for i := 1; i <= 2*N+2; i++ {
				if !numMap[i] {
					num = i
					numMap[i] = true
					break
				}
			}
			fmt.Println(num)
		} else { // 青木くん
			aoki := ni()
			if aoki == 0 {
				return
			} else {
				numMap[aoki] = true
			}
		}
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
