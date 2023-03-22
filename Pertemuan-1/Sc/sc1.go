package main

import "fmt"

func main() { 

	var nilai1 int
	fmt.Println("Masukan Nilai 1 :")
	fmt.Scanln(&nilai1)

	var nilai2 int
	fmt.Println("Masukan Nilai 2 :")
	fmt.Scanln(&nilai2)

    var operator string
	fmt.Println("Masukan Operator (x, +, -, :, %) :")
	fmt.Scanln(&operator)

	switch operator {
		case "x":
		    fmt.Println(nilai1 * nilai2)
		case "+":
			fmt.Println(nilai1 + nilai2)
		case "-":
			fmt.Println(nilai1 - nilai2)
		case ":":
			fmt.Println(nilai1 / nilai2)
		case "%":
			fmt.Println(nilai1 % nilai2)
		default:
            fmt.Println("Invalid Operator")
	} 
}