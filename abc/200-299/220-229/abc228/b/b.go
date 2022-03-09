package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N, X int, friends []int) int {
	knew := make(map[int]bool)
	// 秘密を知った友達i+1は、その秘密を友達friends[i+1]に教えます
	// これを既に知っている人に当たるまで繰り返す
	currentFriends := X - 1 // 初期ターゲットのindex
	for i := 0; i < N; i++ {
		if knew[currentFriends] {
			break
		}
		knew[currentFriends] = true                  // 現在の人を既知にする
		currentFriends = friends[currentFriends] - 1 // 現在の人が次に教える人の値-1(0indexed)を次のターゲットにする
	}
	return len(knew)
}

func main() {
	N, X := ni(), ni()
	friends := make([]int, N)
	for i := 0; i < N; i++ {
		friends[i] = ni()
	}
	fmt.Println(solve(N, X, friends))
}

var sc = bufio.NewScanner(os.Stdin)

func ni() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
