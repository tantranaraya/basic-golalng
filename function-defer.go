package main

import "fmt"

// Defer function adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai dieksekusi
// Defer function akan selalu dieksekusi walaupun terjadi error di function yang dieksekusi

func logging() {
	fmt.Println("selesai menjalankan function")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("menjalankan aplikasi")
	result := 10 / value
	fmt.Println("Hasil ", result)
}

func main() {
	runApplication(0)
}
