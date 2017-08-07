package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/speijnik/logrus-modular.v1"
)

func main() {
	log := logrus.New()
	log.Level = logrus.DebugLevel
	log.Out = os.Stdout
	rootLogger := modular.NewRootLogger(log)
	testModule := rootLogger.GetOrCreateChild("test", logrus.InfoLevel)
	testTestModule := rootLogger.GetOrCreateChild("test.test", logrus.DebugLevel)
	// No-op, log level for "test" module is Info
	testModule.Debug("No-op, log-level is info")
	testModule.Info("Info message")
	testTestModule.Debug("Debug message of child module")
	// Logs "test2", log level for "test.test" module is Debug
	// Set level of "test" module to Info, which will propagate to child
	// module "test.test"
	testModule.SetLevel(logrus.InfoLevel)
	// No-op, log level for "test.test" module is Info
	testTestModule.Debug("No-op, SetLevel propagated Info level to child")
	testTestModule.Info("Another info message")
}
