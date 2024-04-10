package main

import (
	"fmt"
	"runtime"
)

// ===== FOR =====
// 3 components separated by initial instruction,
// condition expression and post instruction
func count(limit int) {
	for i := 1; i <= limit; i++ {
		fmt.Println("Count:", i)
	}
}
func countForEqualsWhile(limit int) {
	sum := 1
	for sum <= limit {
		fmt.Println("countForEqualsWhile:", sum)
		sum++
	}
}
func infinitLoop() {
	for {
		fmt.Println("no")
	}
}

//  ===== IF =====
func ifExamples(x, y int) int {
	if x < y {
		fmt.Println("x < y:", "Y greater than")
	}
	if z := x - y; z > y {
		fmt.Println("x - y > y:", "X - Y greater than")
		return z
	} else {
		return z // avaliable just within scope if/else
	}
}

//  ===== SWITCH =====
func switchCases(number int) {
	fmt.Println("Switch number to text: ")
	switch number {
	case 1:
		fmt.Println("Is one")
	case 2:
		fmt.Println("Is two")
	case 3:
		fmt.Println("Is tree")
	case 4:
		fmt.Println("Is four")
	case 5:
		fmt.Println("Is five")
	default:
		fmt.Println("Other number")
	}
	fmt.Println("==================")
	switch {
	case number == 3:
		fmt.Println("Number is equals 3")
	case number > 5:
		fmt.Println("Number is greater than 5")
	case number < 3:
		fmt.Println("Number is less than 5")
	default:
		fmt.Println("Number is not 3 and is not greater than 5")
	}
	fmt.Println("==================")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS.X.")
	default:
		fmt.Println("Operational System:", os)
	}
}

//  ===== DEFER =====
func useDefer(name string) {
	defer fmt.Println("Waited to finish function to print this name: ", name)
	fmt.Println("Initialize counter before to show name...")
	number := 20
	for i := 1; i <= number; i++ {
		defer fmt.Println("Push function to stack with value: ", i)
	}
	count(number)
}

func main() {
	count(2)
	fmt.Println("==================")
	countForEqualsWhile(2)
	fmt.Println("==================")
	// infinitLoop()
	fmt.Println("==================")
	z := ifExamples(9, 4)
	fmt.Println("Z:", z)
	fmt.Println("==================")
	switchCases(3)
	switchCases(10)
	switchCases(2)
	switchCases(4)
	fmt.Println("==================")
	useDefer("Melquisedeque Pereira França")
}
