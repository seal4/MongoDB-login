package main

import (
	"html/template"
	"log"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	var filename = "login.html"
	t, err := template.ParseFiles(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = t.ExecuteTemplate(w, filename, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func loginsubmit(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	log.Println(username, password)
	upload(username, password)
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/login":
		login(w, r)
	case "/login-submit":
		loginsubmit(w, r)
	default:
		log.Println("User has logged onto page!")
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
	log.Println("Started site on :8080")

}
