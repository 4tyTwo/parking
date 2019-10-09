package main

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

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
	commander   commander
}

func newParkingLot(capacity int, commander commander) *parkingLot {
	return &parkingLot{
		capacity:    capacity,
		takenPlaces: make(map[string]int, capacity),
		commander:   commander,
	}
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
	statusCode := http.StatusAccepted
	if err != nil {
		statusCode = http.StatusNotFound
	}
	c.String(statusCode, "")
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
	accessLogPath := os.Getenv("accessLogPath")
	accessLog, err := os.OpenFile(accessLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	checkErr(err)
	gin.DefaultWriter = io.MultiWriter(accessLog)
	commonLogPath := os.Getenv("commonLogPath")
	commonLog, err := os.OpenFile(commonLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	checkErr(err)
	log.SetOutput(commonLog)
	bitrate, err := strconv.ParseUint(os.Getenv("bitrate"), 10, 32)
	timeout, err := strconv.ParseUint(os.Getenv("timeout"), 10, 32)
	checkErr(err)
	ps := newParkingServer(
		*newParkingLot(
			15,
			*newCommander(
				os.Getenv("comportDevice"),
				uint(bitrate),
				uint(timeout),
			),
		),
	)
	router := gin.Default()
	router.GET("/parkingLot", ps.getFreePlaces)
	router.GET("/parkingLot/:code", ps.getCar)
	router.POST("/parkingLot", ps.placeCar)
	router.Run(":8080")
}
