package controllers

import (
	dep "github.com/patrikeh/go-deep"
	train "github.com/patrikeh/go-deep/training"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type dataStr struct {
	nodeInput []float64
	output []float64
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

func getMatchingProfiles(arr []int) (matchingProfiles []PatientProfile) {
	profiles := loadJSONProfile()
	for _, i := range arr {
		for _, j := range profiles {
			if i == j.PatientID {
				matchingProfiles = append(matchingProfiles, j)
			}
		}
	}
	return
}

//Learning global learning func
func Learning(x []int) {
	fmt.Println("deep patientsmatch")
	fmt.Println(x)
	profiles := getMatchingProfiles(x)
	fmt.Println("matching profiles below")
	fmt.Println(profiles)
	
	var data []train.Example

	for _, inst := range profiles {
		i := train.Example{
			Input: []float64{
				inst.Bmi,
				inst.BpD,
				inst.BpS,
				inst.PatientID,
				inst.Pulse,
				inst.Resp,
				inst.SugarAF,
				inst.SugarBF,
				inst.Temp,
			}, Response: []float64{
				0,
			}
			
		}
	}

	net := dep.NewNeural(&dep.Config{
		Inputs: 9,
		Layout: []int{2,2,2,1},
		Activation: dep.ActivationLinear,
		Mode: dep.ModeBinary,
		Weight: dep.NewNormal(0.1, 0.1),
		Bias: true,
	})

	optimizer := training.NewSGD(0.05, 0.1, 1e-6, true)
	// params: optimizer, verbosity (print stats at every 50th iteration)
	trainer := training.NewTrainer(optimizer, 50)

}
