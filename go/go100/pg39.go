package main

import (
	"fmt"
)

func main() {
	var number int
	var array = [11]int{1, 4, 6, 9, 13, 16, 19, 28, 40, 100}
	fmt.Printf("insert a new number:")
	fmt.Scanf("%d", &number)
	array[10] = number
	for i := 10; i > 0; i-- {
		if array[i] < array[i-1] {
			array[i], array[i-1] = array[i-1], array[i]
		} else {
			break
		}
	}
	fmt.Println(array)
}
