package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// parameter dikasih operator * (dalam kasus ini untuk menegaskan pointer menuju ke struct Address)
func ChangeCountry(address *Address) {
	address.Country = "Indonesia"
}

func main() {

	var alamat = Address{
		City:     "Kepanjen",
		Province: "Jawa Timur",
		Country:  "",
	}

	// ChangeCountry(&alamat)
	// atau
	alamatPointer := &alamat
	ChangeCountry(alamatPointer)

	fmt.Println(alamat)

}
