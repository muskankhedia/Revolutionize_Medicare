package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// GetProfileHandler returns the profile details of the patient id
func GetProfileHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("In GetProfileHandler")

	// prevent CORS error
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var form PatientProfile
	var users []PatientProfile

	if err := parseProfileForm(r, &form); err != nil {
		panic(err)
	}
	jsonByteValue, err := ioutil.ReadFile("datastore/profile.json")
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
				form.SugarAF = users[i].SugarAF
				form.SugarBF = users[i].SugarBF
				form.BpS = users[i].BpS
				form.BpD = users[i].BpD
				form.Bmi = users[i].Bmi
				form.Temp = users[i].Temp
				form.Pulse = users[i].Pulse
				form.Resp = users[i].Resp
				form.Gender = users[i].Gender
			}
		}
	}
	res, err := json.Marshal(form)
	w.Write(res)
}

// UpdateProfileHandler updates the profile details of the patient id
func UpdateProfileHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("In UpdateProfileHandler")

	// prevent CORS error
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var form PatientProfile
	if err := parseProfileForm(r, &form); err != nil {
		panic(err)
	}

	// appending the new user to the list of users
	jsonByteValue, err := ioutil.ReadFile("datastore/profile.json")
	if err != nil {
		panic(err)
	}
	var patients []PatientProfile
	if len(jsonByteValue) != 0 {
		err = json.Unmarshal(jsonByteValue, &patients)
	}
	patients = append(patients, form)
	result, err := json.Marshal(patients)
	err = ioutil.WriteFile("datastore/profile.json", result, 0777)
	if err != nil {
		panic(err)
	}
	w.Write([]byte("true"))
}

// PatientProfile is a struct that stores patient profile data
type PatientProfile struct {
	PatientID int     `json:"patientid"`
	SugarBF   float64 `json:"bsugar"`
	SugarAF   float64 `json:"asugar"`
	BpD       float64 `json:"dbp"`
	BpS       float64 `json:"sbp"`
	Bmi       float64 `json:"bmi"`
	Temp      float64 `json:"temp"`
	Pulse     float64 `json:"pulse"`
	Resp      float64 `json:"resp"`
	Gender    string  `json:"gender"`
}
