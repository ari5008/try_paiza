package levelup

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var a, b, q int
// 	fmt.Scan(&a, &b)
// 	fmt.Scan(&q)

// 	table := make([][]int, 100)
// 	for i := 0; i < 100; i++ {
// 		table[i] = []int{}
// 	}

// 	for i := 0; i < q; i++ {
// 		var op, x int
// 		fmt.Scan(&op, &x)

// 		hash := (a*x + b) % 100

// 		if op == 1 {
// 			table[hash] = append(table[hash], x)
// 		} else if op == 2 {
// 			found := false
// 			for _, value := range table[hash] {
// 				if value == x {
// 					found = true
// 					break
// 				}
// 			}
// 			if found {
// 				fmt.Println("Yes")
// 			} else {
// 				fmt.Println("No")
// 			}
// 		}
// 	}

// 	for i := 0; i < 100; i++ {
// 		for j, value := range table[i] {
// 			if j > 0 {
// 				fmt.Print(" ")
// 			}
// 			fmt.Print(value)
// 		}
// 		fmt.Println()
// 	}
// }
