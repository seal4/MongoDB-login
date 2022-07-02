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

func loginSubmit(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	log.Println(username, password)
	switch {
	case username == "" || password == "":
		log.Println("username or password field has no value")
	default:
		Compare(username)
		if username == "" {
			log.Println("username did not match MongoDB database")
		}
		//CHECK DATABASE FOR MATCHING USERNAME AND PASSWORD HASH!
		//Compare(bson.D{username}) //compares but fails because it is not a bson type
	}
}
