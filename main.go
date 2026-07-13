package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/student", ViewStudent)
	http.HandleFunc("/students/new", NewStudent)

	fmt.Println("server running")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
