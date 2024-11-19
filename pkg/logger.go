package pkg

import (
	_ "embed"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger is a wrapper around the charmbracelet/log.Logger
var (
	logger  *log.Logger
	once    sync.Once
	release bool
)

// logFile is defined at the package level to manage its lifecycle
var LogFile *lumberjack.Logger

// init initializes the logger instance.
func init() {
	// Load the .env file (once)
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Fail to read .env")
		os.Exit(1)
	}

	if os.Getenv("APPLICATION_MODE") == "release" {
		release = true
	} else {
		release = false
	}

	var applicationDir string

	// TODO: must do testing runtime linux (create selection)
	runtime := runtime.GOOS
	switch runtime {
	case "windows":
		applicationDir = os.Getenv("DATA_DIR_WIN")
	case "linux":
		dir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal("Fail to find home directory")
			os.Exit(1)
		}

		applicationDir = dir + "/"
	}
	// Create logs directory if it doesn't exist
	dir := fmt.Sprintf("%s%s", applicationDir, os.Getenv("APPLICATION_NAME"))
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		log.Fatal("Error creating log directory:", err)
		os.Exit(1)
	}

	pathFile := fmt.Sprintf("%s%s/%s", applicationDir, os.Getenv("APPLICATION_NAME"), os.Getenv("FILENAME_LOG"))
	LogFile = &lumberjack.Logger{
		Filename:   pathFile, // Log file name
		MaxSize:    1,        // Maximum size in megabytes before rotation
		MaxBackups: 30,       // Maximum number of backup log files to keep
		MaxAge:     30,       // Maximum number of days to retain old log files
		Compress:   false,    // Compress old log files
	}

	once.Do(func() {
		multiWriter := io.MultiWriter(LogFile, os.Stdout)
		logger = log.New(multiWriter)
		logger.SetFormatter(log.TextFormatter)
		logger.SetReportCaller(false)
		logger.SetReportTimestamp(true)

	})
}

// Log logs messages with the logger
func Log(level log.Level, message string) {
	// Get the current caller's information
	pc, file, line, ok := runtime.Caller(1) // 1 means the caller of this function
	if !ok {
		file = "unknown"
		line = 0
	}

	// Get only the base name of the file
	file = filepath.Base(file) // Extract the base file name

	// Get the function name
	funcName := runtime.FuncForPC(pc).Name()
	funcName = filepath.Base(funcName) // Only keep the function name

	// Log the message with the modified caller information
	if !release {
		switch level {
		case log.InfoLevel:
			logger.Info(fmt.Sprintf("<%s:%s:%d>: %s", file, funcName, line, message))
		case log.WarnLevel:
			logger.Warn(fmt.Sprintf("<%s:%s:%d>: %s", file, funcName, line, message))
		case log.ErrorLevel:
			logger.Error(fmt.Sprintf("<%s:%s:%d>: %s", file, funcName, line, message))
		case log.FatalLevel:
			logger.Fatal(fmt.Sprintf("<%s:%s:%d>: %s", file, funcName, line, message))
		case log.DebugLevel:
			logger.Debug(fmt.Sprintf("<%s:%s:%d>: %s", file, funcName, line, message))
		default:
			logger.Info(fmt.Sprintf("<%s:%s:%d>: %s", file, funcName, line, message)) // Default to info level
		}
	} else {
		switch level {
		case log.InfoLevel:
			logger.Info(fmt.Sprintf(": %s", message))
		case log.WarnLevel:
			logger.Warn(fmt.Sprintf(": %s", message))
		case log.ErrorLevel:
			logger.Error(fmt.Sprintf(": %s", message))
		case log.FatalLevel:
			logger.Fatal(fmt.Sprintf(": %s", message))
		case log.DebugLevel:
			logger.Debug(fmt.Sprintf(": %s", message))
		default:
			logger.Info(fmt.Sprintf(": %s", message)) // Default to info level
		}
	}

}
