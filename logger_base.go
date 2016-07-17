package modular

import "github.com/Sirupsen/logrus"

var _ Logger = (*loggerBase)(nil)

type loggerBase struct {
	moduleLogger ModuleLogger

	fields logrus.Fields
}

func (lb *loggerBase) WithField(key string, value interface{}) Logger {
	return lb.WithFields(logrus.Fields{
		key: value,
	})
}

func (lb *loggerBase) WithFields(fields logrus.Fields) Logger {
	// Merge existing fields with new fields...

	mergedFields := make(logrus.Fields, len(lb.fields)+len(fields))

	for fieldName, fieldValue := range lb.fields {
		mergedFields[fieldName] = fieldValue
	}

	for fieldName, fieldValue := range fields {
		mergedFields[fieldName] = fieldValue
	}

	return &loggerBase{
		moduleLogger: lb.moduleLogger,
		fields:       mergedFields,
	}
}

func (lb *loggerBase) WithError(err error) Logger {
	return lb.WithField(logrus.ErrorKey, err)
}

func (lb *loggerBase) newEntry(level logrus.Level) *logrus.Entry {
	moduleLogger := lb.GetModuleLogger()
	effectiveLevel := moduleLogger.GetLevel()
	if effectiveLevel < level {
		return nil
	}

	rootLogger := moduleLogger.GetRoot()
	moduleFieldName := rootLogger.GetModuleField()

	fields := make(logrus.Fields, len(lb.fields)+1)
	fields[moduleFieldName] = moduleLogger.GetModuleName()
	for fieldName, fieldValue := range lb.fields {
		fields[fieldName] = fieldValue
	}
	return &logrus.Entry{
		Logger: rootLogger.GetLogger(),
		Data:   fields,
	}
}

func (lb *loggerBase) Debugf(format string, args ...interface{}) {
	if entry := lb.newEntry(logrus.DebugLevel); entry != nil {
		entry.Debugf(format, args...)
	}
}

func (lb *loggerBase) Infof(format string, args ...interface{}) {
	if entry := lb.newEntry(logrus.InfoLevel); entry != nil {
		entry.Infof(format, args...)
	}
}

func (lb *loggerBase) Printf(format string, args ...interface{}) {
	lb.Infof(format, args...)
}

func (lb *loggerBase) Warnf(format string, args ...interface{}) {
	if entry := lb.newEntry(logrus.WarnLevel); entry != nil {
		entry.Warnf(format, args...)
	}
}

func (lb *loggerBase) Warningf(format string, args ...interface{}) {
	lb.Warnf(format, args...)
}

func (lb *loggerBase) Errorf(format string, args ...interface{}) {
	if entry := lb.newEntry(logrus.ErrorLevel); entry != nil {
		entry.Errorf(format, args...)
	}
}

func (lb *loggerBase) Fatalf(format string, args ...interface{}) {
	if entry := lb.newEntry(logrus.FatalLevel); entry != nil {
		entry.Fatalf(format, args...)
	}
}

func (lb *loggerBase) Panicf(format string, args ...interface{}) {
	if entry := lb.newEntry(logrus.PanicLevel); entry != nil {
		entry.Panicf(format, args...)
	}
}

func (lb *loggerBase) Debug(args ...interface{}) {
	if entry := lb.newEntry(logrus.DebugLevel); entry != nil {
		entry.Debug(args...)
	}
}

func (lb *loggerBase) Info(args ...interface{}) {
	if entry := lb.newEntry(logrus.InfoLevel); entry != nil {
		entry.Info(args...)
	}
}

func (lb *loggerBase) Print(args ...interface{}) {
	lb.Info(args...)
}

func (lb *loggerBase) Warn(args ...interface{}) {
	if entry := lb.newEntry(logrus.WarnLevel); entry != nil {
		entry.Warn(args...)
	}
}

func (lb *loggerBase) Warning(args ...interface{}) {
	lb.Warn(args...)
}

func (lb *loggerBase) Error(args ...interface{}) {
	if entry := lb.newEntry(logrus.ErrorLevel); entry != nil {
		entry.Error(args...)
	}
}

func (lb *loggerBase) Fatal(args ...interface{}) {
	if entry := lb.newEntry(logrus.FatalLevel); entry != nil {
		entry.Fatal(args...)
	}
}

func (lb *loggerBase) Panic(args ...interface{}) {
	if entry := lb.newEntry(logrus.PanicLevel); entry != nil {
		entry.Panic(args...)
	}
}

func (lb *loggerBase) Debugln(args ...interface{}) {
	if entry := lb.newEntry(logrus.DebugLevel); entry != nil {
		entry.Debugln(args...)
	}
}

func (lb *loggerBase) Infoln(args ...interface{}) {
	if entry := lb.newEntry(logrus.InfoLevel); entry != nil {
		entry.Infoln(args...)
	}
}

func (lb *loggerBase) Println(args ...interface{}) {
	lb.Infoln(args...)
}

func (lb *loggerBase) Warnln(args ...interface{}) {
	if entry := lb.newEntry(logrus.WarnLevel); entry != nil {
		entry.Warnln(args...)
	}
}

func (lb *loggerBase) Warningln(args ...interface{}) {
	lb.Warnln(args...)
}

func (lb *loggerBase) Errorln(args ...interface{}) {
	if entry := lb.newEntry(logrus.ErrorLevel); entry != nil {
		entry.Errorln(args...)
	}
}

func (lb *loggerBase) Fatalln(args ...interface{}) {
	if entry := lb.newEntry(logrus.FatalLevel); entry != nil {
		entry.Fatalln(args...)
	}
}

func (lb *loggerBase) Panicln(args ...interface{}) {
	if entry := lb.newEntry(logrus.PanicLevel); entry != nil {
		entry.Panicln(args...)
	}
}

func (lb *loggerBase) GetModuleLogger() ModuleLogger {
	return lb.moduleLogger
}
