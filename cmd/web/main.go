package main

import (
	"fmt"
	"github.com/taninchot0919/learning-go/pkg/handlers"
	"net/http"
)

const portNumber = ":9000"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
