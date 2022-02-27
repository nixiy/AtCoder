package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
}

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

// xのy乗
func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

// でかい方を返す
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

// ex chmin(&p, v)
func chmin(p interface{}, v interface{}) {
	switch v.(type) {
	case int:
		a, ok := p.(*int)
		if !ok {
			return
		}
		if vv := v.(int); *a > vv {
			*a = vv
		}
	case float32:
		a, ok := p.(*float32)
		if !ok {
			return
		}
		if vv := v.(float32); *a > vv {
			*a = vv
		}
	case float64:
		a, ok := p.(*float64)
		if !ok {
			return
		}
		if vv := v.(float64); *a > vv {
			*a = vv
		}
	}
}

// ex chmax(&p, v)
func chmax(p interface{}, v interface{}) {
	switch v.(type) {
	case int:
		a, ok := p.(*int)
		if !ok {
			return
		}
		if vv := v.(int); *a < vv {
			*a = vv
		}
	case float32:
		a, ok := p.(*float32)
		if !ok {
			return
		}
		if vv := v.(float32); *a < vv {
			*a = vv
		}
	case float64:
		a, ok := p.(*float64)
		if !ok {
			return
		}
		if vv := v.(float64); *a < vv {
			*a = vv
		}
	}
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

// string[i]のように取得するとbyteで取得できてしまう
// 中間処理でruneを使用して部分文字を取得する
func getRune(str string, index int) string {
	rs := []rune(str)
	return string(rs[index])
}

// sのリバースを返す
// ex abcd → dcba
func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}

type intStack []int

func (stack *intStack) push(i int) {
	*stack = append(*stack, i)
}

func (stack *intStack) pop() int {
	result := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return result
}

type intQueue []int

func (queue *intQueue) enqueue(i int) {
	*queue = append(*queue, i)
}

func (queue *intQueue) dequeue() int {
	result := (*queue)[0]
	*queue = (*queue)[1:]
	return result
}

func init() {
	sc.Split(bufio.ScanWords)
}
