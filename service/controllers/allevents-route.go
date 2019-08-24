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

	var data []EventBlock
	json.Unmarshal(byteValue, &data)
	if err != nil {
		fmt.Println("err")
	}

	fmt.Println(data)

	var list []EventBlock

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

//UpdateSuccessHandler Updates the success 
func UpdateSuccessHandler(w http.ResponseWriter, r *http.Request) {

	// prevent CORS error
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	eventid, err := strconv.Atoi(r.FormValue("eventid"))
	success, err := strconv.ParseBool(r.FormValue("success"))

	fmt.Println(eventid)

	jsonFile, err := os.Open("datastore/events.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var data []EventBlock
	json.Unmarshal(byteValue, &data)
	if err != nil {
		fmt.Println("err")
	}

	for i := 0; i < len(data); i++ {

		if eventid == data[i].EventID {
			data[i].Success =  success
			break
		} 
	}

	result, err := json.Marshal(data)
	err = ioutil.WriteFile("datastore/events.json", result, 0777)
	if err != nil {
		panic(err)
	}

	w.Write([]byte(`success field updated`))

}