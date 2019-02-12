package main

import "fmt"




// a -> value 1    ->    内存地址   	-> 			0xc000016080
// &符号作用	 	-> 		取址
// *符号作用	 	-> 		取值

func main() {
	var a int = 1
	var b *int = &a
	var c **int = &b														// **int 的意思就是c的取值的取值是一个int类型
	var x int = *b															// b的取值 -> 地址a的取值 -> 1

	fmt.Println("a = ", a)											// 值 a
	fmt.Println("&a = ", &a)										// 内存 a
	fmt.Println("*&a = ", *&*&*&a)							// 对 a 取它的地址 -> &a
	fmt.Println("b = ", b)											// 地址 a
	fmt.Println("&b = ", &b)										// 地址 b
	fmt.Println("*&b = ", *&b)									// 地址	b 的值, 也就是 地址 a
	fmt.Println("c = ", c)											// 地址	b
	fmt.Println("&c = ", &c)										// 地址	c
	fmt.Println("*&c = ", *&c)									// 地址	c 的值, 也就是 &b 即地址b
	fmt.Println("***&*&*&*&c = ", ***&*&*&*&c)	// ***&*&*&*&c 等价于 -> **c -> c的取值的取值 -> 地址b的取值的取值 -> 地址a的取值 -> a
	fmt.Println("**c = ", **c)									// **c -> c的取值的取值 -> 地址b的取值的取值 -> 地址a的取值 -> a -> 1
	fmt.Println("x = ", x)											// x -> *b -> *&a -> a
}
