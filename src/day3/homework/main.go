package main

import "fmt"

func multiplication() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d ", i, j, i*j)
		}
		fmt.Printf("\n")
	}
}

func isPrefectNumber(n int) bool {
	res := 0
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			res += i
		}
	}
	if res == n {
		fmt.Println(res, n)
		return true
	} else {
		return false
	}
}

func isPalindrome(str string) bool {

	if len(str) > 2 {
		return str[len(str)-1] == str[0] && isPalindrome(str[1:len(str)-2])
	} else if len(str) == 2 {
		return str[len(str)-1] == str[0]
	}
	return true

}

func countNumber(str string) {
	res := 0
	for _, v := range str {
		switch {
		case v <= 90 && v >= 65 || v >= 97 && v <= 122:
			res += 1
		}
	}
	fmt.Println(res)
}

func bigNumber(a, b string) {

}

func main() {
	multiplication()
	isPrefectNumber(6)
	fmt.Println(isPalindrome("aba"))
	countNumber("ABC abc")
}
