package main

import (
	"fmt"
)

func main() {
	var a, n, count int
	var sn, tn int
	fmt.Printf("please input a and n ")
	fmt.Scanf("%d%d\n", &a, &n)
	for count < n {
		tn = tn + a
		sn = sn + tn
		a = a * 10
		count++
	}
	fmt.Printf("a+aa+...=%d\n", sn)
}
