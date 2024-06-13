package main

import "fmt"

func main() {

	fmt.Println("Ver 1.0")

	for i := 0; i < 10; i++ {
		fmt.Printf("i = %d\n", i)
	}

	fmt.Println("ver 2.0 - while")

	i := 0
	for i < 10 {
		fmt.Printf("i = %d\n", i)
		i++
	}

	fmt.Println("Ver 3.0")
	for i := 0; ; i++ {

		fmt.Printf("i = %d\n", i)

		if i >= 10 {
			break
		}
	}

	fmt.Println("Ver 4.0")

	for i := 0; i < 10; {

		fmt.Printf("i = %d\n", i)
		i++
	}

	fmt.Println("Ver 5.0")

	no := 0
	for {
		fmt.Printf("no = %d\n", no)
		no++

		if no >= 10 {
			break
		}
	}

	fmt.Println("Ver 6.0")

	for i := range 10 {
		fmt.Println(i)
	}

	for range 10 {
		fmt.Println("do something")
	}

	//using labels

X_LOOP:
	for i := 0; i < 10; i++ {
		for y := 0; y < 10; y++ {
			fmt.Printf("i = %d , y = %d\n", i, y)
			if i == y {
				fmt.Println("+++++++++++++++++++++++")
				continue X_LOOP
			}
		}
	}
}
