package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/student", ViewStudent)
	http.HandleFunc("/students", ViewStudents)
	http.HandleFunc("/students/new", NewStudent)
	http.HandleFunc("/students/create", RegisterStudent)
	http.HandleFunc("/students/edit", UpdateForm)
	http.HandleFunc("/students/update", UpdateStudent)
	http.HandleFunc("/students/delete", DeleteStudent)
	http.HandleFunc("/students/confirmdelete", ConfirmDeleteSudent)
	http.Handle("/static/", http.StripPrefix("/static/", 
	http.FileServer(http.Dir("static")),),)


	fmt.Println("server running")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
