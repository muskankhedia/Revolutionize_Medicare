package main

import (
	"fmt"
	"log"
	"net/http"
	"crypto/sha256"
	"io/ioutil"
	"os"
	"encoding/json"
	"encoding/hex"
	"github.com/gorilla/mux"
	"strings"
	"strconv"
	"github.com/muskankhedia/Revolutionize_Medicare/service/controllers"
)



func init() {
	controllers.PatientIDsMatch = []int{}
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")
}

func error404(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, `404 page not found`)
}

func routeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, `Check`)
}

func init() {
	log.Println("initializing chain")
	jsonFile, err := os.Open("datastore/init.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var data []controllers.Event
	json.Unmarshal(byteValue, &data)
	if err != nil {
		fmt.Println("err")
	}
	var block controllers.EventBlock

	// copying
	for i:=0; i< len(data); i++ {
		block.Event = data[i].Event
		block.PatientID = data[i].PatientID
		block.TimeSFO = data[i].TimeSFO
		block.Success = data[i].Success
		block.Medicine = data[i].Medicine
		controllers.Chain = append(controllers.Chain, block)
	}

	// making hashes
	for i:=0; i< len(data); i++ {
		var data string
		if i == 0 {
			data = controllers.Chain[i].Event + string(controllers.Chain[i].PatientID) + string(controllers.Chain[i].TimeSFO) + strconv.FormatBool(controllers.Chain[i].Success) + strings.Join(controllers.Chain[i].Medicine, ",")
		} else {
			data = controllers.Chain[i-1].Event + string(controllers.Chain[i-1].PatientID) + string(controllers.Chain[i-1].TimeSFO) + strconv.FormatBool(controllers.Chain[i].Success) + strings.Join(controllers.Chain[i].Medicine, ",")
		}
		inst := sha256.New()
		inst.Write([]byte(data))

		controllers.Chain[i].Hash = hex.EncodeToString(inst.Sum(nil))
	}
	log.Println("done chain building")
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/login", controllers.LoginHandler)
	r.HandleFunc("/get_profile", controllers.GetProfileHandler)
	r.HandleFunc("/update_profile", controllers.UpdateProfileHandler)
	r.HandleFunc("/add_data", controllers.AddDataHandler)
	r.HandleFunc("/signup", controllers.SignupHandler).Methods("POST")
	r.HandleFunc("/allevents", controllers.AllEventsHandler)
	r.HandleFunc("/suggestmedicines", controllers.SuggestHandler)
	r.HandleFunc("/update_Success", controllers.UpdateSuccessHandler)

	r.NotFoundHandler = http.HandlerFunc(error404)

	log.Fatal(http.ListenAndServe(":9000", r))
}
