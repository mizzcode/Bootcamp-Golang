package main

import "fmt"

// function adalah tempat menyimpan banyaknya program dan akan di eksekusi 
// ketika sebuah function di panggil

// jika sebuah function memiliki parameter maka ketika function di panggil
// harus kita isi parameter/argument nya sesuai tipe data nya
func sayHello(firstName string, lastName string) {
	fmt.Println("Halo!", firstName, lastName)
}

// function juga bisa mengembalikan nilai. ketika function bisa mengembalikan nilai
// bisa kita simpan di variabel
func getHello(fullname string) string { 
	if fullname == "" {
		return "Hello MasBro"
	} else {
		return "Hello " + fullname
	}
}

// multiple return value // function yang mengembalikan lebih dari 1 nilai
func tes(nama string) (string, string, string) { 
	return nama, "Kamu", "Sehat"
}

func main() {
	firstName := "Mizz"
	lastName := "Kun"

	sayHello(firstName, lastName)
	sayHello("Jani", "Chan")


	getHello := getHello("")

	fmt.Println(getHello)
}