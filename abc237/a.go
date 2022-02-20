package main

import (
	"fmt"
	"math"
)

func main() {
	var input int
	fmt.Scanf("%d", &input)
	if pow(-2, 31) <= input && input < pow(2, 31) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
