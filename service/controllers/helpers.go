package controllers

import (
	"net/http"
	"strconv"
)

func parseForm(r *http.Request, data *SignupForm) error {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	data.PatientID, _ = strconv.Atoi(r.PostFormValue("patientid"))
	data.Name = r.PostFormValue("name")
	data.Email = r.PostFormValue("email")
	data.Age = r.PostFormValue("age")
	data.DOB = r.PostFormValue("dob")
	data.BloodGroup = r.PostFormValue("bg")
	data.BirthType = r.PostFormValue("b")
	return nil
}
