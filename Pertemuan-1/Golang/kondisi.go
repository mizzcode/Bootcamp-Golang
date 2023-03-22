package main

import "fmt"

func main() {
	name := "Mizz"
	// 1. if
	// tidak menggunakan tanda kurung, jika kondisi pada if true maka blok code dijalankan
	if name == "Mizz" {
		fmt.Println("Hello", name)
	} else if name == "Jani" {
		fmt.Println("Hello", name)
	} else if name == "Jul" {
		fmt.Println("Hello", name)
	} else {
		fmt.Println("Kamu siapa")
	}

	// length := len(name)

	// if length > 5 {
	// 	fmt.Println("Nama Terlalu Panjang")
	// } else {
	// 	fmt.Println("Berhasil Daftar")
	// }

	// shorthand
	if length := len(name); length > 5 {
		fmt.Println("Nama Terlalu Panjang")
	} else {
		fmt.Println("Berhasil Daftar")
	}

	// shorthand lagi
	if nilai := 70; nilai > 80 {
		fmt.Println("Kamu Lulus")
	}  else {
		fmt.Println("Kamu Salah Jurusan")
	}

	// 2. Switch
	nama := "No Name"

	switch nama { 
		case "Mizz": 
			fmt.Println("Hello", nama)
		case "Jani":
			fmt.Println("Hello", nama)
        case "Jul":
			fmt.Println("Hello", nama)
		default:
			fmt.Println("Kamu siapa")
	}

	var names string
	fmt.Println("Masukkan Nama:")
	fmt.Scanln(&names)

	fmt.Println("Perkenalkan Nama Saya:", names)
}