/*
 * @Author: leoking
 * @Date: 2022-09-01 16:24:01
 * @LastEditTime: 2022-09-01 17:09:03
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
	// 将dev分支的divide合并值main后此方法存在了
	fmt.Println(calculator.Divide(2, 0))
	fmt.Println(conv.String(1.1))
}
