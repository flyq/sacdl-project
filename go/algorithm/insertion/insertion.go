package main

import "fmt"

func insertSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			j := i - 1
			temp := nums[i]
			for j >= 0 && nums[j] > temp {
				nums[j+1] = nums[j]
				j--
			}
			nums[j+1] =temp
		}
	}
}


func main() {
	arr := []int{1,8,8,1,0,6,8,7,8,3,0}
	insertSort(arr)
	for i := 0; i < 11; i++{
		fmt.Print(arr[i])
	}
	fmt.Println()
}

 
