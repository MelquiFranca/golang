package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}
func Walking(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	defer close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walking(t1, ch1)
	go Walking(t2, ch2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		fmt.Println(v1, v2)
		if ok1 != ok2 || v1 != v2 {
			return false
		}
		if ok1 == false {
			return true
		}
	}
}

func main() {
	result := Same(tree.New(2), tree.New(1))
	fmt.Println(result)
	result = Same(tree.New(1), tree.New(1))
	fmt.Println(result)
}
