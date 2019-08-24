package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	// "strings"
	"encoding/json"
	"io/ioutil"
	"os"
)

//EventDetails contains event details
type EventDetails struct {
	PatientID int      `json:"PatientID"`
	Event     string   `json:"Event"`
	Medicine  []string `json:"Medicine"`
	TimeSFO   int      `json:"TimeSFO"`
	Success   bool     `json:"Success"`
}

//AllEventsHandler returns all the events of the patient
func AllEventsHandler(w http.ResponseWriter, r *http.Request) {

	// prevent CORS error
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	patientID, err := strconv.Atoi(r.FormValue("patientid"))
	fmt.Println(patientID)

	jsonFile, err := os.Open("datastore/events.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var data []EventDetails
	json.Unmarshal(byteValue, &data)
	if err != nil {
		fmt.Println("err")
	}

	fmt.Println(data)

	var list []EventDetails

	for i := 0; i < len(data); i++ {
		if data[i].PatientID == patientID {
			list = append(list, data[i])
		}
	}

	j, err := json.Marshal(list)
	if err != nil {
		panic(err)
	}
	w.Write(j)

}
