package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // ที่ใช้ * ด้วยเพราะชี้ pointer ไปยัง request ที่รับมา
		n, err := fmt.Fprintf(w, "Hello World") // เหมือนเอา hello world มาใส่ใน response

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Number of byte :", n)
	})
	
	// คือเจ้า listenAndServe มันรีเทิน err ออกมาแต่เราไม่สนใจอะไรเพราะไม่ได้ใช้ก็เลยใช้ _ ไว้
	_ = http.ListenAndServe(":9000", nil)
}
