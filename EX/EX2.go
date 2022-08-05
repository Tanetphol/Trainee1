package EX

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

var order string
var product_code = map[string]int{}
var product_name []string
var product_price = map[string]int{}
var product_quantity = map[string]int{}

func Cal2() {
	product_code["milk"] = 1
	product_name = append(product_name, "milk")
	product_price["milk"] = 10
	product_quantity["milk"] = 1
	for {
		showstock()
	out:
		for {
			showmenu()
		out2:
			for {
				getorder()
				if order == "get" {
					//var boo1 bool = get()
					get()
					break out2
					/*if !boo1 {
						break out2
					}*/
					
				} else if order == "add" {
					add()
					break out
				} else if order == "exit" {
					return
				} else {
					fmt.Println("Try again")
				}

			}
		}

	}
}
func showstock() {
	fmt.Println("########   	 STOCK 		   ########\n")
	fmt.Println("    ---Product------Quantity---\n	")
	if len(product_quantity) != 0 {
		for key, value := range product_quantity {
			if value != 0 {
				fmt.Printf("	%s		%d		\n\n", key, value)
			}
		}
	} else {
		fmt.Println("Sorry, We have nothing in stock")
	}
}
func showmenu() {
	fmt.Println("#############################################")
	fmt.Print("	Menu		  Price \n\n")
	if len(product_price) != 0 {
		for key, value := range product_price {
			fmt.Printf("	%s     	   %d\n",key,value)

		}
	} else {
		fmt.Println("	ขออภัย สินค้าหมด")
	}
}
func getorder() {
	fmt.Println("\n--Command list--")
	fmt.Print("--get, To order--\n--add, To add new product--\n--exit, To close program--\n")
	fmt.Scanf("%s", &order)
	screen.Clear()
}
func get() bool {
	var getorder string
	var getprice int
	var change int
	abletobuy := make(map[string]int)
	fmt.Print("Input money : ")
	fmt.Scanf("%d", &getprice)
	screen.Clear()
	for key, value := range product_price {
		if getprice >= value {
			change = getprice - value
			abletobuy[key] = change
		}
	}
	if len(product_price) == 0{
		fmt.Println("สินค้าหมด")
		return false
	}
	if len(abletobuy) == 0 {
		fmt.Println("จำนวนเงินของท่านไม่เพียงพอ")
		time.Sleep(3*time.Second)
		screen.Clear()
		return false
	}
	showmenu()
	fmt.Print("Input product :")
	fmt.Scanf("%s", &getorder)
	screen.Clear()
	for key, value := range abletobuy {
		if getorder == key {
			product_quantity[key] -= 1
			fmt.Printf("Your product code : %d\n", product_code[key])
			fmt.Printf("Your change : %d\n\n", value)
			time.Sleep(5 * time.Second)
			screen.Clear()
		}
		if product_quantity[key] == 0 {
			deletemap(getorder)
		}
	}
	
	return true
}
func add() {
	var name string
	var price int
	var amount int
	fmt.Println("Input : Name,Price,Amount")
	fmt.Scanf("%s %d %d", &name, &price, &amount)
	if len(product_code) != 0 {
		product_code[name] = len(product_code) + 1
	} else {
		product_code[name] = 1
	}
	product_name = append(product_name, name)
	product_price[name] = price
	product_quantity[name] = amount
	fmt.Println(product_code)
	fmt.Println(product_name)
	fmt.Println(product_price)
	fmt.Println(product_quantity)
}
func deletemap(getorder string) {
	delete(product_code, getorder)
	delete(product_price, getorder)
	delete(product_quantity, getorder)
	for index, value := range product_name {
		if getorder == value {
			product_name[index] = ""
		}
	}
}
