package EXbeaw

import (
	"fmt"
)


func Ex1() {
	fmt.Println(jellyCutting(2,5,4))
}

func jellyCutting(a, b, c int) int {
	var res int
	area := []int {a,b,c}
	for {
		if area[0] == 1 && area[1] == 1 && area[2] == 1 {
			break
		}
		
		index := findmax(area)
		if area[index]%2 != 0 {
			area[index] = area[index] - 1
		}
		area[index] = area[index] / 2
		res += 1
	}
	return res
}
func findmax (area []int) (int){
	var max,ind int
	for index,value := range area{
		if value >= max{
			max = value
			ind = index
		}
	}
	return ind
}
