// Package modular provides modular logging for logrus
package modular

import "github.com/sirupsen/logrus"

// NewRootLogger creates a new root logger, that wraps the passed logrus.Logger
func NewRootLogger(logger *logrus.Logger) RootLogger {
	loggerLevel := logger.Level
	logger.Level = logrus.TraceLevel
	lr := &loggerRoot{
		logger:      logger,
		moduleField: DefaultModuleField,
	}

	lr.root = lr
	lr.children = make(map[string]*loggerModule)
	lr.level = loggerLevel
	lr.moduleLogger = lr

	return lr
}
