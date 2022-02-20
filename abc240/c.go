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
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func nextStr() string {
	sc.Scan()
	return sc.Text()
}

func Powerset1(nums []int, posi int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}

	length := int(math.Pow(2, float64(len(nums))))
	result := make([][]int, length)

	for i := 0; i < length; i++ {
		bi := i
		s := []int{}
		for _, n := range nums {
			if n > posi {
				break
			}
			if bi%2 != 0 {
				s = append(s, n)
			}
			bi = bi / 2

		}
		result[i] = s
	}

	return result
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	posi := nextInt() // 正解の座標の位置

	b_a := make([]int, n)
	aTotal := 0

	for i := 0; i < n; i++ {
		a := nextInt()
		b := nextInt()
		aTotal += a
		b_a[i] = b - a
	}

	ret := Powerset1(b_a, posi)
	// a < bになる性質を利用して、aの総合計を出す。
	// aの総合計がposiを超えていたら即時No
	diff := posi - aTotal
	if diff < 0 {
		fmt.Println("No")
	} else if diff == 0 {
		fmt.Println("Yes")
	} else {
		for _, arr := range ret {
			arTotal := 0
			for _, ar := range arr {
				arTotal += ar
			}
			if arTotal == diff {
				fmt.Println("Yes")
				return
			}
		}
		fmt.Println("No")
	}
}
