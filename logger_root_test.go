package modular

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestLoggerRoot_GetLogger(t *testing.T) {
	rl := &loggerRoot{
		logger: &logrus.Logger{},
	}

	require.EqualValues(t, rl.logger, rl.GetLogger())
}

func TestLoggerRoot_GetModuleField(t *testing.T) {
	rl := &loggerRoot{
		moduleField: "test",
	}

	require.EqualValues(t, "test", rl.GetModuleField())

	rl.moduleField = DefaultModuleField
	require.EqualValues(t, DefaultModuleField, rl.GetModuleField())
}

func TestLoggerRoot_SetModuleField(t *testing.T) {
	rl := &loggerRoot{
		moduleField: "test",
	}

	rl.SetModuleField("test2")
	require.EqualValues(t, "test2", rl.GetModuleField())

	rl.SetModuleField("")
	require.EqualValues(t, DefaultModuleField, rl.GetModuleField())
}

func TestLoggerRoot_GetOrCreateChild(t *testing.T) {
	rl := NewRootLogger(logrus.New())
	require.NotNil(t, rl)
	moduleLogger := rl.GetOrCreateChild("test", logrus.DebugLevel)
	require.NotNil(t, moduleLogger)
	require.EqualValues(t, logrus.DebugLevel, moduleLogger.GetLevel())
	require.EqualValues(t, "test", moduleLogger.GetModuleName())

	// Test nesting
	nestedModuleLogger := moduleLogger.GetOrCreateChild("nested", logrus.ErrorLevel)
	require.NotNil(t, nestedModuleLogger)
	require.EqualValues(t, logrus.ErrorLevel, nestedModuleLogger.GetLevel())
	require.EqualValues(t, "test.nested", nestedModuleLogger.GetModuleName())

}
