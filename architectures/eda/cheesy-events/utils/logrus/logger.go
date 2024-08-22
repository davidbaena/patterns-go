package logrus

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: false,
	})
}

func NewLogger(module string) *logrus.Entry {
	return log.WithField("Module", module)
}
