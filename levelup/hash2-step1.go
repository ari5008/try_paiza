package levelup

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var n int
// 	fmt.Scan(&n)

// 	arrs := make([]int, 10)
// 	for i := range arrs {
// 		arrs[i] = -1
// 	}

// 	for i := 0; i <= n-1; i++ {
// 		var num int
// 		fmt.Scan(&num)

// 		Index := hash(num)
// 		for arrs[Index] != -1 {
// 			Index = hash(Index + 1)
// 		}

// 		arrs[Index] = num

// 	}
// 	for _, number := range arrs {
// 		fmt.Println(number)
// 	}
// }

// func hash(num int) int {
// 	return num % 10
// }
