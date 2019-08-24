package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//ProfileHandler returns the profile details of the patient id
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("In ProfileHandler")

	// prevent CORS error
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var form SignupForm
	var users []SignupForm

	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	jsonByteValue, err := ioutil.ReadFile("../datastore/users.json")
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
				form.Name = users[i].Name
				form.Email = users[i].Email
				form.Age = users[i].Age
				form.DOB = users[i].DOB
				form.BloodGroup = users[i].BloodGroup
				form.BirthType = users[i].BirthType
			}
		}
	}
	res, err := json.Marshal(form)
	if err != nil {
		panic(err)
	}
	// return ([]byte(res))
	w.Write(res)

}
