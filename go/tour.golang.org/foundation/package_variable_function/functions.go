package main

import "fmt"

func add(x int, y int) (sum int) {
	sum = x + y
	return  sum
}

func main() {
	fmt.Println(add(42, 13))
}
