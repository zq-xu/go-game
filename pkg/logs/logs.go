package logs

import (
	"log"
	"os"

	"github.com/rotisserie/eris"
	"github.com/sirupsen/logrus"
)

const (
	defaultLogrusLogLevel = logrus.InfoLevel
)

var Logger = logrus.New()

// InitLogger
func InitLogger(level string) {
	Logger.SetOutput(os.Stdout)

	Logger.SetFormatter(&logrus.JSONFormatter{DisableHTMLEscape: true})
	l, err := getLogrusLevel(level)
	if err != nil {
		log.Fatal("log init failed", err)
	}
	Logger.SetLevel(l)

	Logger.Infof("Succeed to init log with level %s", Logger.Level.String())
}

func getLogrusLevel(str string) (logrus.Level, error) {
	if str == "" {
		return defaultLogrusLogLevel, nil
	}

	l, err := logrus.ParseLevel(str)
	if err != nil {
		return defaultLogrusLogLevel, eris.Wrapf(err, "failed to parse logrus level from %s", str)

	}

	return l, nil
}
