package levelup

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var n int
// 	fmt.Scan(&n)

// 	arrs := make([][]int, 10)

// 	for i := 0; i <= n-1; i++ {
// 		var num int
// 		fmt.Scan(&num)

// 		Index := hash(num)
// 		arrs[Index] = append(arrs[Index], num)
// 	}

// 	fmt.Println(arrs)
// 	for _, number := range arrs {
// 		if number != nil {
// 			for _, value := range number {
// 				fmt.Print(value)
// 				fmt.Print(" ")
// 			}
// 			fmt.Println()
// 		}
// 	}
// }

// func hash(num int) int {
// 	return num % 10
// }
