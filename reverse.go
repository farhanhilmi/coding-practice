package main

import (
	"fmt"
	"strings"
)

func reverseStr(str string) {
	for i := len(str) - 1; i >= 0; i-- {
		fmt.Print(string(str[i]))
	}
}

func reverseWord(text string) {
	words := strings.Split(text, " ")
	for _, v := range words {
		for i := len(v) - 1; i >= 0; i-- {
			fmt.Print(string(v[i]))
		}
		fmt.Print(" ")
	}
}

// func main() {
// 	str := "Ini budi dan yakinkan lah lula"
// 	// reverseStr(str)
// 	reverseWord(str)
// }
