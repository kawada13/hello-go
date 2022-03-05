package main

import "fmt"

func Sum(c chan int) {
	s := <-c
	fmt.Println(s)
	fmt.Println(s)
	fmt.Println(s)
}

func main() {
	ch1 := make(chan int)
	go Sum(ch1)

	i := 0
	ss := 0

	for i < 100 {
		ch1 <- i
		i++
	}

	fmt.Println(ss)
}
