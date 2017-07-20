package main

import "fmt"


func main() {
	arr := []int{1,8,8,1,0,6,8,7,8,3,0}
	quickSort(arr)
	for i := 0; i < 11; i++{
		fmt.Print(arr[i])
	}
	fmt.Println()
}

func quickSort(arr []int) []int {
	innerQuickSort(arr, 0, len(arr) - 1)
	return arr
}

func innerQuickSort(arr []int, lo, hi int) {
	fmt.Println("lo, hi", lo, hi)
	if lo >= hi { return }
	mid := innerParrtition(arr, lo, hi)
	innerQuickSort(arr, lo, mid)
	innerQuickSort(arr, mid + 1, hi)
}

func innerParrtition(arr []int, lo, hi int) int {
	i, j := lo + 1, hi
	k := arr[lo]
	for {
		for i < hi && k >= arr[i] {
			i++
		}
		for j > lo && k <= arr[j] {
			j--
		}
		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[lo], arr[j] = arr[j], arr[lo]
	return j
}
	
