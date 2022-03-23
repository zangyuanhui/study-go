package main

import "fmt"

var x, y int

// 这种因式分解关键字的写法一般用于声明全局变量
var (
	a int
	b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

//g, h := 123, "hello"

func main() {
	//这种不带声明格式的只能在函数体中出现
	g, h := 123, "hello"
	println(x, y, a, b, c, d, e, f, g, h)

	{
		// iota	特殊常量 可理解为 const 语句块中的行索引
		const (
			a = iota //0
			b        //1
			c        //2
			d = "ha" //独立值，iota += 1
			e        //"ha"   iota += 1
			f = 100  //iota +=1
			g        //100  iota +=1
			h = iota //7,恢复计数
			i        //8
		)
		fmt.Println(a, b, c, d, e, f, g, h, i)
	}
}
