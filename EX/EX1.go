package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var input string = "h99 ส!"

func main() {
	sum := 0
	defer func() {
		sum = reverseNumber(sum)
		fmt.Print(sum)

	}()
	input = strings.Replace(input, " ", "", -1)
	for _, r := range input {
		if r >= 'ก' && r <= 'ฮ' {
			writethai('ก', 'จ')
		} else if (r > 'a' && r < 'z') || (r > 'A' && r < 'Z') {
			writeeng('A', 'J')
		} else if unicode.IsPunct(r) {
			fmt.Printf("%c", r)
		} else if unicode.IsNumber(r) {
			str := string(r)
			int1, _ := strconv.Atoi(str)
			sum += int1
		}
	}
}
func reverseNumber(num int) int {

	res := 0
	for num > 0 {
		remainder := num % 10
		res = (res * 10) + remainder
		num /= 10
	}
	return res
}
func writeeng(a, b rune) {
	var re []rune
	for ch := a; ch <= b; ch++ {
		re = append(re, ch)
	}
	re2 := string(re)
	fmt.Print(re2)
}
func writethai(a, b rune) {
	var re []rune
	for ch := a; ch <= b; ch++ {
		re = append(re, ch)
	}
	re2 := string(re)
	fmt.Print(re2)
}
