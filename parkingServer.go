package main

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type errorResponce struct {
	Description string
}

type placeCarResponce struct {
	PlaceCode string
}

type freePlacesResponce struct {
	FreePlaces int
}

type parkingLot struct {
	capacity    int
	takenPlaces map[string]int
}

func (pl parkingLot) getFreePlaces() int {
	return pl.capacity - len(pl.takenPlaces)
}

func (pl *parkingLot) getCar(code string) error {
	if _, exists := pl.takenPlaces[code]; !exists {
		return errors.New("Not Found")
	}
	delete(pl.takenPlaces, code)
	// TODO actually get it
	return nil
}

func (pl parkingLot) generateCode() string {
	for true {
		code := generateRandomUppercaseString(5)
		if _, exists := pl.takenPlaces[code]; !exists {
			return code
		}
	}
	return ""
}

func (pl *parkingLot) placeCar() (string, error) {
	if pl.getFreePlaces() == 0 {
		return "", errors.New("Parking Lot is Full")
	}
	code := pl.generateCode()
	pl.takenPlaces[code] = pl.capacity - pl.getFreePlaces()
	return code, nil
}

type parkingServer struct {
	parkingLot parkingLot
}

func newParkingServer(pl parkingLot) *parkingServer {
	return &parkingServer{parkingLot: pl}
}

func (ps *parkingServer) getFreePlaces(c *gin.Context) {
	freePlaces := ps.parkingLot.getFreePlaces()
	c.JSON(http.StatusOK, freePlacesResponce{FreePlaces: freePlaces})
}

func (ps *parkingServer) getCar(c *gin.Context) {
	code := c.Param("code")
	err := ps.parkingLot.getCar(code)
	var empty struct{}
	statusCode := http.StatusAccepted
	if err != nil {
		statusCode = http.StatusNotFound
	}
	c.JSON(statusCode, empty)
}

func (ps *parkingServer) placeCar(c *gin.Context) {
	code, err := ps.parkingLot.placeCar()
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponce{Description: "No free places"})
	} else {
		c.JSON(http.StatusAccepted, placeCarResponce{PlaceCode: code})
	}
}

func main() {
	err := godotenv.Load("parking.env")
	checkErr(err)

	accessLog, err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	checkErr(err)
	gin.DefaultWriter = io.MultiWriter(accessLog)

	commonLog, err := os.OpenFile("common.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	checkErr(err)
	log.SetOutput(commonLog)
	ps := newParkingServer(parkingLot{capacity: 15, takenPlaces: make(map[string]int, 15)})
	router := gin.Default()
	router.GET("/parkingLot", ps.getFreePlaces)
	router.GET("/parkingLot/:code", ps.getCar)
	router.POST("/parkingLot", ps.placeCar)
	router.Run(":8080")
}
