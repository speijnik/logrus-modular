speijnik/logrus-modular
===
[![GoDoc](https://godoc.org/github.com/speijnik/logrus-modular?status.svg)](https://godoc.org/github.com/speijnik/logrus-modular)
[![Build Status](https://travis-ci.org/speijnik/logrus-modular.svg?branch=master)](https://travis-ci.org/speijnik/logrus-modular)
[![codecov](https://codecov.io/gh/speijnik/logrus-modular/branch/master/graph/badge.svg)](https://codecov.io/gh/speijnik/logrus-modular)

Package `speijnik/logrus-modular` implements modular logging for logrus.

This allows creation of a hierarchy of loggers, whereas log-levels may
be inherited. 
The purpose of this library is simplifying the handling of loggers which
can be logically grouped and allowing configuration of log-levels on
such logger groups.

## Install

With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```sh
go get -u github.com/speijnik/logrus-modular
```

## Examples

```go
package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/speijnik/logrus-modular"
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
```
