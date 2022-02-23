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

func nextStr() string {
	sc.Scan()
	return sc.Text()
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func max(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

// 数値配列をuniqして返す
func uniq(input []int) (uniq []int) {
	m := make(map[int]bool)
	for _, elm := range input {
		if !m[elm] {
			m[elm] = true
			uniq = append(uniq, elm)
		}
	}
	return uniq
}

// 等差数列の和
// ただし初項a、項数n、末項lがわかっている場合
func sumArithmeticProgression_l(n, a, l int) int {
	fmt.Println("n, a, l: ", n, a, l)
	return n * (a + l) / 2
}

// 等差数列の和
// ただし初項a、項数n、公差dがわかっている場合
func sumArithmeticProgression_d(n, a, d int) int {
	return (n / 2) * (2*a + (n-1)*d)
}

// string[i]のように取得するとbyteで取得できてしまう
// 中間処理でruneを使用して部分文字を取得する
func getStrForIndex(str string, index int) string {
	rs := []rune(str)
	return string(rs[index])
}

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	// code...
}
