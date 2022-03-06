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
		return string(((s-'a')+byte(shift))%ALPHABET + 'a')
	} else if 'A' <= s && s <= 'Z' {
		return string(((s-'A')+byte(shift))%ALPHABET + 'A')
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
	existMap := make(map[int]int)
	for i := 0; i < len(S); i++ {
		d := diff(S[i], T[i])
		existMap[d] = d
		if len(existMap) > 1 {
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
