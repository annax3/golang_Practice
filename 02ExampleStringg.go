package main

import "fmt"

func main() {
	s := "the quick brown fox"

	a := len(s)                 //19
	b := s[:3]                  //the
	c := s[4:9]                 //quick
	d := s[4:11]                //quick b
	e := s[:4] + "slow" + s[9:] //the slow brown fox

	//s[5]= 'a'
	s += "es"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(s)
}
