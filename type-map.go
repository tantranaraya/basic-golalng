package main

import "fmt"

func main() {

	// person := map[string]string{
	// 		"name":    "My Name is",
	// 		"address": "My Address is",
	// }

	// atau

	// person := make(map[string]string)
	// person["name"] = "my name"
	// person["address"] = "my address"

	// bentuk pertama

	person := map[string]string{
		"name":    "My Name is",
		"address": "My Address is",
	}

	// penambahan atau mengubah data key dan value diluar deklarasi
	person["title"] = "Tukang Koding"

	// delete(map, "key")

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])
	// mendapatkan jumlah data di map -> len(map)
	fmt.Println(len(person))

	delete(person, "title")

	// data setelah dihapus
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])
	fmt.Println(len(person))

	// bentuk kedua
	user := make(map[string]string)

	user["name"] = "we name are"
	user["address"] = "we address are"

	fmt.Println(user["name"])
	fmt.Println(user["address"])
	fmt.Println(len(person))

}
