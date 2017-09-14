package main

import (
	"fmt"
)

func main() {
	var array = [10]int{1, 4, 6, 9, 13, 16, 19, 28, 40, 100}
	count := len(array)
	for i := 0; i > count/2; i++ {
		array[i], array[count-i-1] = array[count-i-1], array[i]
	}
	fmt.Println(array)
}
