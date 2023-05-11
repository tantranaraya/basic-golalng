package main

import "fmt"

func main() {

	var title [4]string

	title[0] = "String satu"
	title[1] = "String dua"
	title[2] = "String tiga"

	fmt.Println(title[0])
	fmt.Println(title[1])
	fmt.Println(title[2])

	var nilai = [4]int{
		99,
		90,
		22,
	}

	fmt.Println(nilai)
	fmt.Println(nilai[0])
	fmt.Println(nilai[1])
	fmt.Println(nilai[2])

	// len panjang array bukan jumlah data di dalam array
	fmt.Println("panjang title", len(title))
	fmt.Println("panjang nilai", len(nilai))

}
