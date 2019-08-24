package controllers

import (
	"encoding/json"
	"io/ioutil"
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

	// appending the new user to the list of users
	jsonByteValue, err := ioutil.ReadFile("users.json")
	if err != nil {
		panic(err)
	}
	var users []SignupForm
	if len(jsonByteValue) != 0 {
		err = json.Unmarshal(jsonByteValue, &users)
	}
	users = append(users, form)
	result, err := json.Marshal(users)
	err = ioutil.WriteFile("users.json", result, 0777)
	if err != nil {
		panic(err)
	}

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
