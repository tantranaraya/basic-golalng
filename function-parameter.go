package main

import "fmt"

func saySomething(firstName string, lastName string) {

	fmt.Println("Kamu siapa ?", firstName, lastName)

}

func getAlways(number int) {
	for number := 0; number < 10; number++ {
		fmt.Println(number)
	}

}

func main() {

	firstName := "Noeya"
	saySomething(firstName, "Nei")
	saySomething("neil", "nial")

	getAlways(10)
}
