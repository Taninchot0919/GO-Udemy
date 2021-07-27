package main

import (
	"fmt"
)

type myStruct struct {
	Firstname string
}

// เราจะเรียก (m *myStruct) ว่า receiver
// เหมือนว่ามันจะลิ้งกับ struct ที่เราสร้างขึ้นมาเลย ลองดูจากตัวอย่างที่เรารันจะทำให้เราเห็นภาพได้มากขึ้น
func (m *myStruct) printFirstName() string {
	return m.Firstname
}

func main() {
	var myVar myStruct
	myVar.Firstname = "Taninchot"

	myVar2 := myStruct{
		Firstname: "Taninchot2",
	}

	fmt.Println("myVar is set to :", myVar.printFirstName())
	fmt.Println("myVar2 is set to :", myVar2.printFirstName())

}
