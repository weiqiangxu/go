// 01-10自动推导类型和赋值的区别
package main

import "fmt"

func main() {
	// a = 10//对于未定义的变量不能直接赋值-除非自动推导赋值
	a := 10
	fmt.Println("a=", a)
	//自动推导，只能使用一次，用于初始化那次-同一个变量名只能使用1次，但是赋值可以N次，这就是区别
	b := 10
	// b := 30 //前面已经有变量b不能再新建一个变量b
}
