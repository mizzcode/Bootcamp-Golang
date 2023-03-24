package main

import "fmt"

func main() {
	clubs := [...]string {
		"Emyu",
		"Barcelona",
		"Real Madrid",
		"PSG",
		"Indonesia",
		"San Francisco",
        "Belgium",
		"Polandia",
	}

	for i := 0; i < len(clubs); i++ { 
		fmt.Println(clubs[i])
	}
}

