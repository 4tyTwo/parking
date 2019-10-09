package main

import (
	"io"
	"log"
	"os"
	"strconv"

	"github.com/4tyTwo/parking/commander"
	"github.com/4tyTwo/parking/lot"
	"github.com/4tyTwo/parking/server"
	"github.com/4tyTwo/parking/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("parking.env")
	utils.CheckErr(err)
	accessLogPath := os.Getenv("accessLogPath")
	accessLog, err := os.OpenFile(accessLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	utils.CheckErr(err)
	gin.DefaultWriter = io.MultiWriter(accessLog)
	commonLogPath := os.Getenv("commonLogPath")
	commonLog, err := os.OpenFile(commonLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	utils.CheckErr(err)
	log.SetOutput(commonLog)
	bitrate, err := strconv.ParseUint(os.Getenv("bitrate"), 10, 32)
	timeout, err := strconv.ParseUint(os.Getenv("timeout"), 10, 32)
	utils.CheckErr(err)
	ps := server.New(
		*lot.New(
			15,
			*commander.New(
				os.Getenv("comportDevice"),
				uint(bitrate),
				uint(timeout),
			),
		),
	)
	router := gin.Default()
	router.GET("/parkingLot", ps.GetFreePlaces)
	router.GET("/parkingLot/:code", ps.GetCar)
	router.POST("/parkingLot", ps.PlaceCar)
	router.Run(":8080")
}
