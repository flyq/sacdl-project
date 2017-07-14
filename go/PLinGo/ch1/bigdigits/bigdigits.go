package main

import (
       "fmt"
       "log"
       "os"
       "path/filepath"
)

var bigDigits = [][]string{
    {"0000000", "0     0", "0     0", "0     0", "0     0", "0     0", "0000000"},
    {"   1   ", "   1   ", "   1   ", "   1   ", "   1   ", "   1   ", "   1   "}, 
    {"2222222", "      2", "      2", "2222222", "2      ", "2      ", "2222222"},
    {"3333333", "      3", "      3", "3333333", "      3", "      3", "3333333"},
    {"4     4", "4     4", "4     4", "4444444", "      4", "      4", "      4"},
    {"5555555", "5      ", "5      ", "5555555", "      5", "      5", "5555555"},
    {"6666666", "6      ", "6      ", "6666666", "6     6", "6     6", "6666666"},
    {"7777777", "      7", "      7", "      7", "      7", "      7", "      7"},
    {"8888888", "8     8", "8     8", "8888888", "8     8", "8     8", "8888888"} ,
    {"9999999", "9     9", "9     9", "9999999", "      9", "      9", "9999999"},
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("usage: %s <whole-number>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	
	stringOfDigits := os.Args[1]
	for row := range bigDigits[0] { // for row := 0; row < len(bigDigits[0]); row++
		line := ""
		for column := range stringOfDigits {
			digit := stringOfDigits[column] - '0'
			if 0 <= digit && digit <= 9 {
				line += bigDigits[digit][row] + " "
			} else {
				log.Fatal("invalid whole number")
			}
		}    
		fmt.Println(line)
	}
}

