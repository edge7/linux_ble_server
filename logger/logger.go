package ed7_logger

import (
	"io"
	logging "log"
	"os"

	logrus "github.com/sirupsen/logrus"
)

var (
	log *logrus.Logger
)

func init() {
	f, err := os.OpenFile("ble_server.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		logging.Fatalf("error opening file: %v", err)
	}

	log = logrus.New()

	log.SetOutput(os.Stderr)
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "Mon Jan 2 15:04:05 MST 2006"
	log.SetFormatter(customFormatter)
	customFormatter.FullTimestamp = true

	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)
}

// Info ...
func Info(format string, v ...interface{}) {
	log.Infof(format, v...)
}

// Warn ...
func Warn(format string, v ...interface{}) {
	log.Warnf(format, v...)
}

// Error ...
func Error(format string, v ...interface{}) {
	log.Errorf(format, v...)
}
