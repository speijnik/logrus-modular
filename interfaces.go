package modular

import "github.com/sirupsen/logrus"

// Logger defines the baseline logger interface
type Logger interface {
	// WithField extends the current logger's fields with the given field and value and returns a new logger
	WithField(key string, value interface{}) Logger
	// WithFields extends the current logger's fields with the given fields and returns a new logger
	WithFields(fields logrus.Fields) Logger
	// WithError extends the current logger's fields with an error field and returns a new logger
	WithError(err error) Logger

	Tracef(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})

	Trace(args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Print(args ...interface{})
	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})

	Traceln(args ...interface{})
	Debugln(args ...interface{})
	Infoln(args ...interface{})
	Println(args ...interface{})
	Warnln(args ...interface{})
	Warningln(args ...interface{})
	Errorln(args ...interface{})
	Fatalln(args ...interface{})
	Panicln(args ...interface{})

	// GetModuleLogger returns the associated ModuleLogger
	GetModuleLogger() ModuleLogger
}

// ModuleLogger defines the interface implemented by a module logger
type ModuleLogger interface {
	Logger

	// GetModuleName returns the (full) module name
	GetModuleName() string

	// SetLevel sets the module's log level to the given level and recursively propagates this change
	// to all children.
	SetLevel(level logrus.Level)
	// GetLevel returns the module's log level
	GetLevel() logrus.Level

	// GetRoot returns the associated RootLogger
	GetRoot() RootLogger

	// GetChild returns the child with the given name
	GetChild(moduleName string) (ModuleLogger, error)
	// CreateChild creates a child with the given name
	CreateChild(moduleName string, defaultLevel logrus.Level) (ModuleLogger, error)
	// GetOrCreateChild tries returns an existing child or creates it, if it is missing
	GetOrCreateChild(moduleName string, defaultLevel logrus.Level) ModuleLogger
}

// RootLogger defines the interface implemented by a root logger
type RootLogger interface {
	ModuleLogger

	// GetLogger returns the underlying logrus.Logger
	GetLogger() *logrus.Logger

	// GetModuleField returns the name of the module field
	GetModuleField() string
	// SetModuleField sets the module field.
	// Sets field to DefaultModuleField if empty string is passed in.
	SetModuleField(field string)
}
