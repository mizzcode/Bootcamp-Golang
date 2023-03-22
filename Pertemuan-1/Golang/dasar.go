package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	// tipe data int = bilangan bulat
	fmt.Println(2003)

	// tipe data float = bilangan decimal
	fmt.Println(9.9)

	// tipe data string = bilangan string
	// kumpulan karakter atau text
	fmt.Println("Hello World")

    // tipe data bool = bilangan boolean
	// yaitu type data untuk benar atau salah
    fmt.Println(true)
    fmt.Println(false)

    // tipe data nil = bilangan nil
    fmt.Println(nil)

    // tipe data map = bilangan map
    m := make(map[string]int)
    m["a"] = 1
    m["b"] = 2
    m["c"] = 3
    fmt.Println(m)

    // tipe data slice = bilangan slice
    slice := []int{1, 2, 3}
    fmt.Println(slice)

    // tipe data array = bilangan array
    array := [3]int{1, 2, 3}
	fmt.Println(array)

	// variable 
	// wadah atau tempat untuk menyimpan value
	// menggunakan keyword var dan type data
	var name string
    name = "Mizz Kun"
    fmt.Println("Nama:", name)

	// tanpa type data
	var bootcamp = "Genbox Academy"
	fmt.Println(bootcamp)

	// tanpa keyword var dan tanpa menggunakan type data
	age := 19
	fmt.Println("Umur:", age)

	// const
	// sebuah variable yang tidak bisa di ubah data nya
	const (
        pi = 3.14
        e  = 2.7
    )
    fmt.Println(pi)
    fmt.Println(e)

	var (
		firstName = "Jani"
		lastName = "Chan"
	)
	fmt.Println(firstName, lastName)

	// Operasi Aritmatika
	// + - * / %
	var (
		a = 1
        b = 2
        c = 3
	)
	hasil := a * b * c
	fmt.Println(hasil)

	// Perbandingan
	// < > <= >= != ==

	fmt.Println(10 > 5)

	// Penugasan
	// Assignment
	var tes int = 10

	tes = tes + 10
	fmt.Println(tes)

	tes += 10
	fmt.Println(tes)


	// unary
	i := 12
	i++ // increment i = i + 1
	i-- // decrement i = i - 1

	fmt.Println(i > 12 && 1 == 1 || !false)

}