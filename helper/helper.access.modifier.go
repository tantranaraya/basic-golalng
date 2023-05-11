package helper

import "fmt"

/**
Nama function atau variabel jika berawalan huruf besar maka bisa diakses langsung dari file lain (bersifat global)
Nama function atau variabel jika berawalan huruf kecil maka hanya bisa diakses di dalam file itu sendiri (bersifat private)
*/

var version = "1.0.0 beta"
var Application = "Belajar Golang"

func sayLearning(name string) string {

	return "Yuk learning" + name + version
}

func SayPractice(name string) string {
	fmt.Println(sayLearning("Akses dari dalam file"))
	return "Yuk Practice" + name

}
