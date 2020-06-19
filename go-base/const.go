package main

import "unsafe"

// 函数必须是内置函数，否则编译不过！！！
const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
)

func main(){
	println(a, b, c)
}