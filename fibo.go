package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	beforeValue := 0
	value := 0
	return func() int {
		newValue := 0
		if value == 0 || beforeValue == value {
			newValue = value
			beforeValue = value
			value++
			return newValue
		}
		newValue = value + beforeValue
		beforeValue = value
		value = newValue
		return value
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
