package main

import "fmt"

func main() {
	fmt.Println("Learn Array")

	// Array adalah sebuah tipe data yang menampung banyak value dengan tipe data apapun misalnya object, string, int, bool, dll
	// menggunakan index. Dimulai dari 0

	// var color [3]string 

	// color[0] = "green"
	// color[1] = "black"
	// color[2] = "white"

	// // akses data array pakai index
	// fmt.Println(color)
	// fmt.Println(color[0])
	// fmt.Println(color[1])
	// fmt.Println(color[2])

	// fmt.Println("Panjang Array", len(color))

	// // ubah data array menggunakan index
	// color[0] = "red"
	// color[1] = "green"
	// color[2] = "blue"
	
    // fmt.Println(color)

	// clubs := [...]string {
	// 	"Emyu",
	// 	"Barcelona",
	// 	"Real Madrid",
	// 	"PSG",
	// 	"Indonesia",
	// }

	// fmt.Println(clubs)
	// fmt.Println(len(clubs))
	
	// // menggunakan index
	// for i, club := range clubs { 
	// 	fmt.Println(i+1, club)
	// }

	// fmt.Println("===========")

	// // tanpa index
	// for _, club := range clubs { 
	// 	fmt.Println(club)
	// }

	// Slice
	// Slice adalah potongan dari array
	// array[low:high] = akses array dari index low dan berakhir sebelum index high
	// array[low:] = akses array dari index low dan berakhir sampai akhir array
	// array[:high] = akses array dari index 0 dan berakhir sampai sebelum index high
    // array[:] = akses array dari index 0 dan berakhir sampai akhir array

	// slice1 := clubs[1:3]
	// fmt.Println(slice1)
	// fmt.Println("Length Slice:", len(slice1))

	// slice2 := clubs[:5]
	// fmt.Println(slice2)
	// fmt.Println("Length Slice:", len(slice2))

	// slice3 := clubs[1:]
	// fmt.Println(slice3)
	// fmt.Println("Length Slice:", len(slice3))

	// slice4 := clubs[:]
	// fmt.Println(slice4)
	// fmt.Println("Length Slice:", len(slice4))

	// length = 2 || capacity = 5
	newSlice := make([]string, 2, 5) 

	newSlice[0] = "MJ"
	newSlice[1] = "Peter Parker"

	fmt.Println(newSlice)
	
	// // jika array length array tidak cukup, maka append akan membuat data array baru
	// slice5 := append(newSlice, "Mizz")
	// fmt.Println(slice5)

	// bisa juga tanpa di masukan ke var
	newSlice = append(newSlice, "Tes")
	fmt.Println(newSlice)

	// perbedaan slice dan array
	// slice
	// jika data array tidak memenuhi length nya maka akan di isi 0
	slice := make([]int, 3, 5)
	slice[0] = 1
	slice[1] = 2

	// array
	array := [...]int {
		1,
		2,
	}

	fmt.Println(slice)
	fmt.Println(array)

}