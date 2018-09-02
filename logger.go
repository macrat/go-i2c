package i2c

import (
	"github.com/sirupsen/logrus"
)

var (
	lg = logrus.New()

	PanicLevel = logrus.PanicLevel
	FatalLevel = logrus.FatalLevel
	ErrorLevel = logrus.ErrorLevel
	WarnLevel = logrus.WarnLevel
	InfoLevel = logrus.InfoLevel
	DebugLevel = logrus.DebugLevel
)

func SetLogLevel(level logrus.Level) {
	lg.SetLevel(level)
}
