package main

import (
	"io"
	"log"

	"github.com/jacobsa/go-serial/serial"
)

type commander struct {
	device io.ReadWriteCloser
	buffer int
}

func newCommander(device string, bitrate, timeout uint) *commander {
	opts := serial.OpenOptions{
		PortName:              device,
		BaudRate:              bitrate,
		InterCharacterTimeout: timeout, //ms
	}
	port, err := serial.Open(opts)
	checkErr(err)
	return &commander{
		device: port,
		buffer: 64,
	}
}

func (comm *commander) writeCommand(command string) error {
	bytes := []byte(command)
	comm.device.Write(bytes)
	resp := make([]byte, 10)
	_, err := comm.device.Read(resp)
	log.Print(string(resp))
	return err
}
