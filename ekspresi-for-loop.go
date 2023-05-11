package main

import "fmt"

func main() {

	// bentuk pertama init statement
	counter := 1

	for counter <= 10 {
		fmt.Println("perulangan ke", counter)
		counter++
	}

	// bentuk kedua post statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("perulangan bentuk kedua, ke", counter)
	}

	// -----------------------------------------------------------------------------

	// for range

	fmt.Println("for range :")

	slice := []string{
		"noe",
		"nei",
		"neo",
	}

	// for range bentuk pertama

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// for range bentuk kedua

	for index, value := range slice {
		fmt.Println("index", index, "=", value)
	}

	// _, jika menampilkan data saja tanpa menampilkan index,atau sebaliknya
	for _, value := range slice {
		fmt.Println("data saja", value)
	}

	// ---------------------------------------------------------------------------------------------

	// for range map

	person := make(map[string]string)

	person["name"] = "noi"
	person["title"] = "tukang tukangan"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
