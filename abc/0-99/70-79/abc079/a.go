package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func nextStr() string {
	sc.Scan()
	return sc.Text()
}

// string[i]のように取得するとbyteで取得できてしまう
// 中間処理でruneを使用して部分文字を取得する
func getRune(str string, index int) string {
	rs := []rune(str)
	return string(rs[index])
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	n := nextStr()
	if getRune(n, 0) == getRune(n, 1) && getRune(n, 1) == getRune(n, 2) || getRune(n, 1) == getRune(n, 2) && getRune(n, 2) == getRune(n, 3) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
