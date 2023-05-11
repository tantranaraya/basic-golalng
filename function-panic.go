package main

import "fmt"

// Panic adalah function yang bisa dipakai untuk menghentikan program
// Panic function biasanya dipanggil ketika terjadi error pada saat program kita berjalan
// Saat panic function dipanggil, program akan terhenti, tetapi defer function akan tetap dieksekusi

func endApp() {
	fmt.Println("Aplikasi selese")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi kacau")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
}
