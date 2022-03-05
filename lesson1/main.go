package main

import "fmt"

func Sum(s ...int) {
	fmt.Println(s)
	fmt.Println(s)
	fmt.Println(s)
	fmt.Println(s)
	fmt.Println(s)
}

func main() {
	ch1 := make(chan int)
	ch1 <- 1

	ch1 <- 1

	fmt.Println(len(ch1))

}
