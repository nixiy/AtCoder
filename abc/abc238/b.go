package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func max(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	rotateArray := make([]int, n)
	var pizza [360]bool // 切込みを入れた角度(0-indexed)を管理する
	currentAngle := 0   // 現在の12時方向の角度

	// 入力
	for i := 0; i < n; i++ {
		rotateArray[i] = nextInt()
	}

	// 初期12時位置に切り込みを入れてスタート
	pizza[0] = true
	for _, angle := range rotateArray {
		currentAngle += angle // 時計回りにangle度回転させる

		// 360度を超えていたら360度に収める
		// ex 390度 → 360度引く → 30度
		currentAngle %= 360
		pizza[currentAngle] = true // 現在角度に切り込みフラグを立てる
	}

	maxCentralAngle := 0 // 最大の角度
	angle := 0           // 切り込み間の中心角
	for i := 0; i < 360; i++ {
		if pizza[i] {
			maxCentralAngle = max(maxCentralAngle, angle) // 最大中心角の更新
			angle = 0
		}
		angle++
	}
	maxCentralAngle = max(maxCentralAngle, angle) // 最大中心角の更新
	fmt.Println(maxCentralAngle)
}
