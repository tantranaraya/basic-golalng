package main

import "fmt"

/**
- Struct adalah template data atau prototipe data yang digunakan untuk menggabungkan nol atau lebih tipe data
  lainnya dalam satu kesatuan
- Struct biasanya representasi data dalam aplikasi yang kita buat
- Data di struct disimpan dalam field
- Struct seperti class / abstract di OOP atau juga constructor,  di golang yang prosedural berbasis funtion tidak mengenal class
- Sederhananya Struct adalah kumpulan dari field

*/

type Customer struct {
	Name, Address string
	Age           int
}

func main() {

	// Bentuk pemanggilan data

	var data Customer
	data.Name = "Neena"
	data.Address = "Vrindavan"
	data.Age = 34

	fmt.Println(data)
	fmt.Println(data.Name)
	fmt.Println(data.Address)
	fmt.Println(data.Age)

	// Bentuk struct literal

	datas := Customer{
		Name:    "Naina",
		Address: "Aligarh",
		Age:     35,
	}

	fmt.Println(datas)
	fmt.Println(datas.Name)
	fmt.Println(datas.Address)
	fmt.Println(datas.Age)

}
