package main

import "fmt"

func main() {

	// name := "Noe"

	//  if
	var name = "Noe"

	if name == "Noe" {
		fmt.Println("halo", name)
	}

	// if else
	if name == "Nei" {
		fmt.Println("hoi", name)
	} else {
		fmt.Println("kamu siapa ?")
	}

	// if else if
	if name == "nei" {
		fmt.Println("hallo", name)
	} else if name == "Noe" {
		fmt.Println("halo", name)
	} else {
		fmt.Println("kamu siapa lagi ?")
	}

	// short statement if

	// bentuk pertama
	var length = len(name)

	if length < 5 {
		fmt.Println("karakter terlalu pendek (nih bentuk pertama)")
	}

	// bentuk kedua -> langsung pake declarative assignment
	if length := len(name); length < 5 {
		fmt.Println("karakter terlalu pendek (nih bentuk kedua)")
	}
}
