package utils

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strings"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)
	log.SetReportCaller(true)
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		//DisableColors: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf("[ %s:%d ]:", f.File, f.Line)
		},
	})
}

func GetLog() *logrus.Logger {
	return log
}

func extractFileName(file string) string {
	parts := strings.Split(file, "/")
	return parts[len(parts)-1]
}
