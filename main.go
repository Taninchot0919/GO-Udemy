package main

import "fmt"

func main() {
	var myString string
	myString = "Green"

	fmt.Println("myString is set to", myString)

	changeUsingPointer(&myString) // การใช้แบบนี้เหมือนเป็นการส่งค่า memmory
	fmt.Println(myString)
}

// ถ้าหากเราไม่เปลี่ยนไปใช้ pointer ค่ามันไม่ได้เปลี่ยนไปจริงๆ
// โดยการทำแบบนี้เนี่ยทำให้ function นี้ไม่ต้อง return
func changeUsingPointer(s *string) { // รอรับค่า memmory
	// เหมือนเราแค่เอา memory มาเปลี่ยนค่าตรงๆเลย
	fmt.Println("s is", s)
	fmt.Println("*s is", *s)
	fmt.Println("&s is", &s)

	newValue := "Red"
	*s = newValue
}
