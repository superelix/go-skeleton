package common

import (
	"io"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logInit sync.Once
	logger  *logrus.Logger
)

const (
	maxSize    = 10 // Max size in megabytes before rotation
	maxBackups = 3  // Maximum number of old log files to retain
	maxAge     = 7  // Maximum number of days to retain old log files
)

func initializeLogger() {

	if _, err := os.Stat(ProjectRootPath() + "logs"); os.IsNotExist(err) {
		if err = os.MkdirAll(ProjectRootPath()+"logs", os.ModePerm); err != nil {
			panic(err)
		}
	}

	relativeLogFilePath := "logs/" + "application.log"
	logRotation := &lumberjack.Logger{
		Filename:   ProjectRootPath() + relativeLogFilePath,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   false,
	}

	logger = logrus.New()

	logger.SetOutput(io.MultiWriter(os.Stdout, logRotation))
	logger.SetLevel(logrus.InfoLevel)
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.JSONFormatter{})
}

func GetLogger() *logrus.Logger {
	logInit.Do(func() {
		initializeLogger()
	})
	return logger
}
