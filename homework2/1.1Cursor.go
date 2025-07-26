package main

import "fmt"

func Adder(a *int) {
	*a += 10
}

func main() {
	a := 5
	Adder(&a)
	fmt.Println(a) // Output: 15
}
