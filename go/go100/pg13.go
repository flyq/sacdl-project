package main

import (
	"fmt"
)

func main() {
	for num := 100; num < 1000; num++ {
		i := num / 100
		j := num / 10 % 10
		k := num % 10
		if i*i*i+j*j*j+k*k*k == num {
			fmt.Printf("%d^3+%d^3+%d^3=%d\n", i, j, k, num)
		}
	}
}
