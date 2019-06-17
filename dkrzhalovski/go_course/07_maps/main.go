package main

import "fmt"

func main() {
	myMap := make(map[string]string)

	myMap["Dado"] = "Ana"
	myMap["Ana"] = "Dado"

	fmt.Println(myMap["Dado"])

	quickMap := map[string]string{"1": "2", "2": "3"}

	fmt.Println(quickMap["1"])
}
