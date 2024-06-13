package main

import "fmt"

func main() {

	num := 22

	if num%2 == 0 {
		fmt.Printf("%d is even \n", num)
	} else {
		fmt.Printf("%d is odd", num)
	}

	if no := 22; no%2 == 0 {
		fmt.Printf("%d is even", no)
	} else {
		fmt.Printf("%d is odd", no)

	}
}
