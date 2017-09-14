package main

import (
	"fmt"
)

type B bool

func main() {
	var score int = 0
	fmt.Scanf("%d", &score)
	fmt.Println(B(score >= 90).C("A", B(score >= 60).C("B", "C")))
}

//Go语言没有三元表达式,自写函数模仿。
func (b B) C(t, f interface{}) interface{} {
	if bool(b) == true {
		return t
	}
	return f
}
