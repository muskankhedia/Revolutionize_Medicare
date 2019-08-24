package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// SignupHandler will parse the sign up form and create a new Patient
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("At signup route")

	// prevent CORS error
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}

	b, err := json.Marshal(form)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println(string(b))

	http.Redirect(w, r, "/profile", http.StatusFound)
}

// SignupForm contains the details entered by the user in the signup form
type SignupForm struct {
	Name       string `schema:"name"`
	Email      string `schema:"email"`
	Password   string `schema:"password"`
	Age        string `schema:"age"`
	DOB        string `schema:"dob"`
	BloodGroup string `schema:"bloodgroup"`
	BirthType  string `schema:"birthtype"`
}
