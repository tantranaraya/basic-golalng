package main

import "fmt"

// function recover berfungsi untuk menangkap pesan panic

func endApp() {
	// ini adalah defer function
	// recover yang benar diletakan di defer function

	message := recover()
	if message != nil {
		fmt.Println("terjadi error dengan pesan", message)
	}
	fmt.Println("aplikasi selese")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi kacau")
	}
	fmt.Println("aplikasi berjalan")
}

func main() {
	runApp(true)
}
