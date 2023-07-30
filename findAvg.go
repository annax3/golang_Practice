package main

import (
	"fmt"
	"os"
)

func main() {
	var sum float64

	var n int

	for {
		var val float64

		//if _,err := fmt.Fscanln(os.Stdin, &val); err != nil {
		//	break
		//}

		// these two are same

		//_, err := fmt.Fscanln(os.Stdin, &val)
		//if err != nil {
		//	break
		//}

		_, err := fmt.Fscanln(os.Stdin, &val)
		if err != nil {
			break
		}

		sum += val
		n++
	}

	if n == 0 {
		fmt.Println(os.Stderr, "no values")
		os.Exit(-1)
	}
	fmt.Println("The average is ", sum/float64(n))
}
