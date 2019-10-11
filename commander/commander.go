package commander

import (
	"io"

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
		DataBits:              8,       // TODO: configure
		StopBits:              1,       // TODO: configure
		InterCharacterTimeout: timeout, //ms
	}
	port, err := serial.Open(opts)
	utils.CheckErr(err)
	return &Commander{
		device: port,
		buffer: 64,
	}
}

func (comm *Commander) writeCommand(command string) error {
	bytes := []byte(command)
	_, err := comm.device.Write(bytes)
	return err
}
