package main

import (
	"fmt"
)

func main() {
	putchar(5)
}
func putchar(n int) {
	var c byte
	if n >= 1 {
		fmt.Scanf("%c", &c)
		putchar(n - 1)
		fmt.Printf("%c", c)
	}
}
