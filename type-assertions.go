package main

import "fmt"

/**
Type assertions merupakan kemampuan merubah tipe data menjadi tipe data yang kita inginkan
*/

func random() interface{} {
	return "Whoaps"
}

func main() {
	// result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	// menimbulkan panic ketika tipe data yang tidak sesuai untuk dirubah
	// resultInt := result.(int)
	// fmt.Println(resultInt)

	// untuk mencegah panic maka sebaiknya memakai switch
	result := random()

	switch value := result.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("tipe data tidak diketahui / tidak diijinkan")
	}
}
