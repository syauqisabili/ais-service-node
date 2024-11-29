package serial

import (
	"fmt"
	"strings"

	"github.com/BertoldVdb/go-ais"
	"github.com/BertoldVdb/go-ais/aisnmea"
	"github.com/charmbracelet/log"
	"gitlab.com/elcarim-optronic-indonesia/elcas-service-node/config/network"
	"gitlab.com/elcarim-optronic-indonesia/elcas-service-node/pkg"
	"go.bug.st/serial" // Updated import path

	target "gitlab.com/elcarim-optronic-indonesia/elcas-service-node/internal/service/ais"
)

// Global variable for the serial port
var port serial.Port

func Init() error {
	// Get config instance
	networkCfg := network.Get()

	serialCfg := &serial.Mode{
		BaudRate: int(networkCfg.Serial.Baudrate),
		Parity:   serial.Parity(networkCfg.Serial.Parity),
		DataBits: int(networkCfg.Serial.DataBits),
		StopBits: serial.StopBits(networkCfg.Serial.StopBits),
	}

	var err error
	port, err = serial.Open(networkCfg.Serial.ComPort, serialCfg)
	if err != nil {
		return err
	}

	pkg.Log(log.InfoLevel, fmt.Sprintf("SerialPort Opening on %s...", networkCfg.Serial.ComPort))

	return nil
}

func Run() {
	// Get config instance
	networkCfg := network.Get()
	nm := aisnmea.NMEACodecNew(ais.CodecNew(false, false))

	for {
		buffer := make([]byte, 1024*8) // Buffer to store incoming data
		var n int

		for {
			n, err := port.Read(buffer)
			if err != nil {
				pkg.Log(log.InfoLevel, fmt.Sprintf("Error reading from serial port: %s", err.Error()))
				continue
			}
			if n == 0 {
				break
			}
		}

		message := string(buffer[:n])
		pkg.Log(log.InfoLevel, fmt.Sprintf("(%s): %s", networkCfg.Serial.ComPort, strings.ReplaceAll(message, "\r\n", "")))

		parts := strings.Split(message, "\r\n")
		for _, part := range parts {
			// Remove  0x00 character
			part = strings.ReplaceAll(part, "\x00", "")
			if part == "" {
				continue
			}

			// Parse ais sentence
			decoded, err := nm.ParseSentence(part)
			if err != nil {
				pkg.Log(log.ErrorLevel, err.Error())
				continue
			}

			// Check decoded ais message is null
			if decoded != nil {
				// send ais msg to handler via dispatcher
				err := target.Handler(decoded.Packet)
				if err != nil {
					pkg.Log(log.ErrorLevel, err.Error())
				}
			}
		}
	}
}
