package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter a color:")
	var color string
	fmt.Scanln(&color)

	fmt.Println("Enter a plural noun:")
	var plural string
	fmt.Scanln(&plural)

	fmt.Println("Enter a celebrity name:")
	var celebrity string
	fmt.Scanln(&celebrity)

	fmt.Println("Roses are", color)
	fmt.Println(plural, "are blue")
	fmt.Println("I love", celebrity)
}
