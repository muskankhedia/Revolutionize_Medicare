package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

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
func Learning(x []PatientIDsMatch) {
	fmt.Println("deep patientsmatch")
	fmt.Println(x)
	profiles := getMatchingProfiles(x)
	fmt.Println("matching profiles below")
	fmt.Println(profiles)

	var data train.Examples
	for _, inst := range profiles {

		var success bool
		var typer float64
		for _, x := range PatientMatch {
			if x.ID == inst.PatientID {
				success = x.Success
				fmt.Println("this (((((((((((((((((")
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

	net := dep.NewNeural(&dep.Config{
		Inputs:     8,
		Layout:     []int{8, 8, 8, 1},
		Activation: dep.ActivationLinear,
		Mode:       dep.ModeBinary,
		Weight:     dep.NewNormal(0.1, 0.1),
		Bias:       true,
	})

	optimizer := train.NewSGD(0.05, 0.1, 1e-6, true)
	// params: optimizer, verbosity (print stats at every 50th iteration)
	trainer := train.NewTrainer(optimizer, 50)

	trains, heldout := data.Split(0.5)

	trainer.Train(net, trains, heldout, 10000)
	fmt.Println("below output")
	fmt.Println(net.Predict([]float64{80, 120, 12, 80, 22, 97, 72, 18}))
}
