package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/pic"
)

func getSlice() {
	numbers := [10]int{3, 6, 1, 2, 8, 23, 545, 12}

	var sli1 []int = numbers[2:8]
	var sli2 []int = numbers[5:8]
	fmt.Println("Matriz:", numbers)
	fmt.Println("Slice 1:", sli1)
	fmt.Println("Slice 2:", sli2)
	sli2[1] = 999
	fmt.Println("======= Slice2[1] = 999 =======")
	fmt.Println("Matriz:", numbers)
	fmt.Println("Slice 1:", sli1)
	fmt.Println("Slice 2:", sli2)
	fmt.Println("Comprimento | Capacidade => sli1:", len(sli1), cap(sli1))
}
func getMatriz() {
	var a [3]string
	a[0] = "Hello"
	a[1] = "World, "
	a[2] = "Melqui"
	fmt.Println(a)
}
func getLiteralSlice() {
	fmt.Println("============== Literal Slices")
	q := []int{3, 4, 5}
	r := []bool{true, true, false, true}
	s := []struct {
		i int
		b bool
	}{
		{2, false},
		{1, true},
		{5, false},
	}
	fmt.Println(q, r, s)

	var sli3 []int
	fmt.Println("Valor ZERO da slice:", sli3)
	if sli3 == nil {
		fmt.Println("sli3 is nil")
	}
}
func sliceMake() {
	s := make([]int, 5)
	fmt.Println("capacidade S:", cap(s))
	fmt.Println("tamanho S:", len(s))
	fmt.Println("Slice S by Make:", s)
	c := s[:2]
	fmt.Println("capacidade C:", cap(c))
	fmt.Println("tamanho C:", len(c))
	fmt.Println("Slice C by Make:", c)

	c = c[:cap(c)]
	fmt.Println("Alterado tamanho de C usando Capacidade de C:", len(c))
	fmt.Println("capacidade C:", cap(c))
	fmt.Println("Slice C by Make:", c)
}
func sliceWithSlices() {
	board := [10][]string{
		[]string{"-", "_", "*"},
		[]string{"_", "-", "*"},
		[]string{"*", "_", "-"},
	}
	fmt.Println("Slice de slice:", board)
	board[0][0] = "X"
	fmt.Println("Slice de slice:", board)
	board[2][2] = "*"
	fmt.Println("Slice de slice:", board)
	for i, v := range board { //para ignorar o indice ou o valor usa-se _
		fmt.Printf("Item %v: %s\n", i, strings.Join(v, " "))
	}
}
func addValuesSlice(value int) []int {
	sl := []int{4, 5, 1}
	sl = append(sl, value)
	return sl
}
func Pic(dx, dy int) [][]uint8 {
	slice := make([][]uint8, dy)
	for i := range slice {
		childSlice := make([]uint8, dx)
		for x := range childSlice {
			childSlice[x] = uint8(x ^ i)
		}
		slice[i] = childSlice
	}
	return slice
}
func main() {
	getMatriz()
	fmt.Println("==============")
	getSlice()
	fmt.Println("==============")
	sliceMake()
	fmt.Println("==============")
	sliceWithSlices()
	slice := addValuesSlice(58)
	fmt.Println("Slice incrementado com append:", slice)

	pic.Show(Pic)
}
