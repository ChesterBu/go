package main

import (
	"fmt"
	"strings"
)

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
			res++
		}
	}
	fmt.Println(res)
}

func bigNumber(str1, str2 string) (result string) {

	if len(str1) == 0 && len(str2) == 0 {
		result = "0"
		return
	}

	if len(str1) < len(str2) {
		str1 = strings.Repeat("0", len(str2)-len(str1)) + str1
	} else if len(str1) > len(str2) {
		str2 = strings.Repeat("0", len(str1)-len(str2)) + str2
	}

	var index = len(str1) - 1
	var left int
	for index >= 0 {
		c1 := str1[index] - '0'
		c2 := str2[index] - '0'
		sum := int(c1) + int(c2) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0'
		result = fmt.Sprintf("%c%s", c3, result)
		index--
	}
	if left == 1 {
		result = fmt.Sprintf("1%s", result)
	}
	return
}

func main() {
	multiplication()
	isPrefectNumber(6)
	fmt.Println(isPalindrome("aba"))
	countNumber("ABC abc")
	fmt.Println(bigNumber("1623", "456"))

}
