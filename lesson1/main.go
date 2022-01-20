package main

import (
	"fmt"
)

func main() {
	var i int = 100
	var i32 int32 = int32(i)
	fmt.Println(i, i32)

	fmt.Printf("%T\n", i)
	fmt.Printf("%T\n", i32)

}
