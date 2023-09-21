package try

// package main

// import "fmt"

// type People struct {
// 	Number      int
// 	Old         int
// 	Children    bool
// 	alcoholFlag bool
// 	TotalPrice  int
// }

// func main() {
// 	var (
// 		N int
// 		K int
// 	)
// 	fmt.Scan(&N, &K)

// 	peoples := []People{}
// 	for i := 0; i <= N-1; i++ {
// 		var old int
// 		fmt.Scan(&old)
// 		if old >= 20 {
// 			people := People{
// 				Number:   i + 1,
// 				Old:      old,
// 				Children: false,
// 			}
// 			peoples = append(peoples, people)
// 		} else {
// 			people := People{
// 				Number:   i + 1,
// 				Old:      old,
// 				Children: true,
// 			}
// 			peoples = append(peoples, people)
// 		}
// 	}

// 	for i := 0; i <= K-1; i++ {
// 		var (
// 			num   int
// 			dish  string
// 			price int
// 		)
// 		fmt.Scan(&num, &dish, &price)

// 		if dish == "softdrink" {
// 			peoples[num-1].TotalPrice += price
// 		} else if dish == "food" {
// 			if peoples[num-1].alcoholFlag {
// 				peoples[num-1].TotalPrice += price - 200
// 			} else {
// 				peoples[num-1].TotalPrice += price
// 			}
// 		} else if dish == "alcohol" {
// 			if peoples[num-1].Children {
// 				continue
// 			} else {
// 				peoples[num-1].TotalPrice += price
// 				peoples[num-1].alcoholFlag = true
// 			}
// 		}
// 	}

// 	for _, people := range peoples {
// 		fmt.Println(people.TotalPrice)
// 	}

// }

