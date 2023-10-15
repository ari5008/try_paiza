package skilcheck

// package main

// import "fmt"

// func main() {

// 	var N, M int
// 	fmt.Scan(&N, &M)

// 	route_map := make([][]int, N)
// 	for i := 0; i < N; i++ {
// 		route_map[i] = make([]int, 0, M)
// 		for j := 0; j < M; j++ {
// 			var price int
// 			fmt.Scan(&price)

// 			route_map[i] = append(route_map[i], price)
// 		}
// 	}

// 	var X int
// 	fmt.Scan(&X)
// 	var sum int
// 	var next_price int
// 	var previous_price int
// 	var previous_route_follow int
// 	var previous_station_follow int

// 	for i := 0; i < X; i++ {
// 		var route_follow, station_follow int
// 		fmt.Scan(&route_follow, &station_follow)
// 		next_price = route_map[route_follow-1][station_follow-1]
// 		if previous_route_follow == 0 || previous_station_follow == 0 {
// 			previous_price = 0
// 		} else if route_follow != previous_route_follow {
// 			previous_price = route_map[route_follow-1][previous_station_follow-1]
// 		} else {
// 			previous_price = route_map[previous_route_follow-1][previous_station_follow-1]
// 		}
// 		spare := swapSign(next_price, previous_price)
// 		sum += spare
// 		previous_route_follow = route_follow
// 		previous_station_follow = station_follow
// 	}

// 	fmt.Println(sum)
// }

// func swapSign(next_price int, previous_price int) int {
// 	price := next_price - previous_price
// 	if price < 0 {
// 		price = -price
// 	}
// 	return price
// }
