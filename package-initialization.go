package main

import (
	"basic-golang/database"
	"fmt"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
