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

func main() {
	sc.Split(bufio.ScanWords)
	s := nextStr()
	s_len := len(s)
	prefix_a := 0
	suffix_a := 0

	// 行頭のaの個数
	for i := 0; i < s_len; i++ {
		if s[i] == 'a' {
			prefix_a++
		} else {
			break
		}
	}

	// 行末のaの個数
	for i := s_len - 1; i >= 0; i-- {
		if s[i] == 'a' {
			suffix_a++
		} else {
			break
		}
	}

	// aのみで構成された文字列の場合OK
	if prefix_a == s_len {
		fmt.Println("Yes")
		return
	}

	// 行頭と行末のaの個数が合わない場合即時No
	if prefix_a > suffix_a {
		fmt.Println("No")
		return
	}

	for s_index := prefix_a; s_index < (s_len - suffix_a); s_index++ {
		if s[s_index] != s[prefix_a+s_len-suffix_a-s_index-1] {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
