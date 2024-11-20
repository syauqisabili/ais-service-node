package config

import (
	"gitlab.com/elcarim-optronic-indonesia/elcas-service-node/config/network"
	"gitlab.com/elcarim-optronic-indonesia/elcas-service-node/pkg"

	"fmt"
	"os"
	"runtime"
	"strconv"

	"github.com/charmbracelet/log"
	"gopkg.in/ini.v1"
)

func getDir() (string, error) {
	var applicationDir string
	runtime := runtime.GOOS
	switch runtime {
	case "windows":
		applicationDir = os.Getenv("DATA_DIR_WIN")
	case "linux":
		dir, err := os.UserHomeDir()
		if err != nil {
			return "", pkg.ErrorStatus(pkg.ErrDirNotExist, "Fail to find home directory")
		}

		applicationDir = dir + "/"
	}

	return applicationDir, nil
}

func setDefaultValue() error {
	pkg.Log(log.InfoLevel, "Set default config... ")
	var port int64
	var _ error

	config := network.Get()

	// grpc server
	config.GrpcServer.Ip = os.Getenv("DEFAULT_GRPC_SERVER_URI")

	port, _ = strconv.ParseInt(os.Getenv("DEFAULT_GRPC_SERVER_PORT"), 10, 16)
	config.GrpcServer.Port = uint16(port)

	// default udp
	config.SourceAis = network.SourceUdpMulticast

	// Udp server
	config.UdpNet.Ip = os.Getenv("DEFAULT_UDP_MULTICAST_URI")
	port, _ = strconv.ParseInt(os.Getenv("DEFAULT_UDP_MULTICAST_PORT"), 10, 16)
	config.UdpNet.Port = uint16(port)

	//Serial
	config.Serial.ComPort = os.Getenv("DEFAULT_COM_PORT")
	config.Serial.Baudrate = network.Baudrate9600
	config.Serial.DataBits = network.DataBits8
	config.Serial.StopBits = network.StopBits1
	config.Serial.Parity = network.NoneParity
	config.Serial.FlowControl = network.NoneFlow

	// Debug
	config.DebugMode = network.DisableDebug

	network.Set(config)

	return nil
}

func show() error {
	pkg.Log(log.InfoLevel, "Show config... ")
	// Get application dir
	applicationDir, err := getDir()
	if err != nil {
		return err
	}

	// Check config.ini if it doesn't exist
	pathFile := fmt.Sprintf("%s%s/%s", applicationDir, os.Getenv("APPLICATION_NAME"), os.Getenv("FILENAME_CONFIG"))
	if _, err := os.Stat(pathFile); err != nil {
		return pkg.ErrorStatus(pkg.ErrFileNotExist, fmt.Sprintf("%s does not exist!", pathFile))
	}

	// Read config.ini
	config, err := ini.Load(pathFile)
	if err != nil {
		return pkg.ErrorStatus(pkg.ErrReadFile, fmt.Sprintf("Fail to read %s ", pathFile))
	}

	// Iterate over all sections
	for _, section := range config.Sections() {

		if section.Name() == "DEFAULT" {
			continue
		}
		pkg.Log(log.InfoLevel, fmt.Sprintf("[%s]", section.Name()))

		// Iterate over all keys in the current section
		for _, key := range section.Keys() {
			pkg.Log(log.InfoLevel, fmt.Sprintf(" %s = %s", key.Name(), key.Value()))
		}
	}

	return nil
}

func write() error {
	pkg.Log(log.InfoLevel, "Write config... ")

	// Param
	var param string

	// Get application dir
	applicationDir, err := getDir()
	if err != nil {
		return err
	}

	// Check config.ini if it doesn't exist
	pathFile := fmt.Sprintf("%s%s/%s", applicationDir, os.Getenv("APPLICATION_NAME"), os.Getenv("FILENAME_CONFIG"))
	if _, err := os.Stat(pathFile); err != nil {
		return pkg.ErrorStatus(pkg.ErrFileNotExist, fmt.Sprintf("%s does not exist!", pathFile))
	}

	// Read config.ini
	settings, err := ini.Load(pathFile)
	if err != nil {
		return pkg.ErrorStatus(pkg.ErrReadFile, fmt.Sprintf("Fail to read %s ", pathFile))
	}

	// Get config instance
	config := network.Get()

	// Grpc server section
	sec, err := settings.NewSection("grpc_server")
	if err != nil {
		return pkg.ErrorStatus(pkg.ErrWriteFile, fmt.Sprintf("Fail to write to %s ", pathFile))
	}
	_, err = sec.NewKey("ip", config.GrpcServer.Ip)
	if err != nil {
		return pkg.ErrorStatus(pkg.ErrWriteFile, fmt.Sprintf("Fail to write to %s ", pathFile))
	}
	_, err = sec.NewKey("port", strconv.FormatUint(uint64(config.GrpcServer.Port), 10))
	if err != nil {
		return pkg.ErrorStatus(pkg.ErrWriteFile, fmt.Sprintf("Fail to write to %s ", pathFile))
	}

	// Ais section
	sec, err = settings.NewSection("ais")
	if err != nil {
		return pkg.ErrorStatus(pkg.ErrWriteFile, fmt.Sprintf("Fail to write to %s ", pathFile))
	}
	switch config.SourceAis {
	case network.SourceSerial:
		param = "serial"
	default:
		param = "udp"
	}
	_, err = sec.NewKey("source", param)
	if err != nil {
		return pkg.ErrorStatus(pkg.ErrWriteFile, fmt.Sprintf("Fail to write to %s ", pathFile))
	}

	// Udp server section
	sec, err = settings.NewSection("udp_server")
	if err != nil {
		return pkg.ErrorStatus(pkg.ErrWriteFile, fmt.Sprintf("Fail to write to %s ", pathFile))
	}
	_, err = sec.NewKey("ip", config.UdpNet.Ip)
	if err != nil {
		return pkg.ErrorStatus(pkg.ErrWriteFile, fmt.Sprintf("Fail to write to %s ", pathFile))
	}
	_, err = sec.NewKey("port", strconv.FormatUint(uint64(config.UdpNet.Port), 10))
	if err != nil {
		return pkg.ErrorStatus(pkg.ErrWriteFile, fmt.Sprintf("Fail to write to %s ", pathFile))
	}

	// Serial section
	sec, err = settings.NewSection("serial")
	if err != nil {
		return pkg.ErrorStatus(pkg.ErrWriteFile, fmt.Sprintf("Fail to write to %s ", pathFile))
	}
	_, err = sec.NewKey("port", config.Serial.ComPort)
	if err != nil {
		return pkg.ErrorStatus(pkg.ErrWriteFile, fmt.Sprintf("Fail to write to %s ", pathFile))
	}
	_, err = sec.NewKey("baudrate", strconv.FormatUint(uint64(config.Serial.Baudrate), 10))
	if err != nil {
		return pkg.ErrorStatus(pkg.ErrWriteFile, fmt.Sprintf("Fail to write to %s ", pathFile))
	}
	switch config.Serial.Parity {
	case network.NoneParity:
		param = "none"
	case network.OddParity:
		param = "odd"
	case network.EvenParity:
		param = "even"
	case network.MarkParity:
		param = "mark"
	case network.SpaceParity:
		param = "space"
	default:
		param = "none"
	}
	_, err = sec.NewKey("parity", param)
	if err != nil {
		return pkg.ErrorStatus(pkg.ErrWriteFile, fmt.Sprintf("Fail to write to %s ", pathFile))
	}
	_, err = sec.NewKey("databits", strconv.FormatUint(uint64(config.Serial.DataBits), 10))
	if err != nil {
		return pkg.ErrorStatus(pkg.ErrWriteFile, fmt.Sprintf("Fail to write to %s ", pathFile))
	}
	_, err = sec.NewKey("stopbits", strconv.FormatUint(uint64(config.Serial.StopBits), 10))
	if err != nil {
		return pkg.ErrorStatus(pkg.ErrWriteFile, fmt.Sprintf("Fail to write to %s ", pathFile))
	}
	switch config.Serial.FlowControl {
	case network.NoneFlow:
		param = "none"
	case network.RtsCts:
		param = "rts/cts"
	case network.DtrDsr:
		param = "dtr/dsr"
	case network.Rs485Rts:
		param = "rs485-rts"
	default:
		param = "none"
	}
	_, err = sec.NewKey("flow_control", param)
	if err != nil {
		return pkg.ErrorStatus(pkg.ErrWriteFile, fmt.Sprintf("Fail to write to %s ", pathFile))
	}

	// Save to file
	err = settings.SaveTo(pathFile)
	if err != nil {
		return pkg.ErrorStatus(pkg.ErrSaveFile, fmt.Sprintf("Fail to save to %s ", pathFile))
	}

	return nil
}

func read() error {
	pkg.Log(log.InfoLevel, "Read config... ")

	// Param
	var port uint64
	var _ error

	// Get application dir
	applicationDir, err := getDir()
	if err != nil {
		return err
	}

	// Check config.ini if it doesn't exist
	pathFile := fmt.Sprintf("%s%s/%s", applicationDir, os.Getenv("APPLICATION_NAME"), os.Getenv("FILENAME_CONFIG"))
	if _, err := os.Stat(pathFile); err != nil {
		return pkg.ErrorStatus(pkg.ErrFileNotExist, fmt.Sprintf("%s does not exist!", pathFile))
	}

	// Read config.ini
	settings, err := ini.Load(pathFile)
	if err != nil {
		return pkg.ErrorStatus(pkg.ErrReadFile, fmt.Sprintf("Fail to read %s ", pathFile))
	}

	// Get config instance
	config := network.Get()

	// Grpc server section
	section := settings.Section("grpc_server")
	config.GrpcServer.Ip = section.Key("ip").String()
	port, _ = section.Key("port").Uint64()
	config.GrpcServer.Port = uint16(port)

	// Source ais
	section = settings.Section("ais")
	switch section.Key("source").String() {
	case "serial":
		config.SourceAis = network.SourceSerial
	default:
		config.SourceAis = network.SourceUdpMulticast
	}

	// Udp server
	section = settings.Section("udp_server")
	config.UdpNet.Ip = section.Key("ip").String()
	port, _ = section.Key("port").Uint64()
	config.UdpNet.Port = uint16(port)

	// Serial
	section = settings.Section("serial")
	config.Serial.ComPort = section.Key("port").String()
	baudrate, _ := strconv.ParseUint(section.Key("baudrate").String(), 10, 32)
	config.Serial.Baudrate = network.Baudrate(uint32(baudrate))

	switch section.Key("parity").String() {
	case "none":
		config.Serial.Parity = network.NoneParity
	case "odd":
		config.Serial.Parity = network.OddParity
	case "even":
		config.Serial.Parity = network.EvenParity
	case "mark":
		config.Serial.Parity = network.MarkParity
	case "space":
		config.Serial.Parity = network.SpaceParity
	default:
		config.Serial.Parity = network.NoneParity
	}

	databits, _ := strconv.ParseUint(section.Key("databits").String(), 10, 32)
	config.Serial.DataBits = network.DataBits(uint8(databits))
	stopbits, _ := strconv.ParseUint(section.Key("stopbits").String(), 10, 32)
	config.Serial.StopBits = network.StopBits(uint8(stopbits))

	switch section.Key("flow_control").String() {
	case "none":
		config.Serial.FlowControl = network.NoneFlow
	case "rts/cts":
		config.Serial.FlowControl = network.RtsCts
	case "dtr/dsr":
		config.Serial.FlowControl = network.DtrDsr
	case "rs485-rts":
		config.Serial.FlowControl = network.Rs485Rts
	default:
		config.Serial.FlowControl = network.NoneFlow
	}

	// Set to handler
	network.Set(config)

	return nil
}

func Get() error {
	pkg.Log(log.InfoLevel, "Get config... ")

	// Get application dir
	applicationDir, err := getDir()
	if err != nil {
		return err
	}

	// Create config.ini if it doesn't exist
	pathFile := fmt.Sprintf("%s%s/%s", applicationDir, os.Getenv("APPLICATION_NAME"), os.Getenv("FILENAME_CONFIG"))
	if _, err := os.Stat(pathFile); err != nil {
		// File does not exist, so create it.
		file, err := os.Create(pathFile)
		if err != nil {
			return pkg.ErrorStatus(pkg.ErrCreateFile, fmt.Sprintf("Fail to create %s", pathFile))
		}
		defer file.Close()

		// Set default value param
		if err := setDefaultValue(); err != nil {
			return pkg.ErrorStatus(pkg.ErrNoParam, fmt.Sprintf("Fail to set default config: %s", err.Error()))
		}

		// Write config to file
		if err := write(); err != nil {
			return err
		}
	}

	// Read config file
	if err := read(); err != nil {
		return err
	}

	// Show config to console
	if err := show(); err != nil {
		return err
	}
	return nil
}
