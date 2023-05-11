package main

import "fmt"

/**
Secara default di golang semua variabel itu dipassing by value, bukan by reference
artinya jika kita mengirim sebuah variabel ke dalam function, method atau variabel lain,
sebenarnya yang dikirimkan adalah duplikasi value nya
*/

/**
Pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama, tanpa menduplikasi data yang ada
Sederhananya dengan kemampuan pointer kita bisa membuat pass by reference.

Kalau pass by value data diduplikasi ke lokasi memory yang berbeda
Kalau pass by reference datanya tetap di lokasi memory yang sama
*/

/**
- Operator "&"
  Untuk membuat sebuah variabel dengan nilai pointer ke variabel yang lain, kita bisa menggunakan operator "&"
  diikuti dengan nama variabelnya. Ini disebut sebagai variabel pointer.
  Saat kita mengubah variabel pointer, maka yang berubah hanya variabel tersebut, variabel lain yang mengacu ke data
  tersebut tidak akan berubah.
- Operator "*"
  Jika kita ingin mengubah seluruh variabel yang mengacu ke data tersebut, kita bisa menggunakan operator "*"
- Function "new"
  Digunakan untuk membuat pointer, namun hanya mengembalikan pointer ke data kosong, tidak mengembalikan ke data awal
*/

/**
Sebaiknya menggunakan pointer, karena jika tidak akan terjadi pembengkakan data di memory akibat duplikasi data
yang berulang kali yang berakibat pemborosan memory
*/

type Address struct {
	City, Province, Country string
}

func main() {

	// Bentuk pertama
	// var address1 Address = Address{"Kepanjen", "Jawa Timur", "Indonesia"}
	// var address2 *Address = &address1
	// var address3 *Address = &address2

	// Bentuk kedua
	address1 := Address{City: "Kepanjen", Province: "Jawa Timur", Country: "Indonesia"}
	address2 := &address1 /** address2 mem-pointer ke address1, artinya address2 satu lokasi memory dengan address1*/
	address3 := &address1

	address2.City = "Gondang Legi"                                                  /** variabel pointer hanya merubah satu field di address1*/
	*address2 = Address{City: "Turen", Province: "Jawa Barat", Country: "Columbia"} /** operator "*" untuk bisa mengubah seluruh variabel address1 / Address */

	fmt.Println("Ini address1", address1)
	fmt.Println("Ini address2", address2)
	fmt.Println("Ini address3", address3) /** data address3 juga ikut berubah sesuai perubahan di address2 karena sama2 mereference ke address1 */

	/** Penggunaan function new() untuk mengosongkan / membuat data baru di pointer pada lokasi memory yang sama */
	// var alamat1 *Address = new(Address)
	// alamat1.City = "Tumpang"

	alamat1 := new(Address)
	alamat2 := alamat1

	alamat2.City = "Ngandat"
	alamat2.Province = "Jawa Tengah"

	fmt.Println("Ini alamat 1", alamat1)
	fmt.Println("Ini alamat 2", alamat2)

}
