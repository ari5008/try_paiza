package levelup

// package main

// import (
// 	"crypto/sha256"
// 	"encoding/binary"
// 	"fmt"
// )

// func main() {
// 	var (
// 		n int
// 	)
// 	fmt.Scan(&n)

// 	for i := 0; i <= n-1; i++ {
// 		var str string
// 		fmt.Scan(&str)


// 		fmt.Println(hash(str))
// 	}
// }

// func hash(str string) int {
// 	hash := sha256.New()
// 	hash.Write([]byte(str))
// 	hashBytes := hash.Sum(nil)
// 	hashValue := binary.BigEndian.Uint32(hashBytes[:4])

// 	return int((hashValue % 10) + 1)
// }