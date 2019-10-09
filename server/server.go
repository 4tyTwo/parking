package server

import (
	"net/http"

	"github.com/4tyTwo/parking/lot"
	"github.com/gin-gonic/gin"
)

// ParkingServer represents a server with ParkingLot state
type ParkingServer struct {
	parkingLot lot.ParkingLot
}

// New returns ParkingServer assosiated with given ParkingLot
func New(pl lot.ParkingLot) *ParkingServer {
	return &ParkingServer{parkingLot: pl}
}

// GetFreePlaces - gin handler
func (ps *ParkingServer) GetFreePlaces(c *gin.Context) {
	freePlaces := ps.parkingLot.GetFreePlaces()
	c.JSON(http.StatusOK, FreePlacesResponse{FreePlaces: freePlaces})
}

// GetCar - gin handler
func (ps *ParkingServer) GetCar(c *gin.Context) {
	code := c.Param("code")
	err := ps.parkingLot.GetCar(code)
	statusCode := http.StatusAccepted
	if err != nil {
		statusCode = http.StatusNotFound
	}
	c.String(statusCode, "")
}

// PlaceCar - gin handler
func (ps *ParkingServer) PlaceCar(c *gin.Context) {
	code, err := ps.parkingLot.PlaceCar()
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Description: "No free places"})
	} else {
		c.JSON(http.StatusAccepted, PlaceCarResponse{PlaceCode: code})
	}
}
