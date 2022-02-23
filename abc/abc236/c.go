package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextStr() string {
	sc.Scan()
	return sc.Text()
}

// 頻出するYes No出力用
func printYesNo(b bool) {
	if b {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	stationNum, rapidNum := nextInt(), nextInt()

	// 各駅を初期化
	station := make([]string, stationNum)
	for i := 0; i < stationNum; i++ {
		station[i] = nextStr()
	}

	// 急行フラグ管理
	rapid := make(map[string]bool)
	for i := 0; i < rapidNum; i++ {
		rapid[nextStr()] = true
	}

	// 表示
	for _, s := range station {
		printYesNo(rapid[s])
	}
}
