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

func parseProfileForm(r *http.Request, data *PatientProfile) error {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	data.PatientID, _ = strconv.Atoi(r.PostFormValue("patientid"))
	data.SugarAF, _ = strconv.ParseFloat(r.PostFormValue("asugar"), 64)
	data.SugarBF, _ = strconv.ParseFloat(r.PostFormValue("bsugar"), 64)
	data.BpD, _ = strconv.ParseFloat(r.PostFormValue("dbp"), 64)
	data.BpS, _ = strconv.ParseFloat(r.PostFormValue("sbp"), 64)
	data.Bmi, _ = strconv.ParseFloat(r.PostFormValue("bmi"), 64)
	data.Temp, _ = strconv.ParseFloat(r.PostFormValue("temp"), 64)
	data.Pulse, _ = strconv.ParseFloat(r.PostFormValue("pulse"), 64)
	data.Resp, _ = strconv.ParseFloat(r.PostFormValue("resp"), 64)
	data.Gender = r.PostFormValue("gender")
	return nil
}
