package try

// package main

// import (
// 	"fmt"
// )

// type Customer struct {
// 	Sum      int
// 	Alcohol  bool
// 	Children bool
// }

// func NewCustomer(children bool) *Customer {
// 	return &Customer{
// 		Sum:      0,
// 		Alcohol:  false,
// 		Children: children,
// 	}
// }

// func (c *Customer) GetSum() int {
// 	return c.Sum
// }

// func (c *Customer) TakeSoftDrink(price int) {
// 	c.Sum += price
// }

// func (c *Customer) TakeFood(price int) {
// 	if c.Alcohol {
// 		c.Sum += price - 200
// 	} else {
// 		c.Sum += price
// 	}
// }

// func (c *Customer) TakeAlcohol(price int) {
// 	if !c.Children {
// 		c.Sum += price
// 		c.Alcohol = true
// 	}
// }

// func main() {
// 	var n, k int
// 	fmt.Scan(&n, &k)

// 	customers := make([]*Customer, n)
// 	for i := 0; i < n; i++ {
// 		var age int
// 		fmt.Scan(&age)
// 		if age >= 20 {
// 			customers[i] = NewCustomer(false)
// 		} else {
// 			customers[i] = NewCustomer(true)
// 		}
// 	}

// 	for i := 0; i < k; i++ {
// 		var index, price int
// 		var order string
// 		fmt.Scan(&index, &order)

// 		index--

// 		if order == "0" {
// 			customers[index].TakeAlcohol(500)
// 		} else {
// 			fmt.Scan(&price)
// 			if order == "food" {
// 				customers[index].TakeFood(price)
// 			} else if order == "softdrink" {
// 				customers[index].TakeSoftDrink(price)
// 			} else {
// 				customers[index].TakeAlcohol(price)
// 			}
// 		}
// 	}

// 	for _, customer := range customers {
// 		fmt.Println(customer.GetSum())
// 	}
// }

