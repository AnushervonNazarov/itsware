package logger

import (
	"log"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

// SetLogger Installing Logger
var (
	Error *log.Logger
)

const (
	LogError      = "logs/error.log"
	LogMaxSize    = 25
	LogMaxBackups = 5
	LogMaxAge     = 30
	LogCompress   = true
)

func Init() error {
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		err = os.Mkdir("logs", 0755)
		if err != nil {
			return err
		}
	}

	fileError, err := os.OpenFile(LogError, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	Error = log.New(fileError, "", log.Ldate|log.Lmicroseconds)

	lumberLogError := &lumberjack.Logger{
		Filename:   LogError,
		MaxSize:    LogMaxSize, // megabytes
		MaxBackups: LogMaxBackups,
		MaxAge:     LogMaxAge,   //days
		Compress:   LogCompress, // disabled by default
		LocalTime:  true,
	}

	Error.SetOutput(lumberLogError)

	return nil
}
