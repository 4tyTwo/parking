package commander

import (
	"io"
	"log"

	"github.com/4tyTwo/parking/utils"

	"github.com/jacobsa/go-serial/serial"
)

// Commander organizing io with serial device
type Commander struct {
	device io.ReadWriteCloser
	buffer int
}

// New returns created commander struct with given
func New(device string, bitrate, timeout uint) *Commander {
	opts := serial.OpenOptions{
		PortName:              device,
		BaudRate:              bitrate,
		DataBits:              8,
		StopBits:              1,
		MinimumReadSize:       4,
		InterCharacterTimeout: 200,
	}
	port, err := serial.Open(opts)
	utils.CheckErr(err)
	return &Commander{
		device: port,
		buffer: 64,
	}
}

func (comm *Commander) WriteCommand(command string) error {
	bytes := []byte(command)
	n, err := comm.device.Write(bytes)
	if n != len(bytes) {
		log.Fatal("Wrote wrong number of bytes")
	}
	resp := make([]byte, 10)
	n, err = comm.device.Read(resp)
	if err != nil {
		log.Fatalf("port.Read: %v", err)
	}
	if string(resp[:n]) != "OK\r\n" {
		log.Fatal("Device responded with not OK")
	}
	return err
}
