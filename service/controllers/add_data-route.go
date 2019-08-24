package controllers

import (
	"encoding/json"
	// "fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// Message takes incoming JSON payload for writing patient details
type Message struct {
	PatientID int      `json:"PatientID"`
	Event     string   `json:"Event"`
	Medicine  []string `json:"Medicine"`
	TimeSFO   int      `json:"TimeSFO"`
	Success   bool     `json:"Success"`
}

//AddDataHandler adds the data into the user details
func AddDataHandler(w http.ResponseWriter, r *http.Request) {

	// prevent CORS error
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	r.ParseForm()

	patientID, err := strconv.Atoi(r.FormValue("patientid"))
	event := r.FormValue("event")
	medicine := strings.Split(r.FormValue("medicine"), ",")
	timeSinceFirstOccurence, err := strconv.Atoi(r.FormValue("time_since_first_occurance"))
	success, err := strconv.ParseBool(r.FormValue("success"))

	var msg Message
	msg = Message{
		PatientID: patientID,
		Event:     event,
		Medicine:  medicine,
		TimeSFO:   timeSinceFirstOccurence,
		Success:   success,
	}
	var data []Message

	jsonByteValue, err := ioutil.ReadFile("datastore/events.json")
	if err != nil {
		panic(err)
	}

	if len(jsonByteValue) != 0 {
		err = json.Unmarshal(jsonByteValue, &data)
	}

	data = append(data, msg)
	result, err := json.Marshal(data)
	err = ioutil.WriteFile("datastore/events.json", result, 0777)
	if err != nil {
		panic(err)
	}

	w.Write([]byte(`true`))
}
