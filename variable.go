package main

import "fmt"

func main() {
	var name string

	name = "Kalimat string"
	fmt.Println(name)

	name = "Kalimat string yang lain"
	fmt.Println(name)

	var people = "All people"
	fmt.Println(people)

	var numint int8
	numint = 107
	fmt.Println(numint)

	var numfloat = 1.55
	fmt.Println(numfloat)

	// assignment dan declaration
	kampung := "Karangploso"
	fmt.Println(kampung, "otomatis string")

	// Multiple variable
	var (
		title        string
		namaDepan    string
		namaBelakang string
	)

	title = "multiple variable"
	namaDepan = "Nama Depannya"
	namaBelakang = "Nama Belakangnya"

	fmt.Println(title)
	fmt.Println(namaDepan)
	fmt.Println(namaBelakang)

}
