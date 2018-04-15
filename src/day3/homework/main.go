package main

import (
	"bytes"
	"fmt"
	"strings"
	"strconv"
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

func bigNumber(a, b string)string {
	if len(a) < len(b) {
		a = strings.Repeat("0", len(b)-len(a)) + a
	} else {
		b = strings.Repeat("0", len(a)-len(b)) + b
	}
	strLen := len(a)
	nums := make([]uint8, strLen)
	addOne := false
	for i := 0; i < strLen; i++ {
		numA := a[strLen-i-1]-'0'
		numB := b[strLen-i-1]-'0'
		sum:= numA + numB
		if addOne {
			sum++
		}
		if sum >9{
			sum -= 10
			addOne = true 
		} else{
			addOne = false
		}
		nums[i] = sum
	}
	result := convertToString(nums)  
    //进位，最前面补1  
    if addOne {  
        result = "1" + result  
    }  
  
    return result 
}

func convertToString(nums []uint8) string{
	var b bytes.Buffer
	for i := len(nums) - 1; i >= 0; i-- {  
        b.WriteString(strconv.Itoa(int(nums[i])))
    }  
    return b.String() 
}

func main() {
	multiplication()
	isPrefectNumber(6)
	fmt.Println(isPalindrome("aba"))
	countNumber("ABC abc")
}
