package udp

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/BertoldVdb/go-ais"
	"github.com/BertoldVdb/go-ais/aisnmea"
	"github.com/charmbracelet/log"
	"gitlab.com/elcarim-optronic-indonesia/elcas-service-node/config/network"
	target "gitlab.com/elcarim-optronic-indonesia/elcas-service-node/internal/service/ais"
	"gitlab.com/elcarim-optronic-indonesia/elcas-service-node/pkg"
)

var (
	conn *net.UDPConn
)

func Init() error {
	// Get config instance
	config := network.Get()

	// Resolve UDP address
	port := ":" + strconv.FormatUint(uint64(config.UdpNet.Port), 10)
	udpAddress, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		pkg.Log(log.FatalLevel, err.Error())
		return err
	}

	// Create UDP connection
	conn, err = net.ListenUDP("udp", udpAddress)
	if err != nil {
		pkg.Log(log.FatalLevel, err.Error())
		return err
	}

	pkg.Log(log.InfoLevel, fmt.Sprintf("UDP Listening on %s...", port))

	return nil
}

func Run() {
	nm := aisnmea.NMEACodecNew(ais.CodecNew(false, false))
	for {
		buffer := make([]byte, 1024*8) // Buffer to hold incoming data

		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			pkg.Log(log.InfoLevel, fmt.Sprintf("Error reading from udp: %s", err.Error()))
			continue
		}

		// Process the incoming data
		message := string(buffer[:n])
		pkg.Log(log.InfoLevel, fmt.Sprintf("(%s): %s", remoteAddr, strings.ReplaceAll(message, "\r\n", "")))

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
				if err := target.Handler(decoded.Packet); err != nil {
					pkg.Log(log.ErrorLevel, err.Error())
				}
			}
		}
	}
}
