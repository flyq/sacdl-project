package main

import "fmt"

func ShellSort( arr []int) []int {
	length := len(arr)
	h := 1
	for h < length/3 {
		h = h*3 + 1
	}
	for h >= 1 {
		for i := h; i < length; i++ {
			for j :=i; j >= h && arr[j] < arr[j-h]; j -= h {
				arr[j], arr[j-h] = arr[j-h], arr[j]
			}
		}
		h /= 3
	}
	return arr
}


func main() {
	arr := []int{1,8,8,1,0,6,8,7,8,3,0}
	ShellSort(arr)
	for i := 0; i < 11; i++{
		fmt.Print(arr[i])
	}
	fmt.Println()
}
