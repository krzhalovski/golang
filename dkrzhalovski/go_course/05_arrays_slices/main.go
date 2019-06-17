package main

import "fmt"

func main() {
	var arr [2]string

	arr[0] = "hey"
	arr[1] = "you"

	fmt.Println(arr)

	otherArr := []string{"HEY", "YOU"}
	fmt.Println(otherArr)
}
