package main

import "fmt"

func main() {
	fmt.Println("Welcome To-Do-List")

	var daftarTodoList []string

	for true {
		var kegiatan string
		fmt.Println("Masukkan Kegiatan")
		fmt.Scanln(&kegiatan)

		daftarTodoList = append(daftarTodoList, kegiatan)

		var pertanyaan string
		fmt.Println("Lagi atau tidak (y/n)")
		fmt.Scanln(&pertanyaan)
		
		if pertanyaan == "n" {
            break
        }
	}

	for i, todolist := range daftarTodoList { 
		fmt.Println(i+1, todolist)
	}


}