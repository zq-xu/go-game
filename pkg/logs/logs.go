package logs

import (
	"github.com/rotisserie/eris"
	"github.com/sirupsen/logrus"
)

func InitLog(level string) error {
	l, err := logrus.ParseLevel(level)
	if err != nil {
		return eris.Wrapf(err, "invaild log level %s", err)
	}

	logrus.SetLevel(l)
	logrus.SetFormatter(&logrus.JSONFormatter{DisableHTMLEscape: true})
	return nil
}
