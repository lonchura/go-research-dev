package main

import "fmt"

// 阶乘
func Factorial(x int) (result int) {
	if x == 0 {
		result = 1;
	} else {
		result = x * Factorial(x - 1);
	}
	return;
}

// 斐波那契数列
func fibonacci(n int) (res int) {
	if n <= 2 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

func main() {
	var i int = 15
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(i))

	// fibonacci
	result := 0
	for i := 1; i <= 10; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
}