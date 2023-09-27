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
// 	for i := 0; i <= Q-1; i++ {
// 		var foundNum int
// 		fmt.Scan(&foundNum)

		

// 		flag = binarySearch(nums, foundNum)

// 		// for _, num := range nums {
// 		// 	if num == foundNum {
// 		// 		flag = true
// 		// 		break
// 		// 	} else {
// 		// 		flag = false
// 		// 	}
// 		// }
// 		if flag {
// 			fmt.Println("YES")
// 		} else {
// 			fmt.Println("NO")
// 		}
// 	}

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
