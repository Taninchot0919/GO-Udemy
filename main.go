package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLeg() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Sam",
		Breed: "German Shepperd",
	}

	printInfo(&dog)

	gorilla := Gorilla{
		Name: "Ling",
		Color: "Brown",
		NumberOfTeeth: 38,
	}

	printInfo(&gorilla)
}

func printInfo(a Animal)  {
	fmt.Println("This animals say",a.Says(),"and has",a.NumberOfLeg(),"legs")
}

// -------- Interface --------
// โดยเจ้า interface ในนี้เนี่ยก็เหมือนกันกับ interface ใน java เลย
// พวกที่เป็น receiver ส่วนใหญ่แล้วจะเป็น pointer ดังนั้นควรใช้แบบนี้ดีกว่าเพื่อที่เป็น best practice
func (d *Dog) Says() string {
	return "Booooooooo"
}

func (d *Dog) NumberOfLeg() int {
	return 4
}

func (g *Gorilla) Says() string {
	return "Urghhhhhhh"
}

func (g *Gorilla) NumberOfLeg() int {
	return 2
}