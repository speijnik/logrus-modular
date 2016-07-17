package modular

import (
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestNewRootLogger(t *testing.T) {
	logger := &logrus.Logger{
		Level: logrus.InfoLevel,
	}
	rl := NewRootLogger(logger)
	require.NotNil(t, rl)
	require.EqualValues(t, logger, rl.GetLogger())
	require.EqualValues(t, "", rl.GetModuleName())
	require.EqualValues(t, DefaultModuleField, rl.GetModuleField())
	require.EqualValues(t, logrus.InfoLevel, rl.GetLevel())
	require.EqualValues(t, rl, rl.GetRoot())
	require.EqualValues(t, rl, rl.GetModuleLogger())
}
