package main

import "fmt"

func main() {
	a := 2   // declaration and initialization
	b := 3.9 // declaration and initialization

	fmt.Printf("a type: %T\n", a)  //a type: int
	fmt.Printf("a value: %v\n", a) //a value: 2
	fmt.Printf("b type: %T\n", b)  //b type: float64
	fmt.Printf("b value: %v\n", b) //b value: 3.9

	fmt.Printf("a : %T  %[1]v\n", a) //a : int  2
	fmt.Printf("a : %T  %[1]v\n", a, a)
	//%[1] means it will reuse the perimeter 1 was upper both line are same

	fmt.Printf("b : %T  %[1]v\n", b) //b value: 3.9

}
