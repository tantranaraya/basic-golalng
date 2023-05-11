package main

import "fmt"

func main() {
	var absensi = 90
	var nilai = 87

	var lulusujian = nilai > 75
	var lulusabsensi = absensi > 80

	var lulus = lulusabsensi && lulusujian
	fmt.Println(lulus)
}
