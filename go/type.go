package main

import "fmt"

type Weekday int

const (
	Sunday Weekday = iota + 1
	Monday
	Tuesday
	Wednesday
	Thurday
	Firday
	Saturday
)

type Flags uint
const (
	FlagUp Flags = 1 << iota 	// 向上
	FlagBroadcast 						// 支持广播访问
	FlagLoopback	 						// 是环回接口
	FlagPointToPoint 					// 属于点对点链路
	FlagMulticast 						// 支持多路广播
)

func main() {
	fmt.Println(
		FlagUp,
		FlagBroadcast,
		FlagLoopback,
		FlagPointToPoint,
		FlagMulticast,
	)
}
