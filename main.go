package main

import (
	"encoding/json"
	"fmt"
	// "github.com/taninchot0919/learning-go/helpers"
)

type Person struct { // เหมือนอารมณ์สร้าง obj มารอ
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	// ตัวอย่าง json
	myJson := `
	[
		{
			"first_name" : "Clark",
			"last_name" : "Kent",
			"hair_color" : "Black",
			"has_dog" : true
		},
		{
			"first_name" : "Bruce",
			"last_name" : "Wayne",
			"hair_color" : "Black",
			"has_dog" : false
		}
	]
	`

	// ------------------ Read Json ------------------
	var unmarshalled []Person // สร้างตัวแปรมาตัวนึงเป็น type Object Person

	//เจ้า json.Unmarshal เนี่ยเป็นการแปลง json ให้เป็น byte แล้วไปเก็บสักที่
	// โดย parameter ตัวแรกต้องการ byte เราก็แปลง string กลายเป็น byte โดยใช้ []byte()
	// ต่อมาก็เก็บไปที่ unmarshalled โดยชี้ไปตรงๆ
	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		fmt.Println("Error unmarshalling json", err)
	}

	fmt.Printf("Unmarshalled : %v", unmarshalled)
	fmt.Println("")

	// ------------------ Write Json ------------------
	// สร้างข้อมูลมาเล่นๆ
	var mySlice []Person

	var m1 Person
	m1.FirstName = "Wally"
	m1.LastName = "West"
	m1.HairColor = "Red"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Diana"
	m2.LastName = "Prince"
	m2.HairColor = "Black"
	m2.HasDog = true

	mySlice = append(mySlice, m2) //เอาข้อมูลมายัดลงให้หมด

	//แปลงข้อมูลให้เป็น json ใช้ MarshalIndent เหมือนเป็น json.Stringify
	// พารามิเต้อตัวที่ 3เหมือนเป็นช่องเว้น white space
	newJson, err := json.MarshalIndent(mySlice, "", "	")

	if err != nil {
		fmt.Println("Error while marshalling", err)
	}

	// เนื่องจากตอนมันแปลงอะ มันรีเทินมาเป็น byte เราต้องแปลงให้มันรีเทินเป็น string
	fmt.Println(string(newJson))
}
