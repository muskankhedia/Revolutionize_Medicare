package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"encoding/hex"
	"strings"
	"crypto/sha256"
)

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

	var msg EventBlock
	msg = EventBlock{
		EventID: len(Chain),
		PatientID: patientID,
		Event:     event,
		Medicine:  medicine,
		TimeSFO:   timeSinceFirstOccurence,
		Success:   success,
		Hash: "",
	}
	i := len(Chain) - 1
	data := Chain[i].Event + string(Chain[i].PatientID) + string(Chain[i].TimeSFO) + strconv.FormatBool(Chain[i].Success) + strings.Join(Chain[i].Medicine, ",")
	inst := sha256.New()
	inst.Write([]byte(data))
	msg.Hash = hex.EncodeToString(inst.Sum(nil))

	Chain = append(Chain, msg)
	result, err := json.Marshal(Chain)
	err = ioutil.WriteFile("datastore/events.json", result, 0777)
	if err != nil {
		panic(err)
	}

	w.Write([]byte(`true`))
}
