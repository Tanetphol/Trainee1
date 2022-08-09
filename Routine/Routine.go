package Routine

import (
	"fmt"
	"time"
)

func buyGlassesAtSevenEleven(c chan string) {
	time.Sleep(1 * time.Second)
	fmt.Println("ซื้อแว่น : ที่เซเว่น : เสร็จแล้ว")
	c <-"ส่งของให้นายA"
}
func buyWatchAtCentral(c chan string) {
	time.Sleep(1 * time.Second)
	fmt.Println("ซื้อนาฬิกา : ที่เซ็นทรัล : เสร็จแล้ว")
}
func buyFruitAtSiamParagon() {
	time.Sleep(1 * time.Second)
	fmt.Println("ซื้อผลไม้ : ที่สยามพารากอน : เสร็จแล้ว")
}
func buyCarAtToyota() {
	time.Sleep(1 * time.Second)
	fmt.Println("ซื้อรถ : ที่ศูนย์โตโยต้า : เสร็จแล้ว")
}

func Routine() {
	// EX.Cal2()
	var messagefrommisterB string
	c := make(chan string)

	start := time.Now()          // เริ่มจับเวลาในการ Run
	go buyGlassesAtSevenEleven(c) //B
	go buyWatchAtCentral(c)       //B
	buyFruitAtSiamParagon()      //A
	buyCarAtToyota()             //A
	// result := ...
	// go sendtomisterA(c)
	messagefrommisterB  = <- c
	if messagefrommisterB == "ส่งของให้นายA"{
		fmt.Println("นายAได้รับของแล้ว")
		
	}
	
	fmt.Println("ใช้เวลาในการ Run ทั้งสิ้น : ", time.Since(start), " วินาที") // แสดงเวลาที่ Run ทั้งหมด
}
