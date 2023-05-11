package main

import "fmt"

// function()(named return values and typedata)

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Noe"
	middleName = "Neil"
	lastName = "Nea"

	// return firstName, middleName, lastName
	// atau
	// return

	return
}

func main() {

	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName, middleName, lastName)

}
