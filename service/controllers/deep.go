package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"os"
	"strconv"
	// "encoding/json"

	dep "github.com/patrikeh/go-deep"
	train "github.com/patrikeh/go-deep/training"
)

type dataStr struct {
	nodeInput []float64
	output    []float64
}

func loadJSONProfile() []PatientProfile {
	// var inst PatientProfile
	// appending the new user to the list of users
	jsonByteValue, err := ioutil.ReadFile("datastore/profile.json")
	if err != nil {
		panic(err)
	}
	var patients []PatientProfile
	if len(jsonByteValue) != 0 {
		err = json.Unmarshal(jsonByteValue, &patients)
	}
	return patients

}

func getMatchingProfiles(arr []PatientIDsMatch) (matchingProfiles []PatientProfile) {
	profiles := loadJSONProfile()
	for _, i := range arr {
		for _, j := range profiles {
			if i.ID == j.PatientID {
				matchingProfiles = append(matchingProfiles, j)
			}
		}
	}
	return
}

//Learning global learning func
func Learning(x []PatientIDsMatch, medicine string, pid int) float64 {
	fmt.Println("deep patientsmatch")
	fmt.Println(x)

	fmt.Println("Medicine:::", medicine)

	var medicineList []PatientIDsMatch

	for i := 0; i < len(x); i++ {
		if strings.Compare(strings.ToLower(medicine), strings.ToLower(x[i].Medicine)) == 0 {
			medicineList = append(medicineList, x[i])
		}
	}

	profiles := getMatchingProfiles(medicineList)
	fmt.Println("matching profiles below")
	fmt.Println(profiles)

	var data train.Examples
	for _, inst := range profiles {

		var success bool
		var typer float64
		for _, x := range PatientMatch {
			if x.ID == inst.PatientID {
				success = x.Success
				break
			}
		}
		if success {
			typer = 1
		} else {
			typer = 0
		}
		i := train.Example{
			Input: []float64{
				inst.Bmi,
				inst.BpD,
				inst.BpS,
				inst.Pulse,
				inst.Resp,
				inst.SugarAF,
				inst.SugarBF,
				inst.Temp,
			}, Response: []float64{
				typer,
			},
		}
		data = append(data, i)
	}

	jsonFile, err := os.Open("datastore/profile.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var patientdataArr []PatientProfile
	var pD PatientProfile
	json.Unmarshal(byteValue, &patientdataArr)
	if err != nil {
		fmt.Println("err")
	}

	for i := 0; i < len(patientdataArr); i++ {

		if pid == patientdataArr[i].PatientID {
			pD = patientdataArr[i]
			break
		}
	}


	net := dep.NewNeural(&dep.Config{
		Inputs:     8,
		Layout:     []int{8, 8, 8, 1},
		Activation: dep.ActivationLinear,
		Mode:       dep.ModeBinary,
		Weight:     dep.NewNormal(0.1, 0.1),
		Bias:       true,
	})

	optimizer := train.NewSGD(0.05, 0.1, 1e-6, true)
	trainer := train.NewTrainer(optimizer, 50)

	trains, heldout := data.Split(0.5)

	trainer.Train(net, trains, heldout, 10000)
	fmt.Println("below output")
	// fmt.Println(net.Predict([]float64{80, 120, 12, 80, 22, 97, 72, 18}))
	var deepRes []float64
	deepRes = net.Predict([]float64{pD.Bmi, pD.BpD, pD.BpS, pD.Pulse, pD.Resp, pD.SugarAF, pD.SugarBF, pD.Temp})
	fmt.Println(deepRes[0])
	i := fmt.Sprintf("%.2f", deepRes[0])
	value, _ := strconv.ParseFloat(i, 64)
	return value

}
