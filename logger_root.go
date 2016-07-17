package modular

import (
	"sync"

	"github.com/Sirupsen/logrus"
)

// DefaultModuleField defines the default field to use for the module name
const DefaultModuleField = "module"

type loggerRoot struct {
	loggerModule

	logger *logrus.Logger

	moduleFieldMutex sync.Mutex
	moduleField      string
}

func (lr *loggerRoot) GetLogger() *logrus.Logger {
	return lr.logger
}

func (lr *loggerRoot) GetModuleField() string {
	lr.moduleFieldMutex.Lock()
	defer lr.moduleFieldMutex.Unlock()
	return lr.moduleField

}

func (lr *loggerRoot) SetModuleField(field string) {
	if field == "" {
		field = DefaultModuleField
	}
	lr.moduleFieldMutex.Lock()
	defer lr.moduleFieldMutex.Unlock()

	lr.moduleField = field
}
