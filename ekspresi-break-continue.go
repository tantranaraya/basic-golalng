package main

import "fmt"

func main() {

	// break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println("perulangan break", i)
	}

	// continue
	for i := 0; i < 10; i++ {

		// menampilkan angka ganjil
		if i%2 == 0 {
			continue
		}

		fmt.Println("perulangan ganjil", i)
	}

	for i := 0; i < 10; i++ {

		// menampilkan setelah 5
		if i <= 5 {
			continue
		}

		fmt.Println("perulangan setelah 5", i)
	}
}
