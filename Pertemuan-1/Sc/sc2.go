package main

import "fmt"

func bintang(panjang int) string {
	hasil := ""
	for i := 0; i <= panjang; i++ { 
		for j := 0; j < i; j++ {
			hasil += "* "
		}
		hasil += "\n"
	}
	return hasil
}

func main() { 
	fmt.Println(bintang(3))
}