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
	jsonByteValue, err := ioutil.ReadFile("../datastore/users.json")
	if err != nil {
		panic(err)
	}
	var users []SignupForm
	if len(jsonByteValue) != 0 {
		err = json.Unmarshal(jsonByteValue, &users)
	}
	form.PatientID = len(users) + 1
	users = append(users, form)
	result, err := json.Marshal(users)
	err = ioutil.WriteFile("../datastore/users.json", result, 0777)
	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/profile", http.StatusFound)
}

// SignupForm contains the details entered by the user in the signup form
type SignupForm struct {
	PatientID  int    `schema:"patientid" json:"PatientID"`
	Name       string `schema:"name" json:"Name"`
	Email      string `schema:"email" json:"Email"`
	Age        string `schema:"age" json:"Age"`
	DOB        string `schema:"dob" json:"Dob"`
	BloodGroup string `schema:"bloodgroup" json:"BloodGroup"`
	BirthType  string `schema:"birthtype" json:"BirthType"`
}
