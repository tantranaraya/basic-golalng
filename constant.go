package main

import "fmt"

func main() {

	// Constant harus langsung ada nilai nya dan nilai nya tidak bisa dirubah
	const namaDepan string = "Nama Depannya"
	const namaBelakang = "Nama Belakang tanpa tipedata"
	const value = 100

	fmt.Println(namaDepan)
	fmt.Println(namaBelakang)
	fmt.Println(value)

	// Multiple constant

	const (
		alamat string = "alamat"
		kota          = "kota"
		nomor         = 454
	)

	fmt.Println(alamat)
	fmt.Println(kota)
	fmt.Println(nomor)

}
