package main

import (
	"fmt"
)
//判断是100-200内的素数
func isPrime() {
	for i := 100; i < 201; i++ {
		flag := true
		for j := 2; j < i/2; j++ {
			if i%j == 0 {
				flag = false
			}
		}
		if flag {
			fmt.Println(i)
		}
	}
}
//是否为水仙花数
func isNarcissistic(n int) bool {
	i, j, k := (n/100)%10, (n/10)%10, n%10
	sum := i*i*i + j*j*j + k*k*k
	return sum == n
}
//阶乘
func factorial(n int) int {
	if n > 1 {
		return n * factorial(n-1)
	} else {
		return 1
	}
}

func main() {
	isPrime()
	fmt.Println(isNarcissistic(153))
	fmt.Println(factorial(4))
}
