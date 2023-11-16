package main

import (
	"fmt"
	"strconv"
)

func main() {
	str1 := "10"
	num1, _ := strconv.ParseInt(str1, 2, 64)
	fmt.Printf("num1 type:%T, num1 = %v\n", num1, num1)

	str2 := "3.1415926"
	num2, _ := strconv.ParseFloat(str2, 64)
	fmt.Printf("num2 type:%T, num2 = %v\n", num2, num2)

	str3 := "true"
	b, _ := strconv.ParseBool(str3)
	fmt.Printf("b type:%T, b = %v\n", b, b)

	str4 := "z"
	myChar := byte(str4[0])
	fmt.Printf("myChar type:%T, myChar = %v\n", myChar, myChar)
}
