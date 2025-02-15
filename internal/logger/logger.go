package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func NewLogger(level, file string) (*logrus.Logger, error) {
	log := logrus.New()
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		return nil, err
	}
	log.SetLevel(lvl)

	f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	log.SetOutput(f)

	return log, nil
}
