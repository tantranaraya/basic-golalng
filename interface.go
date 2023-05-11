package main

import "fmt"

/**
- Interface adalah tipe data abstrak dan tidak memiliki implementasi langsung
- Sebuah interface berisikan definisi2 method
- biasanya interface digunakan sebagai kontrak (contract)
*/

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func sayAnything(hasName HasName) {
	fmt.Println("I wanna say to you", hasName.GetName())
}

func main() {
	var data Person
	data.Name = "Hae-Joon"

	sayAnything(data)

	var dataAnimal Animal
	dataAnimal.Name = "Picky Blinders Cat"

	sayAnything(dataAnimal)
}
