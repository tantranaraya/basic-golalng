package main

import "fmt"

// Closure adalah kemampuan sebuah function berinteraksi dengan data2 disekitarnya dalam scope yang sama
// scope adalah lingkup kerja sebuah blok function atau variabel

func main() {
	name := "nora"
	counter := 0
	increment := func() {
		// deklarasi ulang variabel name (:= bukan =) agar tidak meng-overwrite data name := nora dari scope di atasnya
		name := "neira"
		fmt.Println("increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()
	fmt.Println(counter)
	// variabel name dari scope di atas (name := "nora")
	fmt.Println(name)

}
