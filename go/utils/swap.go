package main

import "fmt"

// a -> value 1    ->    内存地址   	-> 			0xc000016080
// &符号作用	 	-> 		取址
// *符号作用	 	-> 		取值

func main() {
	var a = 3
	var b = 4
	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)
}

func swap(a, b *int) {
	*a, *b = *b, *a
}
