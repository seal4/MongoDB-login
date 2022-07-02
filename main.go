package main

import (
	"html/template"
	"log"
	"net/http"
)

func signup(w http.ResponseWriter, r *http.Request) {
	var filename = "signup.html"
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

func signupSubmit(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	log.Println(username, password)
	switch {
	case username == "" || password == "":
		log.Println("username or password field has no value")
	default:
		upload(username, password)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/signup":
		signup(w, r)
	case "/signup-submit":
		signupSubmit(w, r)
	case "/login":
		login(w, r)
	case "/login-submit":
		loginSubmit(w, r)
	default:
		log.Println("User has logged onto page!")
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
	log.Println("Started site on :8080")
	login(nil, nil)
}
