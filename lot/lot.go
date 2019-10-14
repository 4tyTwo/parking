package lot

import (
	"errors"
	"strconv"
	"strings"

	"github.com/4tyTwo/parking/commander"
	"github.com/4tyTwo/parking/utils"
)

// ParkingLot represents a parking lot structure
type ParkingLot struct {
	capacity    int
	takenPlaces map[string]int
	commander   commander.Commander
}

// New returns ParkingLot with given capacity and commander
func New(capacity int, commander commander.Commander) *ParkingLot {
	return &ParkingLot{
		capacity:    capacity,
		takenPlaces: make(map[string]int, capacity),
		commander:   commander,
	}
}

// GetFreePlaces returns the amount free places at ParkingLot
func (pl ParkingLot) GetFreePlaces() int {
	return pl.capacity - len(pl.takenPlaces)
}

// GetCar proceeds car retrieval procedure if there is a car assosiated with given code
func (pl *ParkingLot) GetCar(code string) error {
	n, exists := pl.takenPlaces[code]
	if !exists {
		return errors.New("Not Found")
	}
	delete(pl.takenPlaces, code)
	err := pl.commander.WriteCommand("get " + strconv.Itoa(n) + "\n") // Stringbuilder perhaps?
	utils.CheckErr(err)
	return nil
}

// PlaceCar proceeds car enplacement procedure if there is any free places
func (pl *ParkingLot) PlaceCar() (string, error) {
	freePlaces := pl.GetFreePlaces()
	if freePlaces == 0 {
		return "", errors.New("Parking Lot is Full")
	}
	code := pl.generateCode()
	place := pl.capacity - freePlaces
	err := pl.commander.WriteCommand("put " + strconv.Itoa(place) + "\n") // Stringbuilder perhaps?
	utils.CheckErr(err)
	pl.takenPlaces[code] = place
	return code, nil
}

func (pl ParkingLot) generateCode() string {
	for true {
		// supposingly parking lot will never have capacity > 36^5
		code := strings.ToUpper(utils.GenerateRandomString(5))
		if _, exists := pl.takenPlaces[code]; !exists {
			return code
		}
	}
	return ""
}
