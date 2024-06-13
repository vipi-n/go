package main

import "fmt"

func main() {

	//no := 22

	switch no := 22; no % 2 {
	case 0:
		fmt.Printf("%d is even", no)

	case 1:
		fmt.Printf("%d is odd", no)
	}

	//simulate nested-if in switch case

	num := 21

	switch {
	case num%2 == 0:
		fmt.Printf("%d is even", num)

	case num%2 == 1:
		fmt.Printf("%d is odd", num)

	default:
		fmt.Printf("%d is not even/odd")
	}

	//fall-through (opposite of "break")
	switch no := 4; no {

	case 0:
		fmt.Println("no == 0")
		fallthrough

	case 1:
		fmt.Println("no == 1")
		fallthrough

	case 2:
		fmt.Println("no == 2")
		fallthrough

	case 3:
		fmt.Println("no == 3")
		fallthrough

	case 4:
		fmt.Println("no <= 4")
		fallthrough

	case 5:
		fmt.Println("no <=5")
		fallthrough

	case 6:
		fmt.Println("no <=6")
	}

	//fallthrough applied

	fmt.Println("fallthrough applied")

	switch plan := "SUPER"; plan {

	case "SUPREME":
		fmt.Println("[SUPREME] : Offline download allowed!")
		fallthrough
	case "SUPER":
		fmt.Println("[SUPER] : For a family of 3")
		fallthrough
	case "PRO":
		fmt.Println("[PRO] : Create your own playlist")
		fallthrough
	case "FREE":
		fmt.Println("[FREE] : Listen to the music")
	}

}
