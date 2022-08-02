package main

import (
	"fmt"
	"unicode"
)

var input string = "tanetpho0@"

func main() {
	for _, r := range input {
		if unicode.IsLetter(r) {
			fmt.Println("alphabet checked")
		} else if unicode.IsNumber(r) {
			fmt.Println("number คิคิ")
			fmt.Println("number checked")
		} else {
			fmt.Println("nor number or alphabet")
		}

	}
	fmt.Println("จบบ")
}
