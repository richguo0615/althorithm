package main

import (
	"fmt"
	"strconv"
)

//http://puremonkey2010.blogspot.com/2011/05/c-bitwise-operation.html
//由於電腦進行位元運算比乘法、除法運算快上許多，所以有很多專業的程式設計師，會利用位元運算來取代乘法、除法運算. 優點是程式執行效率增加，缺點是程式碼可讀性變低
func bitwise() {
	var a uint64 = 60
	var b uint64 = 13
	var c uint64 = 0

	bitwiseSwap(1, 2)

	// 偶數 & 1 會是 0
	fmt.Println(2 & 1)

	// 奇數 & 1 會是 1
	fmt.Println(3 & 1)

	fmt.Printf("a: %s\nb: %s \n", strconv.FormatUint(a, 2), strconv.FormatUint(b, 2))

	c = a & b //AND 兩個位元都是1，才是1
	fmt.Printf("Line 1 - Value of c is %d - %s \n", c, strconv.FormatUint(c, 2))

	c = a | b //OR 只要有一個位元是1，就是1
	fmt.Printf("Line 2 - Value of c is %d - %s \n", c, strconv.FormatUint(c, 2))

	c = a ^ b //XOR 兩個位元一樣就得 0，
	fmt.Printf("Line 3 - Value of c is %d - %s \n", c, strconv.FormatUint(c, 2))

	c = a << 2 //相當於 a * 2 * 2 （無條件捨去
	fmt.Printf("Line 4 - Value of c is %d - %s \n", c, strconv.FormatUint(c, 2))

	c = a >> 2 //相當於 a / 2 / 2（無條件捨去
	fmt.Printf("Line 5 - Value of c is %d - %s \n", c, strconv.FormatUint(c, 2))

	c = a &^ b
	fmt.Printf("Line 6 - Value of c is %d - %s \n", c, strconv.FormatUint(c, 2))

	fmt.Println(2)
}

func bitwiseSwap(x, y int) {
	x = x ^ y
	y = x ^ y
	x = x ^ y
	fmt.Printf("x: %d y: %d \n", x, y)
}
