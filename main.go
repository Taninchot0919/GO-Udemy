package main

import "fmt"

func main() {
	animals := []string{"Dog", "Cat", "Fish", "Horse"}

	fmt.Println("range of animals is ", len(animals)) // เอาไว้หา length

	// i เหมือนตัวรันเลข animal ที่ต่อกันคือ เหมือนเป็นตัวแปร ranges animals เหมือนวนลูป array ว่ามีกี่ตัว
	for i, animal := range animals {
		fmt.Println(i, animal)
	}

	fmt.Println("---------------------")

	// เหมือนเราอยากลูปแต่ไม่อยากให้ใส่เลขด้านหน้าก็ใช้วิธีแบบนี้
	for _, animal := range animals {
		fmt.Println(animal)
	}

	fmt.Println("---------------------")

	animalMaps := make(map[string]string)

	animalMaps["Dog"] = "ArTTho"
	animalMaps["Cat"] = "ArTThy"

	// ทีนี้ถ้าหากมันใช้มาจาก map แล้วเนี่ย i ในนี้จะไม่ใช่ตัวเลขละ มันจะเป็น key แทน ส่วน animal ก็จะเป็น value
	for i, animal := range animalMaps {
		fmt.Println(i, animal)
	}

	for _, animal := range animalMaps {
		fmt.Println(animal)
	}

	type User struct {
		Firstname string
		Lastname  string
		Email     string
		Age       int
	}

	users := []User{}

	users = append(users, User{"Taninchot", "Phuwaloertthiwat", "Taninchot0919@live.com", 20})
	users = append(users, User{"Art", "Phuwaloertthiwat", "Taninchot0919@live.com", 22})
	users = append(users, User{"Tanin", "Phuwaloertthiwat", "Taninchot0919@live.com", 21})
	users = append(users, User{"T", "Phuwaloertthiwat", "Taninchot0919@live.com", 24})

	for _, user := range users {
		fmt.Println(user.Firstname, user.Lastname, user.Email, user.Age)
	}

}
