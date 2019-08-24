package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//LoginHandler returns the medical history details of the patient id
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("In login handler")
	var present = "false"
	// prevent CORS error
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var form SignupForm
	var users []SignupForm

	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	jsonByteValue, err := ioutil.ReadFile("datastore/users.json")
	if err != nil {
		panic(err)
	}
	if len(jsonByteValue) != 0 {
		err = json.Unmarshal(jsonByteValue, &users)
		if err != nil {
			panic(err)
		}
		for i := 0; i < len(users); i++ {
			if form.PatientID == users[i].PatientID {
				present = "true"
			}
		}
	}
	log.Println(present)
	w.Write([]byte(present))
}
