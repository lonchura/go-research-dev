package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

// 指针引用替换
/* 定义交换值函数 */
//func swap(x *int, y *int) {
//	var temp int
//	temp = *x    /* 保持 x 地址上的值 */
//	*x = *y      /* 将 y 值赋给 x */
//	*y = temp    /* 将 temp 值赋给 y */
//}

func main() {
	a, b := swap("Mahesh", "Kumar")
	fmt.Println(a, b)
}