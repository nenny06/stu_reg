package main

import (
	"html/template"
	"net/http"
	"strconv"
)

func renderTemplate(w http.ResponseWriter, file string, data any) {
	tmpl, err := template.ParseFiles("templates/base.html", file)

	if err != nil {
		http.Error(w, "Cannot parse file", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	renderTemplate(w, "templates/home.html", nil)
}

func ViewStudent(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/student" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	converted_id, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	for _, student := range students {
		if student.ID == converted_id {
			renderTemplate(w, "templates/student.html", student)
			return
		}
	}
	http.Error(w, "Student not found", http.StatusNotFound)
	return
}

func NewStudent(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/students/new" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	renderTemplate(w, "templates/newstudent.html", nil)
}

func RegisterStudent(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/students/create" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	age := r.FormValue("age")
	course := r.FormValue("course")
	email := r.FormValue("email")

	if firstname == "" || lastname == "" || age == "" || course == "" || email == "" {
		http.Error(w, "Empty parameter", http.StatusBadRequest)
		return
	}

	convert_age, err := strconv.Atoi(age)
	if err != nil {
		http.Error(w, "Age must be a number", http.StatusBadRequest)
		return
	}

	nextid := len(students) + 1

	newstudent := Student{
		ID:        nextid,
		FirstName: firstname,
		LastName:  lastname,
		Age:       convert_age,
		Course:    course,
		Email:     email,
	}

	students = append(students, newstudent)

	http.Redirect(w, r, "/students", http.StatusSeeOther)

}

func UpdateForm(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/students/edit" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	id := r.FormValue("id")

	if id == "" {
		http.Error(w, "empty id parameter", http.StatusBadRequest)
		return
	}

	convert_id, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}

	for _, student := range students {
		if student.ID == convert_id {
			renderTemplate(w, "templates/updateform.html", student)
		}
	}
	http.Error(w, "Student not found", http.StatusBadRequest)
	return
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/students/update" {
		http.NotFound(w,r)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	id := r.FormValue("id")
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	age := r.FormValue("age")
	course := r.FormValue("course")
	email := r.FormValue("email")

	if id == "" || firstname == "" || lastname == "" || age == "" || course == "" || email == "" {
		http.Error(w, "one or more parameter missing", http.StatusBadRequest)
		return
	}

	convert_id, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Wrong age input", http.StatusBadRequest)
		return
	}

	convert_age, err := strconv.Atoi(age)
	if err != nil {
		http.Error(w, "Wrong age input", http.StatusBadRequest)
		return
	}

	for i := range students {
		if students[i].ID == convert_id {
			students[i].FirstName = firstname
			students[i].LastName = lastname
			students[i].Age = convert_age
			students[i].Course = course
			students[i].Email = email
			http.Redirect(w, r, "/students", http.StatusSeeOther)
			return
		}
	}
	http.Error(w, "Internal server error", http.StatusInternalServerError)
	return

}

func ViewStudents(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/students" {
		http.NotFound(w,r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	renderTemplate(w, "templates/students.html", students)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/students/delete" {
		http.NotFound(w,r)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.FormValue("id")

	if id == "" {
		http.Error(w, "ID required", http.StatusBadRequest)
		return
	}

	convert_id, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}

	for i := range students {
		if students[i].ID == convert_id {
			students = append(students[:i], students[i + 1:]...)
			http.Redirect(w,r, "/students", http.StatusSeeOther)
			return
		}
	}
	http.Error(w, "Student not found", http.StatusBadRequest)
	return
}