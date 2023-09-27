package try

// package main

// import (
// 	"fmt"

// )

// func main() {
// 	var (
// 		N int
// 		K int
// 		Q int
// 	)
// 	fmt.Scan(&N, &K, &Q)

// 	var nums []int
// 	for i := 0; i <= N - 1; i++ {
// 		var num int
// 		fmt.Scan(&num)
// 		nums = append(nums, num)
// 	}

// 	nums = append(nums[:K], nums[K - 1:]...)
// 	nums[K] = Q
// 	fmt.Println(nums)

// 	for _, number := range nums {
// 		fmt.Println(number)
// 	}
// }