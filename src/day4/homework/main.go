package main

import "fmt"

func bubbleSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a
}

func insertionSort(a []int) []int {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
	return a

}

func quickSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	pivotIndex := len(a) / 2
	pivot := a[pivotIndex:pivotIndex+1][0]
	a = append(a[:pivotIndex], a[pivotIndex+1:]...)
	var left, right []int
	for _, value := range a {
		if value < pivot {
			left = append(left, value)
		} else {
			right = append(right, value)
		}
	}
	return append(quickSort(left), append([]int{pivot}, quickSort(right)...)...)

}

func selectionSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		min := i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		if min != i {
			a[i], a[min] = a[min], a[i]
		}
	}
	return a
}

func main() {
	a := [...]int{90, 12, 284, 920, 573, 329, 243, 0, -1}
	//fmt.Println(bubbleSort(a[:]))
	//fmt.Println(insertionSort(a[:]))
	//fmt.Println(selectionSort(a[:]))
	fmt.Println(quickSort(a[:]))

}
