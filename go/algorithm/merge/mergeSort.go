package main

import "fmt"

var tmp []int
func MergeSort(arr []int) []int {
	tmp = make([]int, len(arr))
	mergeInnerSort(arr, 0, len(arr) - 1)
	return arr
}

func mergeInnerSort(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}
	if hi-lo == 1 {
		if arr[lo] > arr[hi] {
			arr[lo], arr[hi] = arr[hi], arr[lo]
		}
		return
	}
	mid := (lo + hi) / 2
	mergeInnerSort(arr, lo, mid)
	mergeInnerSort(arr, mid + 1, hi)
	merge(arr, lo, mid, hi)
}

func merge (arr []int, lo, mid, hi int) {
	for i := lo; i <= hi; i++ {
		tmp[i] = arr[i]
	}

	j := mid + 1
	for k := lo; k <= hi; k++ {
		if lo > mid {
			arr[k] = tmp[j]
			j++
		} else if j > hi {
			arr[k] = tmp[lo]
			lo++
		} else if tmp[lo] > tmp[j] {
			arr[k] = tmp[j]
			j++
		} else {
			arr[k] = tmp[lo]
			lo++
		}
	}
}


func main() {
	arr := []int{1,8,8,1,0,6,8,7,8,3,0}
	MergeSort(arr)
	for i := 0; i < 11; i++{
		fmt.Print(arr[i])
	}
	fmt.Println()
}
