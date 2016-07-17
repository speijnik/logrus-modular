package modular

import (
	"testing"
	"github.com/Sirupsen/logrus"
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