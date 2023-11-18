package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i int
	i = 10
	fmt.Println("i=", i)
	var num = 10.11
	fmt.Println("num=", num)
	numm := "tom"
	fmt.Println("numm=", numm)
	var a1, a2, a3 int
	fmt.Println("a1=", a1, "a2=", a2, "a3=", a3)
	var b1, b2, b3 = 100, "tim", 888
	fmt.Println("b1=", b1, "b2=", b2, "b3=", b3)
	c1, c2, c3 := 101, "tem", 666
	fmt.Println("c1=", c1, "c2=", c2, "c3=", c3)
	var a, b = 1, 2
	fmt.Println("a+b=", a+b)
	var n1, n2 = "你", "好"
	fmt.Println("n1+n2=", n1+n2)
	var aa = 1
	fmt.Printf("aa 的数据类型: %T", aa)
	var nn byte = 10
	fmt.Printf("nn 的数据类型:%T , n1占用的字节数是 %d ", nn, unsafe.Sizeof(nn))
	num1 := 5.12
	num2 := 0.12
	num3 := 5.1234e2
	fmt.Println("num1=", num1, "num2=", num2, "num3=", num3)
	fmt.Printf("num3的数据类型: %T", num3)
	var m byte = 'a'
	var n byte = '0'
	fmt.Printf("m = ", m, "n = ", n)
	fmt.Printf("m = %c , n = %c", m, n)
	var z int = '北'
	fmt.Printf("z=%c, z对应码值=%d", z, z)
	var d = false
	fmt.Println("d =", d)
	var address string = "北京长城 110 hello world"
	fmt.Println("address = ", address)
	var nn1 int32 = 100
	var nn2 float32 = float32(nn1)
	fmt.Printf("n1 = %v n2 = %v ", nn1, nn2)
}
