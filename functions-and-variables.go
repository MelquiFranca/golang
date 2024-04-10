package main

import (
	"fmt"
)

// Should have 0 or more parameters
func add(x int, y int) int {
	return x + y
}

// When 2 or more parameters has same type,
// can omit type of all with exception the last
func multiplyItems(x, y int) int {
	return x * y
}

// A function can returns any numbers of results
func addAndMultiply(x int, y int) (int, int) {
	return add(x, y), multiplyItems(x, y)
}

// Named return value (clean return)
func subtract(x int, y int) (z int) {
	z = x - y
	return
}

// Declaring variables and Some basic types
var (
	name string = "Melqui"
	year int    = 33
) // not allowed short statement
func variables() (string, bool, int, int32) {
	var city, userData string
	city = "Rio de Janeiro"
	country := "Brazil" // short statement
	var yearString = string(year)
	userData = name + " - " + yearString + " years - " + city + " - " + country
	var numberInt32 rune = 200
	return userData, false, 1000, numberInt32
}

// Declaration of constants
const PI float32 = 3.14159265359

func main() {
	fmt.Println("add:", add(5, 4))
	fmt.Println("multiplyItems:", multiplyItems(5, 4))
	x, y := addAndMultiply(9, 7)
	fmt.Printf("addAndMultiply:%v, %v", x, y)
	fmt.Println()
	fmt.Println("subtract:", subtract(9, 3))
	name, resultBool, numberInt, numberIntRune := variables()
	fmt.Printf("variables:%v, %v,  %v, %v", name, resultBool, numberInt, numberIntRune)
	fmt.Println()
	fmt.Println("PI value:", PI)
}
