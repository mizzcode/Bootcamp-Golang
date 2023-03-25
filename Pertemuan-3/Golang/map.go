package main

import "fmt"

// Struct
type Biodata struct {
	Name, Address string
	Age int
}

func main() {
	// map data
	// key-value. 
	// key harus unik
	
	data := map[string]string{
		"key1": "value1",
        "key2": "value2",
	}

	mizz := Biodata {
		Name : "Mizz",
		Address : "Jln. Bekasi",
        Age : 19,
	}

	data["testing"] = "testapp"

	fmt.Println("Name:", mizz.Name)
	fmt.Println("Address:", mizz.Address)
	fmt.Println("Age:", mizz.Age)

	fmt.Println(data)
    fmt.Println(len(data))

	for key, value := range data {
        fmt.Println(key, ":", value)
    }
}