package main

import (
	"errors"
	"fmt"
)

// interface yang digunakan sebagai kontrak untuk membuat error, nama interfacenya adalah error

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian tidak boleh dengan nol")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
	// var errorExample error = errors.New("Waduh error !")
	// fmt.Println(errorExample.Error())

	hasil, err := Pembagian(100, 0)

	if err == nil {
		fmt.Println("Hasilnya", hasil)
	} else {
		fmt.Println("Hasilnya", hasil)
		fmt.Println("Error kan ! sudah dibilang !", err.Error())
	}
}
