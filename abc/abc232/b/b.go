package main

import (
	"bufio"
	"fmt"
	"os"
)

const ALPHABET = 26

// 対象の1文字をずらす
func byteShift(s byte, shift int) string {
	if shift < 0 {
		shift += ALPHABET
	}

	if 'a' <= s && s <= 'z' {
		return string(((s-'a')+byte(shift))%26 + 'a')
	} else if 'A' <= s && s <= 'Z' {
		return string(((s-'A')+byte(shift))%26 + 'A')
	} else {
		return string(s)
	}
}

// sをshift文字後ろにずらす O(len(s))
func strShift(s string, shift int) string {
	if shift < 0 {
		shift += ALPHABET
	}

	shiftedStr := ""
	for i := 0; i < len(s); i++ {
		shiftedStr += byteShift(s[i], shift)
	}
	return shiftedStr
}

// aからbまで何文字離れているか
func diff(a, b byte) int {
	return (int(b) + ALPHABET - int(a)) % ALPHABET
}

func solve(S, T string) bool {
	d := diff(S[0], T[0])
	for i := 1; i < len(S); i++ {
		// 文字シフト数が1文字目と異なれば打ち切る
		if diff(S[i], T[i]) != d {
			return false
		}
	}
	return true
}
func main() {
	S, T := ns(), ns()
	if solve(S, T) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

var sc = bufio.NewScanner(os.Stdin)

func ns() string {
	sc.Scan()
	return sc.Text()
}

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
