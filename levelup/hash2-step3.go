package levelup

// package main

// import (
//     "fmt"
// )

// func main() {
//     table := make([][]int, 100)
//     for i := 0; i < 100; i++ {
//         table[i] = []int{}
//     }

//     var a, b, q int
//     fmt.Scan(&a, &b)
//     fmt.Scan(&q)

//     for i := 0; i < q; i++ {
//         var op, x int
//         fmt.Scan(&op, &x)

//         hash := (a*x + b) % 100

//         switch op {
//         case 1:
//             table[hash] = append(table[hash], x)
//         case 2:
//             found := false
//             for _, value := range table[hash] {
//                 if value == x {
//                     found = true
//                     break
//                 }
//             }
//             if found {
//                 fmt.Println("Yes")
//             } else {
//                 fmt.Println("No")
//             }
//         }
//     }

//     for i := 0; i < 100; i++ {
//         for j := 0; j < len(table[i]); j++ {
//             if j > 0 {
//                 fmt.Print(" ")
//             }
//             fmt.Print(table[i][j])
//         }
//         fmt.Println()
//     }
// }
