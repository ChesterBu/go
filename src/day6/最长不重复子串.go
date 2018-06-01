package main

import "fmt"

func lengthOfNonRepeatSubStr(str string)int{
	lastOccurred := make(map[byte]int)
	start :=0
	maxLen :=0
	for i,ch := range []byte(str) {
		lastI ,ok :=lastOccurred[ch]
		if ok && lastI >=start{
			start = lastI+1
		}
		if i-start+1>maxLen{
			maxLen = i-start+1
		}
		lastOccurred[ch]=i
	}
	return maxLen
}



func main() {
	fmt.Println(lengthOfNonRepeatSubStr("abcda"))
	fmt.Println(lengthOfNonRepeatSubStr("aaa"))

}
