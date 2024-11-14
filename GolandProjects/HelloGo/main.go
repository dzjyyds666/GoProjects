package main

import "fmt"

func main() {
	//fmt.Println("Hello Go")
	//a := 1
	//fmt.Println(a)

	//var num uint8 = 200
	//
	//fmt.Printf("num=%v 类型:%T\n", num, num)
	//fmt.Println(unsafe.Sizeof(num))
	//
	//// 精度丢失
	//d := 1129.6
	//fmt.Println(d * 100)

	a := 10
	b := 20
	a, b = b, a
	fmt.Println(a, b)
}
