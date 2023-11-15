package main

import "fmt"

func main() {
	var num1 int = 10
	var num2 float32 = float32(num1)
	var str1 = fmt.Sprintf("%d", num1)
	fmt.Printf("str1 type:%T\nstr1 = %v\nnum2 type:%T\nnum2 =%v\n ", str1, str1, num2, num2)
}
