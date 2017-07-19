package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	world := "world"
	if len(os.Args) > 1 {
		world = strings.Join(os.Args[1:]," ")
	}
	fmt.Println("Hello", world)
}
	
	
