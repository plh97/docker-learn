package main

import "fmt"

func initSlice() {
	s1 := []int{1, 2, 3}
	fmt.Println(s1, len(s1), cap(s1))

	var a *int = new(int)
	var b *int
	var c *int
	var d *int

	s2 := []*int{a, b, c, d}
	fmt.Println(s2, len(s2), cap(s2))
}

func makeSlice() {
	s1 := make([]int, 3)
	fmt.Println(s1, len(s1), cap(s1))

	// s2 := make([]int, 3, 2) // 容量必须大于长度
	// fmt.Println(s2, len(s2), cap(s2))
}

func emptySlice() {
	var nilSlice []int
	emptySlice := []int{}
	fmt.Printf("len = %d, cap = %d, %#v\n", len(nilSlice), cap(nilSlice), nilSlice)
	fmt.Printf("len = %d, cap = %d, %#v\n", len(emptySlice), cap(emptySlice), emptySlice)
}

func coypSlice() {
	data := [...]int{0, 1, 2, 3, 4, 9: 90}
	s := data[:2:3]
	fmt.Println("s: ", s, len(s), cap(s))
	s = append(s, 222, 333)
	fmt.Println("s: ", s, len(s), cap(s))
	fmt.Println("data: ", data, len(data), cap(data))
	fmt.Println(&s[0], &data[0])
}

func reCapSlice() {
	s := make([]int, 0, 1)
	c := cap(s)
	for i := 0; i < 50; i++ {
		s = append(s, i)
		if n := cap(s); n > c {
			fmt.Printf("cap: %d -> %d \n", c, n)
			c = n
		}
	}
}

func coypSlice2() {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := data[8:]
	s2 := data[:5]
	copy(s2, s)
	fmt.Println("s:", s)
	fmt.Println("s2:", s2)
	fmt.Println("data:", data)
}
