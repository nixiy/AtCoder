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

func yes() {
	fmt.Println("Yes")
}

func no() {
	fmt.Println("No")
}

// 頻出するYes No出力用
func printYesNo(b bool) {
	if b {
		yes()
	} else {
		no()
	}
}

func init() {
	sc.Split(bufio.ScanWords)
}

func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}

func solve(s string) {
	s_len := len(s)

	// 行頭のaの個数
	lA := 0
	for i := 0; i < s_len; i++ {
		if s[i] == 'a' {
			lA++
		} else {
			break
		}
	}

	// 行末のaの個数
	rA := 0
	for i := s_len - 1; i >= 0; i-- {
		if s[i] == 'a' {
			rA++
		} else {
			break
		}
	}

	// aのみで構成された文字列の場合OK
	if lA == s_len {
		yes()
		return
	}

	// 右に多い分には左にaを足せるが、左に多い分は調整不可能な為No
	if lA > rA {
		no()
		return
	}

	T := s[lA : s_len-rA] // aを除いた文字列をスライスで取得
	printYesNo(T == reverse(T))
}

func main() {
	s := nextStr()
	solve(s)
}
