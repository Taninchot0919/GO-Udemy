package main

import "fmt"

func main() {
	fmt.Println("Hello, world")

	var whatToSay string
	var i int

	whatToSay = "Taninchot"
	i = 5

	fmt.Println(whatToSay)
	fmt.Println("i is set to", i)

	whatWasSaid, theOtherWasSaid := saySomething()

	fmt.Println("The function return : ", whatWasSaid, " and ", theOtherWasSaid)
}

func saySomething() (string, string) {
	return "something function", "2nd from something function"
}
