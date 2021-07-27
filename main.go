package main

import "fmt"

func main() {
	myVar := "TeeRak"

	switch myVar {
	case "cat":
		fmt.Println("Cat")
	case "dog":
		fmt.Println("Dog")
	case "fish":
		fmt.Println("Fish")
	default:
		fmt.Println("Animal > <")
	}
}
