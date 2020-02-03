package modular

import (
	"testing"

	"errors"

	"bytes"
	"encoding/json"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestLoggerBase_WithFields(t *testing.T) {
	// Create a new logger
	lb := &loggerBase{
		moduleLogger: &loggerModule{},
		fields: logrus.Fields{
			"test": "test",
		},
	}

	// Create a new logger by calling WithFields
	logger := lb.WithFields(logrus.Fields{
		"test2": "test2",
	})
	require.NotNil(t, logger)

	lb2, ok := logger.(*loggerBase)
	require.EqualValues(t, true, ok)
	// Check if the two loggers are not equal
	require.NotEqual(t, lb, lb2)

	// Check if the old logger's fields have not been modified
	require.Len(t, lb.fields, 1)
	// Check if the new logger's fields are an extension of the old logger's fields
	require.Len(t, lb2.fields, 2)

	// Check field contents of new logger
	require.Contains(t, lb2.fields, "test")
	require.Contains(t, lb2.fields, "test2")
	require.EqualValues(t, "test", lb2.fields["test"])
	require.EqualValues(t, "test2", lb2.fields["test2"])

	// Check field contents of old logger
	require.Contains(t, lb.fields, "test")
	require.EqualValues(t, "test", lb.fields["test"])

	// Test override of field value
	logger = lb.WithFields(logrus.Fields{
		"test": "test2",
	})
	lb2, ok = logger.(*loggerBase)
	require.EqualValues(t, true, ok)
	// Check if the two loggers are not equal
	require.NotEqual(t, lb, lb2)

	// Check if the old logger's fields have not been modified
	require.Len(t, lb.fields, 1)
	// Check if the new logger's fields are an extension of the old logger's fields
	require.Len(t, lb2.fields, 1)

	// Check field contents of new logger
	require.Contains(t, lb2.fields, "test")
	require.EqualValues(t, "test2", lb2.fields["test"])

	// Check field contents of old logger
	require.Contains(t, lb.fields, "test")
	require.EqualValues(t, "test", lb.fields["test"])
}

func TestLoggerBase_WithField(t *testing.T) {
	// Create a new logger
	lb := &loggerBase{
		moduleLogger: &loggerModule{},
		fields: logrus.Fields{
			"test": "test",
		},
	}

	logger := lb.WithField("test2", "test2")
	require.NotNil(t, logger)

	lb2, ok := logger.(*loggerBase)
	require.EqualValues(t, true, ok)
	// Check if the two loggers are not equal
	require.NotEqual(t, lb, lb2)

	// Check if the old logger's fields have not been modified
	require.Len(t, lb.fields, 1)
	// Check if the new logger's fields are an extension of the old logger's fields
	require.Len(t, lb2.fields, 2)

	// Check field contents of new logger
	require.Contains(t, lb2.fields, "test")
	require.Contains(t, lb2.fields, "test2")
	require.EqualValues(t, "test", lb2.fields["test"])
	require.EqualValues(t, "test2", lb2.fields["test2"])

	// Check field contents of old logger
	require.Contains(t, lb.fields, "test")
	require.EqualValues(t, "test", lb.fields["test"])
}

func TestLoggerBase_WithError(t *testing.T) {
	// Create a new logger
	lb := &loggerBase{
		moduleLogger: &loggerModule{},
		fields: logrus.Fields{
			"test": "test",
		},
	}

	err := errors.New("test2")

	logger := lb.WithError(err)
	require.NotNil(t, logger)

	lb2, ok := logger.(*loggerBase)
	require.EqualValues(t, true, ok)
	// Check if the two loggers are not equal
	require.NotEqual(t, lb, lb2)

	// Check if the old logger's fields have not been modified
	require.Len(t, lb.fields, 1)
	// Check if the new logger's fields are an extension of the old logger's fields
	require.Len(t, lb2.fields, 2)

	// Check field contents of new logger
	require.Contains(t, lb2.fields, "test")
	require.Contains(t, lb2.fields, "error")
	require.EqualValues(t, "test", lb2.fields["test"])
	require.EqualValues(t, err, lb2.fields["error"])

	// Check field contents of old logger
	require.Contains(t, lb.fields, "test")
	require.EqualValues(t, "test", lb.fields["test"])
}

func TestLoggerBase_GetModuleLogger(t *testing.T) {
	lb := &loggerBase{
		moduleLogger: &loggerModule{},
	}

	require.EqualValues(t, lb.moduleLogger, lb.GetModuleLogger())
}

func TestLoggerBase_NewEntry(t *testing.T) {
	lb := &loggerBase{
		moduleLogger: &loggerModule{
			level: logrus.DebugLevel,
			name:  "test_module",
			root: &loggerRoot{
				logger:      &logrus.Logger{},
				moduleField: "module",
			},
		},
	}

	levels := []logrus.Level{
		logrus.DebugLevel,
		logrus.InfoLevel,
		logrus.WarnLevel,
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	}

	// Test our "base" case: debug logging enabled at module level,
	// newEntry called with debug level.
	// This must return a new entry.
	require.NotNil(t, lb.newEntry(logrus.DebugLevel))

	// Test if newEntry behaves correctly...
	for i, level := range levels {
		if i == 0 {
			// Ignore debug level above
			continue
		}
		lb.moduleLogger.SetLevel(level)
		for j := 0; j < i; j++ {
			require.Nil(t, lb.newEntry(levels[j]), "Level is %s, newEntry(%s) should return nil", level.String(), levels[j].String())
		}
		for j := i; j < len(levels); j++ {
			require.NotNil(t, lb.newEntry(levels[j]), "Level is %s, newEntry(%s) should not return nil", levels[j].String(), levels[j].String())
		}
	}

	// Test if newEntry creates a copy of fields...
	lb.moduleLogger.SetLevel(logrus.DebugLevel)
	lb.fields = logrus.Fields{
		"test":  "test",
		"test2": "test2",
	}

	entry := lb.newEntry(logrus.DebugLevel)
	require.NotNil(t, entry)
	require.Len(t, entry.Data, len(lb.fields)+1)
	require.Contains(t, entry.Data, "test")
	require.EqualValues(t, "test", entry.Data["test"])
	require.Contains(t, entry.Data, "test2")
	require.EqualValues(t, "test2", entry.Data["test2"])
	require.Contains(t, entry.Data, "module")
	require.EqualValues(t, "test_module", entry.Data["module"])
}

func testLogFunction(t *testing.T, level logrus.Level, expectedMessage string, logFn func(*loggerBase)) {
	buffer := bytes.NewBufferString("")

	levels := []logrus.Level{
		logrus.TraceLevel,
		logrus.DebugLevel,
		logrus.InfoLevel,
		logrus.WarnLevel,
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	}

	lb := &loggerBase{
		moduleLogger: &loggerModule{
			level: level,
			name:  "test_module",
			root: &loggerRoot{
				logger: &logrus.Logger{
					Out:       buffer,
					Formatter: &logrus.JSONFormatter{},
					Level:     logrus.TraceLevel,
				},
				moduleField: "module",
			},
		},
	}

	// Check if logging works
	logFn(lb)
	require.NotEmpty(t, buffer.Bytes())

	var data map[string]interface{}
	require.NoError(t, json.Unmarshal(buffer.Bytes(), &data))

	require.EqualValues(t, expectedMessage, data["msg"])
	require.EqualValues(t, "test_module", data["module"])
	lvl, err := logrus.ParseLevel(data["level"].(string))
	require.NoError(t, err)
	require.EqualValues(t, level, lvl)

	if level != logrus.PanicLevel {
		ignoredLevel := logrus.PanicLevel
		foundLevel := false
		for _, lvl := range levels {
			if lvl == level {
				foundLevel = true
			} else if foundLevel {
				ignoredLevel = lvl
				break
			}
		}
		// Check if we are not logging if the level is excluded
		buffer.Reset()
		lb.moduleLogger.SetLevel(ignoredLevel)
		logFn(lb)
		require.Empty(t, buffer.Bytes())
	}
}

func TestLoggerBase_Trace(t *testing.T) {
	testLogFunction(t, logrus.TraceLevel, "test", func(lb *loggerBase) {
		lb.Trace("test")
	})
}

func TestLoggerBase_Tracef(t *testing.T) {
	testLogFunction(t, logrus.TraceLevel, "test: 1", func(lb *loggerBase) {
		lb.Tracef("test: %d", 1)
	})
}

func TestLoggerBase_Traceln(t *testing.T) {
	testLogFunction(t, logrus.TraceLevel, "test", func(lb *loggerBase) {
		lb.Traceln("test")
	})
}

func TestLoggerBase_Debug(t *testing.T) {
	testLogFunction(t, logrus.DebugLevel, "test", func(lb *loggerBase) {
		lb.Debug("test")
	})
}

func TestLoggerBase_Debugf(t *testing.T) {
	testLogFunction(t, logrus.DebugLevel, "test: 1", func(lb *loggerBase) {
		lb.Debugf("test: %d", 1)
	})
}

func TestLoggerBase_Debugln(t *testing.T) {
	testLogFunction(t, logrus.DebugLevel, "test", func(lb *loggerBase) {
		lb.Debugln("test")
	})
}

func TestLoggerBase_Error(t *testing.T) {
	testLogFunction(t, logrus.ErrorLevel, "test", func(lb *loggerBase) {
		lb.Error("test")
	})
}

func TestLoggerBase_Errorf(t *testing.T) {
	testLogFunction(t, logrus.ErrorLevel, "test: 1", func(lb *loggerBase) {
		lb.Errorf("test: %d", 1)
	})
}

func TestLoggerBase_Errorln(t *testing.T) {
	testLogFunction(t, logrus.ErrorLevel, "test", func(lb *loggerBase) {
		lb.Errorln("test")
	})
}

func TestLoggerBase_Info(t *testing.T) {
	testLogFunction(t, logrus.InfoLevel, "test", func(lb *loggerBase) {
		lb.Info("test")
	})
}

func TestLoggerBase_Infof(t *testing.T) {
	testLogFunction(t, logrus.InfoLevel, "test: 1", func(lb *loggerBase) {
		lb.Infof("test: %d", 1)
	})
}

func TestLoggerBase_Infoln(t *testing.T) {
	testLogFunction(t, logrus.InfoLevel, "test", func(lb *loggerBase) {
		lb.Infoln("test")
	})
}

func TestLoggerBase_Print(t *testing.T) {
	testLogFunction(t, logrus.InfoLevel, "test", func(lb *loggerBase) {
		lb.Print("test")
	})
}

func TestLoggerBase_Printf(t *testing.T) {
	testLogFunction(t, logrus.InfoLevel, "test: 1", func(lb *loggerBase) {
		lb.Printf("test: %d", 1)
	})
}

func TestLoggerBase_Println(t *testing.T) {
	testLogFunction(t, logrus.InfoLevel, "test", func(lb *loggerBase) {
		lb.Println("test")
	})
}

func TestLoggerBase_Warn(t *testing.T) {
	testLogFunction(t, logrus.WarnLevel, "test", func(lb *loggerBase) {
		lb.Warn("test")
	})
}

func TestLoggerBase_Warnf(t *testing.T) {
	testLogFunction(t, logrus.WarnLevel, "test: 1", func(lb *loggerBase) {
		lb.Warnf("test: %d", 1)
	})
}

func TestLoggerBase_Warnln(t *testing.T) {
	testLogFunction(t, logrus.WarnLevel, "test", func(lb *loggerBase) {
		lb.Warnln("test")
	})
}

func TestLoggerBase_Warning(t *testing.T) {
	testLogFunction(t, logrus.WarnLevel, "test", func(lb *loggerBase) {
		lb.Warning("test")
	})
}

func TestLoggerBase_Warningf(t *testing.T) {
	testLogFunction(t, logrus.WarnLevel, "test: 1", func(lb *loggerBase) {
		lb.Warningf("test: %d", 1)
	})
}

func TestLoggerBase_Warningln(t *testing.T) {
	testLogFunction(t, logrus.WarnLevel, "test", func(lb *loggerBase) {
		lb.Warningln("test")
	})
}

func TestLoggerBase_Panic(t *testing.T) {
	testLogFunction(t, logrus.PanicLevel, "test", func(lb *loggerBase) {
		defer func() {
			if r := recover(); r == nil {
				t.Fatal("Expected panic, but did not panic")
			}
		}()
		lb.Panic("test")
	})
}

func TestLoggerBase_Panicf(t *testing.T) {
	testLogFunction(t, logrus.PanicLevel, "test: 1", func(lb *loggerBase) {
		defer func() {
			if r := recover(); r == nil {
				t.Fatal("Expected panic, but did not panic")
			}
		}()
		lb.Panicf("test: %d", 1)
	})
}

func TestLoggerBase_Panicln(t *testing.T) {
	testLogFunction(t, logrus.PanicLevel, "test", func(lb *loggerBase) {
		defer func() {
			if r := recover(); r == nil {
				t.Fatal("Expected panic, but did not panic")
			}
		}()
		lb.Panicln("test")
	})
}
