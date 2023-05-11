package main

import "fmt"

/**
nil hanya bisa digunakan pada tipe data interface, function, map, slice, pointer, channel
*/
func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {

	var data map[string]string = NewMap("niu")

	if data == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(data["name"])
	}
}
