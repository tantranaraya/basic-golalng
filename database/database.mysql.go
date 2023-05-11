package database

import "fmt"

var connection string

// function default untuk menginisialisasi konfigurasi
// function init otomatis akan terkesekusi
func init() {

	fmt.Println("mencoba memanggil Init")
	connection = "MySql"
}

func GetDatabase() string {
	return connection
}
