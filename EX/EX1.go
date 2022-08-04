package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var input string = "h99 ส!"

func main() {
	sum := new(int)
	defer func() {
		fmt.Print(*sum)
	}()
	input = strings.Replace(input, " ", "", -1)
	for _, r := range input {
		if r >= 'ก' && r <= 'ฮ' {
			fmt.Printf("ก,ข,ฃ,ค,ฅ,ฆ,ง,จ")
		} else if (r > 'a' && r < 'z') || (r > 'A' && r < 'Z') {
			fmt.Printf("A,B,C,D,E,F,G,H,I,J")
		} else if unicode.IsPunct(r) {
			fmt.Printf("%c", r)
		} else if unicode.IsNumber(r) {
			str := string(r)
			int1, _ := strconv.Atoi(str)
			*sum += int1
		}
	}
}
