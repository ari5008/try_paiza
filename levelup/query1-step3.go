package levelup

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var (
// 		N int
// 		Q int
// 	)
// 	fmt.Scan(&N, &Q)

// 	var nums []int
// 	for i := 0; i <= N-1; i++ {
// 		var num int
// 		fmt.Scan(&num)
// 		nums = append(nums, num)
// 	}

// 	flag := false
// 	newNums := QuickSort(nums)
// 	for i := 0; i <= Q-1; i++ {
// 		var foundNum int
// 		fmt.Scan(&foundNum)

// 		flag = binarySearch(newNums, foundNum)

// 		if flag {
// 			fmt.Println("YES")
// 		} else {
// 			fmt.Println("NO")
// 		}
// 	}

// }

// func QuickSort(numbers []int) []int {
// 	quickSort(numbers, 0, len(numbers)-1)
// 	return numbers
// }

// func quickSort(numbers []int, low int, high int) {
// 	if low < high {
// 		Index := partition(numbers, low, high)
// 		quickSort(numbers, low, Index-1)
// 		quickSort(numbers, Index+1, high)
// 	}
// }

// func partition(numbers []int, low int, high int) int {
// 	pivot := numbers[high]
// 	i := low - 1

// 	for j := low; j < high; j++ {
// 		if numbers[j] < pivot {
// 			i++
// 			numbers[i], numbers[j] = numbers[j], numbers[i]
// 		}
// 	}
// 	numbers[high], numbers[i+1] = numbers[i+1], numbers[high]
// 	return i + 1
// }

// func binarySearch(numbers []int, value int) bool {
// 	left, right := 0, len(numbers)-1
// 	for left <= right {
// 		mid := (left + right) / 2
// 		if numbers[mid] == value {
// 			return true
// 		} else if numbers[mid] < value {
// 			left = mid + 1
// 		} else {
// 			right = mid - 1
// 		}
// 	}
// 	return false
// }
