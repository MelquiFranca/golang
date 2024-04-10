package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type VertexTwo struct {
	X int
	Y int
}

// Structs with pointers
func withPointers() {
	fmt.Println("Using Pointers with structs")
	vert := VertexTwo{1, 2}
	fmt.Println("vert:", vert)
	p := &vert
	(*p).X = 5 // not recommended
	fmt.Println("(*p).X:", (*p).X)
	p.Y = 8 // allowed and recommended
	fmt.Println("p.Y:", p.Y)
	fmt.Println("vert:", vert)
}

// Literal Structs
var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

func literalStructs() {
	fmt.Println(v1, p, v2, v3)
}

func main() {
	var v Vertex = Vertex{5, 8}
	fmt.Println("Vertex:", v)
	v.X = 2
	fmt.Println("Modify Vertex:", v)
	withPointers()
	literalStructs()
}
