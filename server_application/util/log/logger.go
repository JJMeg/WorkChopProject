package log

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

type Config struct {
	Output string `json:"output"`
	Level  string `json:"level"`
	Format string `json:"format"`
}

func Newlogger(cfg *Config) (*logrus.Logger, error) {
	level, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		return nil, err
	}
	logger := logrus.New()
	logger.Level = level

	var output io.Writer
	switch cfg.Output {
	case "stdout":
		output = os.Stdout
	case "stderr":
		output = os.Stderr
	}

	logger.SetOutput(output)

	return logger, nil
}
