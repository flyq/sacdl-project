package main

import (
	//"os"
	//"strings"
	"fmt"
)

func selectSort(nums []int) {
	length := len(nums)
	for i := 0; i < length; i++ {
		maxIndex := 0

		for j := 1; j < length - i; j++ {
			if nums[j] > nums[maxIndex] {
				maxIndex = j
			}
		}
		nums[length-i-1], nums[maxIndex] = nums[maxIndex], nums[length-i-1]
	}
}

func main() {
	arr := []int{1,8,8,1,0,6,8,7,8,3,0}
	selectSort(arr)
	for i := 0; i < 11; i++{
		fmt.Print(arr[i])
	}
	fmt.Println()
}
