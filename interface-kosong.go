package main

import "fmt"

func Whats(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Whats ?"
	}
}

func main() {

	var data interface{} = Whats(1)
	fmt.Println(data)

	empty := Whats(1)
	fmt.Println(empty)
}
