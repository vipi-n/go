package main

import "fmt"

func main() {

	/* 	const Red int = 0
	   	const Blue int = 1
	   	const yelow int = 2
	*/

	/* 	const (
	   		Red   int = 0
	   		Blue  int = 1
	   		yelow int = 2
	   	)
	*/

	const (
		Red    int = iota //0 - (iota * 2) + 1 soemthing like this also we can do
		Blue   int = iota //1
		yellow int = iota //2
	)
	fmt.Printf("Red : %d\n Blue : %d\n Yellow: %d", Red, Blue, yellow)
}
