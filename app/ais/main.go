package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
	"gitlab.com/elcarim-optronic-indonesia/elcas-service-node/config"
	"gitlab.com/elcarim-optronic-indonesia/elcas-service-node/config/network"
	"gitlab.com/elcarim-optronic-indonesia/elcas-service-node/internal/service/ais/serial"
	"gitlab.com/elcarim-optronic-indonesia/elcas-service-node/internal/service/ais/udp"
	"gitlab.com/elcarim-optronic-indonesia/elcas-service-node/pkg"
)

func init() {
	// Load the .env file (once)
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Fail to read .env")
		os.Exit(1)
	}

	pkg.Log(log.InfoLevel, os.Getenv("APPLICATION_NAME")+" "+os.Getenv("APPLICATION_VERSION")+" is running... ")

	// Get config
	if err := config.Get(); err != nil {
		pkg.Log(log.ErrorLevel, "Get config fail!")
		os.Exit(1)
	}
}

func main() {
	networkSettings := network.Get()

	if networkSettings.SourceAis == network.SourceSerial {
		// TODO: Init serial
		if err := serial.Init(); err != nil {
			pkg.Log(log.ErrorLevel, "Init serial fail!")
			os.Exit(2)
		}
		// Run worker
		go serial.Run()
	} else {
		// Init udp server
		if err := udp.Init(); err != nil {
			pkg.Log(log.ErrorLevel, "Init udp server fail!")
			os.Exit(2)
		}

		// Run worker
		go udp.Run()
	}

	// Prevent the main function from exiting
	select {}
}
