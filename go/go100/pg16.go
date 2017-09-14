package main

import (
	"fmt"
)

func main() {
	var m, n, r, x int
	fmt.Scanf("%d%d", &m, &n)
	x = m * n
	for n != 0 {
		r = m % n
		m = n
		n = r
	}
	fmt.Printf("%d %d\n", m, x/m)
}
