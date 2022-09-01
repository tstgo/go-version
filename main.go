/*
 * @Author: leoking
 * @Date: 2022-09-01 16:24:01
 * @LastEditTime: 2022-09-01 16:57:00
 * @LastEditors: leoking
 * @Description:
 */
package main

import (
	"fmt"

	"github.com/tstgo/calculator"
	"github.com/tstgo/conv"
)

func main() {
	fmt.Println("hello,go-version")
	fmt.Println(calculator.Add[int](1, 3))
	// 此方法不存在
	// fmt.Println(calculator.Divide(2,4))
	fmt.Println(conv.String(1.1))
}
