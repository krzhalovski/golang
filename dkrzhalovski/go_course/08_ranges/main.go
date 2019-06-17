package main

import "fmt"

func main() {
	arr := []int{1, 5, 2, 4, 3}

	for i, id := range arr {
		fmt.Println(i, id)
	}

	quickMap := map[string]string{"1": "2", "2": "3"}

	for key, val := range quickMap {
		fmt.Printf("Key: %s, Value: %s\n", key, val)
	}
}
