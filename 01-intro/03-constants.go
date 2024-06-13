package main

import "fmt"

func main() {

	const pi float64 = 3.14
	const pie = 3.14

	//cannot reassign
	//pi = 3.44
	fmt.Println(pi)

	const (
		i int     = 31
		y float64 = 20.99
	)
	//type inference

	const (
		x  = 100
		y1 = 200
	)

}
