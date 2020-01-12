package log

import (
	"os"
	"strings"

	"github.com/VIBHOR94/golang-microservices/src/api/config"
	"github.com/sirupsen/logrus"
)

var (
	Log *logrus.Logger
)

func init() {
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		level = logrus.DebugLevel
	}

	Log = &logrus.Logger{
		Level: level,
		Out:   os.Stdout,
	}
	if config.IsProduction() {
		Log.Formatter = &logrus.JSONFormatter{}
	} else {
		Log.Formatter = &logrus.TextFormatter{}
	}
}

func Info(msg string, tags ...string) {
	if Log.Level < logrus.InfoLevel {
		return
	}
	Log.WithFields(parseFeilds(tags...)).Info(msg)
}

func Debug(msg string, tags ...string) {
	if Log.Level < logrus.DebugLevel {
		return
	}
	Log.WithFields(parseFeilds(tags...)).Debug(msg)
}

func Error(msg string, tags ...string) {
	if Log.Level < logrus.ErrorLevel {
		return
	}
	Log.WithFields(parseFeilds(tags...)).Error(msg)
}

func parseFeilds(tags ...string) logrus.Fields {
	result := make(logrus.Fields, len(tags))
	for _, tag := range tags {
		els := strings.Split(tag, ":")
		result[strings.TrimSpace(els[0])] = els[1]
	}
	return result
}
