package main

import "fmt"

func main() {

	name := "ayana moon"

	switch name {
	case "ayana moon":
		fmt.Println("halo", name)
		fmt.Println("kamu memang benar", name)
	case "nei":
		fmt.Println(name, "kamu bukan yang kukenal")
		fmt.Println("silahkan daftar dulu")
	default:
		fmt.Println("kamu siapa lagi ?", name)
	}

	// short statemant switch
	// memakai boolean
	// bentuk pertama
	var length = len(name)

	switch length > 5 {
	case true:
		fmt.Println("karakter terlalu panjang (1)")
	case false:
		fmt.Println("karakter memenuhi (1)")
	}

	// bentuk kedua

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("karakter terlalu panjang (2)")
	case false:
		fmt.Println("karakter memenuhi (2)")
	}

	// switch tanpa kondisi

	lengthval := len(name)

	switch {
	case lengthval > 5:
		fmt.Println("karakter terlalu panjang (tanpa kondisi)")
	case lengthval > 3:
		fmt.Println("karakter lumayan panjang (tanpa kondisi)")
	default:
		fmt.Println("karakter sesuai")
	}

}
