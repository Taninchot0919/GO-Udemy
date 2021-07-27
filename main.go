package main

import (
	"fmt"

	"github.com/taninchot0919/learning-go/helpers"
)

const numPool = 10

func CalculateValue(intChanel chan int) {
	randomNumber := helpers.RandomNumber(numPool) // ส่งเลขไป random แล้วนำกลับมาให้ randomNumber
	intChanel <- randomNumber                     // เหมือนเป็นการเอาค่ามาใส่ให้กับ chanel ต้องใช้เครื่องหมาย <- ใช้ = ไม่ได้
}
func main() {
	intChanel := make(chan int) // เป็นการสร้าง chanel ของ int ขึ้นมา
	defer close(intChanel)      // defer คือรอทุกตัวรันเสร็จค่อยทำงาน

	go CalculateValue(intChanel) // ใช้ go routine 

	num := <-intChanel // เหมือนเป็นการเอาเลขจาก chanel มาใส่ตัวแปร num
	fmt.Println(num)
}
