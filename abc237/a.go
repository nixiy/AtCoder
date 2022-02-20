package main

import (
	"fmt"
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
