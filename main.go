package main

import (
	"fmt"
)

func SortSlice(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr)-i; j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}

func IncrementOdd(slice []int) {
	for i := 1; i < len(slice); i += 2 {
		slice[i]++
	}
}

func PrintSlice(slice []int) {
	fmt.Println(slice)
}

func ReverseSlice(slice []int) {
	for i, j := 0, len(slice)-1; i < len(slice)/2; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func appendFunc(dst func([]int), src ...func([]int)) func([]int) {
	return func(a []int) {
		dst(a)
		for _, functions := range src {
			functions(a)
		}
	}
}

func main() {
	arr := []int{5, 6, 7, 4, 1, 0}
	// 0 1 4 5 6 7
	// 0 2 4 6 6 8
	// 8 6 6 4 2 0
	newFUnc := appendFunc(SortSlice, IncrementOdd, ReverseSlice, PrintSlice)
	newFUnc(arr)
}
