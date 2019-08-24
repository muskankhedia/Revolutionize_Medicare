package controllers

import (
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
	"encoding/json"
	"strings"
)

//Event contains event details
type Event struct {
	PatientID int `json:"PatientID"`
	Event string `json:"Event"`
	Medicine []string `json:"Medicine"`
	TimeSFO int `json:"TimeSFO"`
	Success bool `json:"Success"`
}

//SuccessRate contains the success rate
type SuccessRate struct {
	Event string `json:"Event"`
	Medicine string `json:"Medicine"`
	SuccessProp float64 `json:"Probability"`
}


//SuggestHandler suggests the effective medicine for the disease and the success rate
func SuggestHandler(w http.ResponseWriter, r *http.Request) {

	// prevent CORS error
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	r.ParseForm()

	// patientID := r.FormValue("patientid")
	event := r.FormValue("event")

	jsonFile, err := os.Open("datastore/init.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var data []Event
	json.Unmarshal(byteValue, &data)
	if (err != nil) {
		fmt.Println("err")
	}

	var list []Event

	for i := 0; i < len(data); i++ {
		if (strings.Compare(strings.ToLower(event), strings.ToLower(data[i].Event)) == 0) {
			list = append(list, data[i])
		}
	}


	if (len(list) != 0) {
		var SuccessList []SuccessRate

		for i := 0; i < len(list); i++ {
			count := 0
			for j := 0; j < len(SuccessList); j++ {
				if (strings.Compare(strings.ToLower(list[i].Medicine[0]), strings.ToLower(SuccessList[j].Medicine)) == 0) {

					if (list[i].Success == true) {
						SuccessList[j].SuccessProp = SuccessList[j].SuccessProp + 1/float64(len(SuccessList))
					} else if (list[i].Success == false) {
						SuccessList[j].SuccessProp = SuccessList[j].SuccessProp - 1/float64(len(SuccessList))
					} 
					count++

				}
			}
			if (count == 0) {
				var suclist SuccessRate
				suclist = SuccessRate{
					Event: list[i].Event,
					Medicine: list[i].Medicine[0],
				}	

				if (list[i].Success == true) {
					suclist.SuccessProp = 1/float64(len(list))
				} else if(list[i].Success == false) {
					suclist.SuccessProp = -1/float64(len(list))
				}	
				SuccessList = append(SuccessList, suclist)
			}
		}
		

		j, err := json.Marshal(SuccessList)
		if err != nil {
			panic(err)
		}
		w.Write(j)

	} else {
		w.Write([]byte(`null`))
	}

}