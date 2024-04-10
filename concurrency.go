package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func channels() {
	fmt.Println("======Channel======")
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[1:len(s)/2], c) // esse será o segundo
	go sum(s[:len(s)/2], c)  // esse será o último
	go sum(s[:], c)          // Esse será o primeiro resultado retornado
	x, y, z := <-c, <-c, <-c
	fmt.Println(x, y, z)
}
func bufferizedChannels() {
	fmt.Println("======Bufferized Channel======")
	ch := make(chan int, 2) //S
	ch <- 1
	ch <- 2
	// Não pode mais ser armazenado valores, pois o canal está chei
	// ch <- 233         // Descomente o início da linha para visualizar o erro
	fmt.Println(<-ch) // Esvaziando canal
	fmt.Println(<-ch) // Esvaziando canal
	//Agora pode receber dados novamente
	ch <- 18
	ch <- 7
	fmt.Println(<-ch) // Esvaziando canal
	fmt.Println(<-ch) // Esvaziando canal
}
func fibo(n int, c chan uint64) {
	fmt.Println("======Range and Close (Fibonnacci)======")
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- uint64(x)
		x, y = y, x+y
	}
	close(c)
}
func closeAndRange() {
	c := make(chan uint64, 100000000)
	go fibo(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
func fiboSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x: //Aguardam chegar algo no canal para executar
			x, y = y, x+y
		case <-quit: //Aguardam chegar algo no canal para executar
			fmt.Println("quit")
			return
		}
	}
}
func selectFunction() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fiboSelect(c, quit)
}
func selectDefaultFunction() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOOM!")
			return // Finaliza o laço FOR
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
func main() {
	// channels()
	// bufferizedChannels()
	// closeAndRange()
	// selectFunction()
	selectDefaultFunction()
}
