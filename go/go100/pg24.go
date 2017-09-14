package main

import (
	"fmt"
)

func main() {
	number := 20
	a, b, s := 2.0, 1.0, 0.0
	for n := 1; n <= number; n++ {
		s = s + a/b
		a, b = a+b, a
	}
	fmt.Printf("sum is %9.6f\n", s)
}
