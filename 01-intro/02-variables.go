package main

import "fmt"

func main() {

	/*
		var userName string
		var id int
		var marks float64

		userName = "vipin"
		id = 1
		marks = 19.2
	*/

	/*
		var userName string = "vipin"
		var id int = 1
		var marks float64 = 19.2
	*/

	/* 	var userName = "vipin"
	   	var id = 1
	   	var marks = 19.9
	*/

	/* 	userName := "Vipin"
	   	id := 1
	   	marks := 99.9
	*/

	/* 	var username, college string
	   	var id, age int
	   	var marks float64

	   	username = "vipin"
	   	college = "PSU"

	   	id = 1
	   	age = 22
	   	marks = 99.1
	*/

	/* 	var (
	   		username, college string
	   		id, age           int
	   		marks             float64
	   	)
	*/

	/* var (
		username, college = "vipin", "CTC"
		id, age           = 1, 21
		marks             = 22.2
	)
	fmt.Printf("Username %s with id %d has marks %f his age is %d and college %s \n", username, id, marks, age, college)
	*/

	/* 	var (
	   		x, y, sum = 100, 200, "sume of %d and %d is %d"
	   		result    = x + y
	   	)
	   	fmt.Printf(sum, x, y, result)
	*/
	x, y, str := 100, 200, "sume of %d and %d is %d"
	result := x + y
	fmt.Printf(str, x, y, result)

}
