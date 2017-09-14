package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func sub(x, y float32) float32 {
	return x - y
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(sub(13.4, 499))
}
