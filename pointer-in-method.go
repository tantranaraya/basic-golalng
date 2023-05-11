package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr." + man.Name
}

func main() {

	boy := Man{Name: "Neo"}
	boy.Married()
	fmt.Println(boy.Name)
}
