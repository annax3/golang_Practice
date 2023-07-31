package main

import "fmt"

func main() {
	s := "hello, world"

	a := len(s)

	fmt.Println(a)

	b := s[:5]

	fmt.Println(b)

	c := s[7:]

	fmt.Println(c)
}
