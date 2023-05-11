package main

import "fmt"

// disini function sebagai value dari variabel

func getPeople(name string) string {
	return "where are you ? " + name
}

func main() {
	getpeople := getPeople
	fmt.Println(getpeople("inaiyah"))
}
