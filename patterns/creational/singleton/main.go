package main

import (
	"fmt"
)

func main() {
	db1 := GetDatabaseConnection()
	db2 := GetDatabaseConnection()

	fmt.Println(db1 == db2) // Output: true
}
