package client

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/4tyTwo/parking/server"
	"github.com/4tyTwo/parking/utils"
)

// GetFreePlaces -
func GetFreePlaces(host string) (server.FreePlacesResponse, error) {
	req, err := http.NewRequest("GET", host+"/parkingLot", nil)
	utils.CheckErr(err) // TODO: check for connection refused
	client := &http.Client{}
	resp, err := client.Do(req)
	utils.CheckErr(err)
	var freePlaces server.FreePlacesResponse
	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		utils.CheckErr(err)
		json.Unmarshal(body, &freePlaces)
		return freePlaces, nil
	}
	if resp.StatusCode == 500 {
		return freePlaces, errors.New("Server Failed")
	}
	return freePlaces, errors.New("Unknown Error")
}

// PlaceCar -
func PlaceCar(host string) (server.PlaceCarResponse, error) {
	req, err := http.NewRequest("POST", host+"/parkingLot", nil)
	utils.CheckErr(err) // TODO: check for connection refused
	client := &http.Client{}
	resp, err := client.Do(req)
	utils.CheckErr(err)
	var placedCar server.PlaceCarResponse
	if resp.StatusCode == 202 {
		body, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		utils.CheckErr(err)
		json.Unmarshal(body, &placedCar)
		return placedCar, nil
	}
	if resp.StatusCode == 500 {
		return placedCar, errors.New("Server Failed")
	}
	if resp.StatusCode == 500 {
		return placedCar, errors.New("Server Failed")
	}
	return placedCar, errors.New("Unknown Error")
}

// GetCar -
func GetCar(host, code string) error {
	req, err := http.NewRequest("GET", host+"/parkingLot/"+code, nil)
	utils.CheckErr(err) // TODO: check for connection refused
	client := &http.Client{}
	resp, err := client.Do(req)
	utils.CheckErr(err)
	if resp.StatusCode == 202 {
		return nil
	}
	if resp.StatusCode == 404 {
		return errors.New("No car found")
	}
	if resp.StatusCode == 500 {
		return errors.New("Server Failed")
	}
	return errors.New("Unknown Error")
}
