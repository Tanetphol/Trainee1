package EXbeaw

import "fmt"

func Ex2(){
	fmt.Println(findprime(5))
}

func findprime(x int) bool{
	if x == 0 || x == 1{
		return false
	}
	for i:= 2 ; i < x/2 ; i++ {
		if x%i == 0{
			return false
		}
	}
	return true
}