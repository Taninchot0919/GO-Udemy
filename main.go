package main

import (
	"fmt"
	"sort"
)

type User struct{
	Firstname string
	LastName string
}
func main() {
	// เหมือนเป็นการบอกว่า key จะเป็นอะไรแล้วก็ value เป็นอะไร
	// myMap := make(map[string]interface{}) // ถ้าเราไม่รู้ value ที่จะรับมาเป็น type อะไรสามารถใส่เป็น interfaces ได้
	myMap := make(map[string]string) 

	myMap["dog"] = "Puddle"
	myMap["other-dog"] = "Mha"

	fmt.Println(myMap["dog"])
	fmt.Println(myMap["other-dog"])

	myMap["dog"] = "IO"
	fmt.Println(myMap["dog"])
	// เหมือนกับเจ้า map เนี่ยเก็บแบบให้ข้างใน [] เป็น key แล้วถ้านำตัวเดิมมาเท่่ากับมันก็เลยเป็นค่าใหม่

	myMap2 := make(map[string]int)
	myMap2["one"] = 1
	myMap2["two"] = 2
	fmt.Println(myMap2["one"])
	fmt.Println(myMap2["two"])

	myMap3 := make(map[string]User)
	me := User{
		Firstname: "Taninchot",
		LastName: "Phuwaleortthiwat",
	}
	myMap3["me"] = me
	fmt.Println(myMap3["me"].Firstname, myMap3["me"].LastName)

	// ------------------- Slice -------------------
	fmt.Println("=============== Slice ===============")
	// คล้ายๆกับ array ใน ภาษาอื่น
	var myString []string

	// การใช้ append เหมือนเป็นการใส่ child ไว้เป็นตัวสุดท้าย
	myString = append(myString, "Taninchot")
	myString = append(myString, "Art")
	fmt.Println(myString)

	var myInt []int
	myInt = append(myInt, 3)
	myInt = append(myInt, 1)
	myInt = append(myInt, 2)
	
	fmt.Println(myInt)

	sort.Ints(myInt)
	fmt.Println(myInt)

	numbers := []int{1,2,3,4,5,6,7,8,9,10}

	fmt.Println(numbers)
	fmt.Println(numbers[6:9])
}
