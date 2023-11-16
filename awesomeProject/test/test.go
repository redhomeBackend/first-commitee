package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i int
	i = 10
	var m int
	num := "tom"
	var a1, a2, a3 int
	var n1 = 100
	var n2 = 200
	var n3 = "Tom"
	var (
		n4 = 300
		n5 = "Mary"
	)
	fmt.Println("HelloWorld", i, m, "m=", num, "num=", a1, a2, a3)
	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3, "n4=", n4, "n5=", n5)
	var a, b = 1, 2
	var m1, m2 = "你", "好"
	fmt.Println("m1+m2=", m1+m2)
	fmt.Println("a+b=", a+b)
	var c = 1
	fmt.Println("c int: %T", c)
	var c1 byte = 10
	fmt.Println("c1 int：%T ， c1 %d", c1, unsafe.Sizeof(c1))
	num1 := 5.12
	num2 := 0.12
	num3 := 5.1234e2
	fmt.Println("num1=", num1, "num2=", num2, "num3", num3)
	fmt.Println("num3float: %T", num3)
	var d byte = 'd'
	var e byte = '0'
	fmt.Println("d=", d, "，e = ", e)
	fmt.Println("d = %c ， e = %c", d, e)
	var f int = '北'
	fmt.Println("f=%c， f3=%d", f)
	var n = false
	fmt.Println("n =", n)
	var address string = "北京长城 110 hello word"
	fmt.Println("address = " + address)
}
