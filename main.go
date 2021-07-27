package main

import (
	"fmt"
	"github.com/taninchot0919/learning-go/helpers"
)

func main() {
	fmt.Println("Hello Module")

	var myVar helpers.SomeType
	myVar.TypeName = "Taninchot"
	myVar.TypeNumber = 20
	fmt.Println(myVar)
}
