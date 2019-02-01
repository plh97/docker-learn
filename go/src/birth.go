package main

import (
	// "fmt"
	"math"
)


/*
 * 生日碰撞问题
 * n个人都不在同一天生日的问题
 *
 * ++++++++++++++++++++++++++++++++++++++++++
 * + n个人				概率
 * ++++++++++++++++++++++++++++++++++++++++++
 * + 1个人  			365/365
 * + 2个人  			365/365 * 364/365
 * + 3个人  			365/365 * 364/365 * 363/365
 * + 10个人 			11%
 * + 20个人 			41%
 * + 30个人 			71%
 * + 40个人 			89%
 * + 50个人 			97%
 * + 60个人 			99%
 * + 70个人 			99.916%
 * + 80个人 			99.991433%
 * + 90个人 			99.999385%
 * + 100个人			99.999969%
 * ++++++++++++++++++++++++++++++++++++++++++
 * + n个人  (!365)/(!n)/(365^n)
 * ++++++++++++++++++++++++++++++++++++++++++
 */


func surprise(n,limit float64) float64 {
	if(n>limit){
		return n * surprise(n-1,limit);
	} else {
		return 1;
	}

}


func birth(num int) float64 {
	/* 声明局部变量 */
	mother := math.Pow(float64(365), float64(num))
	children := surprise(float64(365),float64(365-num))





	// fmt.Printf("分子 %f \n",children);
	// fmt.Printf("分母 %f \n",mother);



	return 1 - children / mother
}