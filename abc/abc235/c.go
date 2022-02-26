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

func init() {
	sc.Split(bufio.ScanWords)
}

func main() {
	seqLen, pairLen := nextInt(), nextInt()
	seq := make([]int, seqLen)
	x := make([]int, pairLen)
	k := make([]int, pairLen)
	// 1 1 2 3 1 2 の場合
	// 以下の様な数値がそれぞれ何番目に登場するのかを記録しておく
	// 2次元配列で宣言すると10億x10億の配列を作ることになり大変な事になる(やってしまった)
	// map[0:[1 2 5] 1:[3 6] 2:[4]]
	mat := make(map[int][]int)

	for i := 0; i < seqLen; i++ {
		seq[i] = nextInt()
	}
	for i := 0; i < pairLen; i++ {
		x[i] = nextInt()
		k[i] = nextInt()
	}

	for i := 0; i < seqLen; i++ {
		// matrix[今見ている数]にi+1を追加する
		mat_i := seq[i] - 1
		mat[mat_i] = append(mat[mat_i], i+1)
	}

	fmt.Println(mat)

	for pi := 0; pi < pairLen; pi++ {
		xp := x[pi] - 1
		kp := k[pi] - 1
		// 対象の数値がseqに存在するか
		// 対象のkpが配列のindex内に収まるか
		if len(mat[xp]) != 0 && len(mat[xp]) > kp {
			fmt.Println(mat[xp][kp])
		} else {
			fmt.Println(-1)
		}
	}
}
