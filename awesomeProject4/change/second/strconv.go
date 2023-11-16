package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1 int = 10
	str1 := strconv.FormatInt(int64(num1), 2)
	fmt.Printf("str1 type:%T, str1 = %v\n", str1, str1)

	var num2 float64 = 3.1415926535
	str2 := strconv.FormatFloat(num2, 'f', 5, 64)
	fmt.Printf("str2 type:%T, str2 = %v\n", str2, str2)

	var b bool = true
	str3 := strconv.FormatBool(b)
	fmt.Printf("str3 type:%T, str3 = %v\n", str3, str3)

	var myChar byte = 'z'
	str4 := string(myChar)
	fmt.Printf("str4 type:%T,str4 = %v\n", str4, str4)
}
