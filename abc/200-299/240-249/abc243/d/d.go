package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(X int, S string) int {
	var strStack StrStack

	// 不要箇所除去
	for _, sByte := range S {
		// LまたはRの直後にUがある場合、下がって上がるすなわち現在地が変わらないため、除去できる。
		if sByte == 'U' && len(strStack) > 0 && (strStack.last() == "L" || strStack.last() == "R") {
			strStack.pop()
		} else {
			strStack.push(string(sByte))
		}
	}

	// 解く
	for _, s := range strStack {
		if s == "U" {
			X /= 2
		}
		if s == "L" {
			X *= 2
		}
		if s == "R" {
			X = X*2 + 1
		}
	}
	return X
}

func main() {
	_, X := ni(), ni()
	S := ns()

	fmt.Println(solve(X, S))
}

var sc = bufio.NewScanner(os.Stdin)

type StrStack []string

func (stack *StrStack) empty() bool   { return len(*stack) == 0 }
func (stack *StrStack) first() string { return (*stack)[0] }
func (stack *StrStack) last() string  { return (*stack)[len(*stack)-1] }
func (stack *StrStack) push(i string) { *stack = append(*stack, i) }

func (stack *StrStack) pop() string {
	result := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return result
}

func ni() int           { sc.Scan(); return atoi(sc.Text()) }
func ns() string        { sc.Scan(); return sc.Text() }
func atoi(a string) int { i, _ := strconv.Atoi(a); return i }

func init() {
	const MaxBuf = 1024 * 1024
	sc.Buffer(make([]byte, MaxBuf), MaxBuf)
	sc.Split(bufio.ScanWords)
}
