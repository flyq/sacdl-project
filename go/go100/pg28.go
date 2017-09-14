package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%d\n", clacAge(5))
}
func clacAge(n int) (age int) {
	if n == 1 {
		age = 10
	} else {
		age = clacAge(n-1) + 2
	}
	return
}
