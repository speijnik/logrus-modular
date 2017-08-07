package modular

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestLoggerModule_GetModuleName(t *testing.T) {
	lm := &loggerModule{
		name: "test",
	}

	require.EqualValues(t, "test", lm.GetModuleName())
}

func TestLoggerModule_GetRoot(t *testing.T) {
	lm := &loggerModule{
		root: &loggerRoot{},
	}

	require.EqualValues(t, lm.root, lm.GetRoot())
}

func TestLoggerModule_GetLocalChildNames(t *testing.T) {
	lm := &loggerModule{
		name: "test.module",
	}

	local, child := lm.getLocalChildNames("test2")
	require.EqualValues(t, "test2", local)
	require.EqualValues(t, "", child)

	local, child = lm.getLocalChildNames("test2.test3")
	require.EqualValues(t, "test2", local)
	require.EqualValues(t, "test3", child)

	local, child = lm.getLocalChildNames("test2.test3.test4")
	require.EqualValues(t, "test2", local)
	require.EqualValues(t, "test3.test4", child)

	// Test stripping of prefix...
	local, child = lm.getLocalChildNames("test.module.test2.test3.test4")
	require.EqualValues(t, "test2", local)
	require.EqualValues(t, "test3.test4", child)
}

func TestLoggerModule_CreateChild(t *testing.T) {
	lm := &loggerModule{
		name:     "test.module",
		children: make(map[string]*loggerModule),
	}

	// Test "local" creation
	child, err := lm.CreateChild("test", logrus.FatalLevel)
	require.NoError(t, err)
	require.NotNil(t, child)
	require.EqualValues(t, "test.module.test", child.GetModuleName())
	require.EqualValues(t, logrus.FatalLevel, child.GetLevel())

	child, err = lm.CreateChild("test", logrus.FatalLevel)
	require.Nil(t, child)
	require.EqualError(t, err, ErrChildExists.Error())

	// Test "local" creation with full name
	child, err = lm.CreateChild("test.module.test2", logrus.FatalLevel)
	require.NoError(t, err)
	require.NotNil(t, child)
	require.EqualValues(t, "test.module.test2", child.GetModuleName())
	require.EqualValues(t, logrus.FatalLevel, child.GetLevel())

	child, err = lm.CreateChild("test.module.test2", logrus.FatalLevel)
	require.Nil(t, child)
	require.EqualError(t, err, ErrChildExists.Error())

	child, err = lm.CreateChild("test2", logrus.FatalLevel)
	require.Nil(t, child)
	require.EqualError(t, err, ErrChildExists.Error())

	// Test nested creation
	child, err = lm.CreateChild("test3.test4", logrus.FatalLevel)
	require.NoError(t, err)
	require.NotNil(t, child)
	require.EqualValues(t, "test.module.test3.test4", child.GetModuleName())
	require.EqualValues(t, logrus.FatalLevel, child.GetLevel())

	// Test nested creation, top-level exists...
	child, err = lm.CreateChild("test3.test5", logrus.FatalLevel)
	require.NoError(t, err)
	require.NotNil(t, child)
	require.EqualValues(t, "test.module.test3.test5", child.GetModuleName())
	require.EqualValues(t, logrus.FatalLevel, child.GetLevel())
}

func TestLoggerModule_GetChild(t *testing.T) {
	lm := &loggerModule{
		name:     "test.module",
		children: make(map[string]*loggerModule),
	}

	// Test: non-existent child
	child, err := lm.GetChild("test.module.nest")
	require.Nil(t, child)
	require.EqualError(t, err, ErrChildNotFound.Error())

	// Create nested children
	child, err = lm.CreateChild("test.module.nest.nest2", logrus.FatalLevel)
	require.NoError(t, err)
	require.NotNil(t, child)

	// Get "local" child
	child, err = lm.GetChild("test.module.nest")
	require.NoError(t, err)
	require.NotNil(t, child)
	require.EqualValues(t, "test.module.nest", child.GetModuleName())

	// Get nested child
	child, err = lm.GetChild("test.module.nest.nest2")
	require.NoError(t, err)
	require.NotNil(t, "test.module.nest.nest2", child)

	// Test nested get, missing prefix
	child, err = lm.getChild("nest.nest2")
	require.NoError(t, err)
	require.NotNil(t, "test.module.nest.nest2", child)
}

func TestLoggerModule_GetOrCreateChild(t *testing.T) {
	lm := &loggerModule{
		name:     "test.module",
		children: make(map[string]*loggerModule),
	}

	// Test "local" creation
	child := lm.GetOrCreateChild("test", logrus.FatalLevel)
	require.NotNil(t, child)
	require.EqualValues(t, "test.module.test", child.GetModuleName())
	require.EqualValues(t, logrus.FatalLevel, child.GetLevel())

	// Test: existing child
	child2 := lm.GetOrCreateChild("test", logrus.FatalLevel)
	require.NotNil(t, child2)
	require.EqualValues(t, "test.module.test", child2.GetModuleName())
	require.EqualValues(t, logrus.FatalLevel, child2.GetLevel())
	require.EqualValues(t, child, child2)
}

func TestLoggerModule_GetLevel(t *testing.T) {
	lm := &loggerModule{
		level: logrus.FatalLevel,
	}

	require.EqualValues(t, logrus.FatalLevel, lm.GetLevel())
}

func TestLoggerModule_SetLevel(t *testing.T) {
	lm := &loggerModule{
		name:     "test.module",
		level:    logrus.FatalLevel,
		children: make(map[string]*loggerModule),
	}
	require.EqualValues(t, logrus.FatalLevel, lm.GetLevel())

	child, err := lm.CreateChild("test.module.nest", logrus.InfoLevel)
	require.NoError(t, err)
	require.NotNil(t, child)
	require.EqualValues(t, logrus.InfoLevel, child.GetLevel())

	lm.SetLevel(logrus.DebugLevel)
	require.EqualValues(t, logrus.DebugLevel, lm.GetLevel())
	require.EqualValues(t, logrus.DebugLevel, child.GetLevel())
}
