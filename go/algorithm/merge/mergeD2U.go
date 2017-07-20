package main

import "fmt"

var tmp []int

/**
 * merge sort from down to up
 */
func MergeDToU(arr []int) []int {
	n := len(arr)
	var sz, lo int
	tmp = make([]int, n)
	for sz=1; sz < n; sz += sz {
		for lo=0; lo<n-sz; lo += sz+sz {
			merge(arr, lo, lo+sz-1, min(lo+sz+sz-1, n-1))
		}
	}
	return arr
}

func min(a,b int) int {
	if a < b {
		return a
	} else {
		return b
	}
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
	MergeDToU(arr)
	for i := 0; i < 11; i++{
		fmt.Print(arr[i])
	}
	fmt.Println()
}
