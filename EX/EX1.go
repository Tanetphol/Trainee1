package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var input string = "h9999 ส!"

func main() {
	sum := 0
	defer func() {
		sum = reverseNumber(sum)
		fmt.Println(sum)

	}()

	input = strings.Replace(input, " ", "", -1)
	// check eng any character from a to z or A to Z
	eng, _ := regexp.Compile(`[a-zA-Z]`)
	// check eng any character from ก to ฮ
	thai, err := regexp.Compile(`[ก-ฮ]`)
	if err != nil {
		fmt.Println("regexp.Compile(`[ก-ฮ]`) ERROR : ", err)
	}

	for _, r := range input {
		if thai.MatchString(string(r)) {
			writethai('ก', 'จ')
		} else if eng.MatchString(string(r)) {
			writeeng('A', 'J')
		} else if unicode.IsPunct(r) {
			fmt.Printf("%c, ", r)
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
		re = append(re, ch, ',', ' ')
	}
	re2 := string(re)
	fmt.Print(re2)
}
func writethai(a, b rune) {
	var re []rune
	for ch := a; ch <= b; ch++ {
		re = append(re, ch, ',', ' ')
	}
	re2 := string(re)
	fmt.Print(re2)
}
