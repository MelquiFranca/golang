package main

import "fmt"

func pointers() {
	number := 40
	var pointer *int
	pointer = &number
	fmt.Println("number:", number)
	fmt.Println("value pointer:", *pointer)
	fmt.Println("address pointer:", pointer)
	*pointer = 25
	fmt.Println("number:", number)
	fmt.Println("value pointer:", *pointer)
}

func main() {
	pointers()
}
