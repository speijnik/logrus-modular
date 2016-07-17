package modular

import "github.com/Sirupsen/logrus"

type Logger interface {
	WithField(key string, value interface{}) Logger
	WithFields(fields logrus.Fields) Logger
	WithError(err error) Logger

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})

	Debug(args ...interface{})
	Info(args ...interface{})
	Print(args ...interface{})
	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})

	Debugln(args ...interface{})
	Infoln(args ...interface{})
	Println(args ...interface{})
	Warnln(args ...interface{})
	Warningln(args ...interface{})
	Errorln(args ...interface{})
	Fatalln(args ...interface{})
	Panicln(args ...interface{})

	GetModuleLogger() ModuleLogger
}

type ModuleLogger interface {
	Logger

	GetModuleName() string

	SetLevel(level logrus.Level)
	GetLevel() logrus.Level

	GetRoot() RootLogger

	GetChild(moduleName string) (ModuleLogger, error)
	CreateChild(moduleName string, defaultLevel logrus.Level) (ModuleLogger, error)
	GetOrCreateChild(moduleName string, defaultLevel logrus.Level) ModuleLogger
}

type RootLogger interface {
	ModuleLogger

	GetLogger() *logrus.Logger

	GetModuleField() string
	SetModuleField(field string)
}
