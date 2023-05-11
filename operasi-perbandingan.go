package main

import "fmt"

func main() {
	// operator perbandingan
	// >
	// <
	// >=
	// <=
	// ==
	// !=

	var title1 = "judul"
	var title2 = "judul"

	var hasil bool = title1 == title2
	fmt.Println(hasil)

	var value1 = 12
	var value2 = 15

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 >= value2)
	fmt.Println(value1 <= value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
