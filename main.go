package main

import (
	"fmt"
	"groupie-tracker/groupie"
	"net/http"
)

func main() {
	http.HandleFunc("/", groupie.Handler)
	http.ListenAndServe(":8080", nil)
	fmt.Print("Listen And Serve Port 8080")
}
