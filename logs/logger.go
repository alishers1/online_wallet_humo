package logs

import (
	"os"

	"github.com/sirupsen/logrus"
)

func InitLogger() error {
	file, err := os.OpenFile("../../logs/logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	} else {
		logrus.SetOutput(file)
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.SetLevel(logrus.InfoLevel)
	return nil
}
