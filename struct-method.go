package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// Struct function / method
func (customer Customer) saySomething() {
	fmt.Println("I wanna say something to", customer.Name)
}

func main() {
	var data Customer
	data.Name = "Neena"
	data.Address = "Vrindavan"
	data.Age = 34

	data.saySomething()
}
