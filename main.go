/*
 * @Author: leoking
 * @Date: 2022-09-01 16:24:01
 * @LastEditTime: 2022-09-01 20:19:46
 * @LastEditors: leoking
 * @Description:
 */
package main

import (
	"fmt"

	"github.com/tstgo/calculator/v2"
	"github.com/tstgo/cond"
	"github.com/tstgo/conv"
)

func main() {
	fmt.Println("hello,go-version")
	fmt.Println(calculator.Add[int](1, 3))
	// 将dev分支的divide合并值main后此方法存在了
	fmt.Println(calculator.Divide(2, 0))
	fmt.Println(calculator.Sub(10, -2))
	fmt.Println(conv.String(1.1))
	fmt.Println(cond.IsZero(0))
}
