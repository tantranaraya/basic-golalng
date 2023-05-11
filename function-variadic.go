package main

import "fmt"

// variable arguments (varargs) artinya datanya bisa menerima lebih dari satu input
// ...int, ...string artinya bisa digunakan banyak data

func sumAll(numbers ...int) int {
	total := 0

	// memakai _, agar tidak menampilkan index pada perulangan for range
	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {

	total := sumAll(10, 10, 10, 10, 10, 10)
	fmt.Println(total)

	// menggunakan slice pada variadic function
	slice := []int{10, 20, 30, 50, 60}
	total = sumAll(slice...)
	fmt.Print("Hasil penggunaan slice ", total)
}
