package main

import "fmt"

func main() {
	// tipe noSim adalah alias dari tipe data string
	type noSim string

	var sim noSim = "90909089098"

	fmt.Println(sim)
	fmt.Println(noSim("545454545454"))

}
