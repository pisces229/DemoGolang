package singleton

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var AppLogrus = initLogrus()
var AppLogwriter = initLogWriter()

func initLogrus() (logger *logrus.Logger) {
	fmt.Println("initLogrus")
	logger = logrus.New()
	logger.SetOutput(os.Stdout)
	// fmt.Println("Flag:", os.O_WRONLY|os.O_CREATE|os.O_APPEND)
	// fmt.Println("FileMode:", os.ModeAppend)
	// logFile, _ := os.OpenFile(AppConfiguration.LogPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	// logger.SetOutput(logFile)
	logger.SetFormatter(&logrus.JSONFormatter{})
	// logger.SetFormatter(&logrus.TextFormatter{})
	logger.SetLevel(logrus.InfoLevel)
	logger.SetReportCaller(true)
	logger.Info("[NewLogrus]")
	return
}

type LogWriter struct {
	logger *logrus.Logger
}

func (logWriter *LogWriter) Printf(format string, v ...interface{}) {
	logstr := fmt.Sprintf(format, v...)
	logWriter.logger.Info(logstr)
}

func initLogWriter() *LogWriter {
	fmt.Println("initLogWriter")
	return &LogWriter{logger: AppLogrus}
}
