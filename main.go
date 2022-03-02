package main

import (
	"fmt"
	"net/http"
)

func main() {

	Set_urls()
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
