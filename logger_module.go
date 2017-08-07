package modular

import (
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

var _ ModuleLogger = (*loggerModule)(nil)

type loggerModule struct {
	loggerBase

	name       string
	levelMutex sync.Mutex
	level      logrus.Level

	root RootLogger

	childrenMutex sync.Mutex
	children      map[string]*loggerModule
}

func (lm *loggerModule) GetModuleName() string {
	return lm.name
}

func (lm *loggerModule) SetLevel(level logrus.Level) {
	lm.levelMutex.Lock()
	defer lm.levelMutex.Unlock()
	lm.level = level

	// Propagate change to children
	lm.childrenMutex.Lock()
	defer lm.childrenMutex.Unlock()
	for _, child := range lm.children {
		child.SetLevel(level)
	}
}

func (lm *loggerModule) GetLevel() logrus.Level {
	lm.levelMutex.Lock()
	defer lm.levelMutex.Unlock()
	return lm.level
}

func (lm *loggerModule) GetRoot() RootLogger {
	return lm.root
}

func (lm *loggerModule) getLocalChildNames(moduleName string) (localName, childName string) {
	if lm.name != "" {
		moduleName = strings.TrimPrefix(moduleName, lm.name+".")
	}

	localName = moduleName
	childName = ""
	if strings.Contains(moduleName, ".") {
		moduleNameParts := strings.Split(moduleName, ".")
		localName = moduleNameParts[0]
		childName = strings.Join(moduleNameParts[1:], ".")
	}

	return
}

func (lm *loggerModule) GetChild(moduleName string) (ModuleLogger, error) {
	lm.childrenMutex.Lock()
	defer lm.childrenMutex.Unlock()

	return lm.getChild(moduleName)
}

func (lm *loggerModule) getChild(moduleName string) (ModuleLogger, error) {
	if lm.name != "" && !strings.HasPrefix(moduleName, lm.name+".") {
		moduleName = strings.Join([]string{lm.name, moduleName}, ".")
	}

	localModuleName, childModuleName := lm.getLocalChildNames(moduleName)

	childModule, ok := lm.children[localModuleName]

	if !ok {
		return nil, ErrChildNotFound
	}

	if childModuleName == "" {
		return childModule, nil
	}

	return childModule.GetChild(moduleName)
}

func (lm *loggerModule) CreateChild(moduleName string, defaultLevel logrus.Level) (ModuleLogger, error) {
	lm.childrenMutex.Lock()
	defer lm.childrenMutex.Unlock()
	return lm.createChild(moduleName, defaultLevel)
}

func (lm *loggerModule) createChild(moduleName string, defaultLevel logrus.Level) (ModuleLogger, error) {
	if lm.name != "" && !strings.HasPrefix(moduleName, lm.name+".") {
		moduleName = strings.Join([]string{lm.name, moduleName}, ".")
	}

	localModuleName, childModuleName := lm.getLocalChildNames(moduleName)
	fullLocalModuleName := localModuleName
	if lm.name != "" {
		fullLocalModuleName = strings.Join([]string{lm.name, localModuleName}, ".")
	}

	childModule, ok := lm.children[localModuleName]

	if ok && childModuleName == "" {
		return nil, ErrChildExists
	} else if ok {
		return childModule.CreateChild(moduleName, defaultLevel)
	}

	// Child does not exist, create it.
	child := &loggerModule{
		loggerBase: loggerBase{
			fields: make(logrus.Fields, 5),
		},
		name:     fullLocalModuleName,
		root:     lm.root,
		level:    defaultLevel,
		children: make(map[string]*loggerModule, 1),
	}
	child.moduleLogger = child

	lm.children[localModuleName] = child

	if childModuleName == "" {
		return child, nil
	}

	return child.CreateChild(moduleName, defaultLevel)
}

func (lm *loggerModule) GetOrCreateChild(moduleName string, defaultLevel logrus.Level) ModuleLogger {
	lm.childrenMutex.Lock()
	defer lm.childrenMutex.Unlock()

	if child, err := lm.getChild(moduleName); err == nil {
		return child
	}

	if child, err := lm.createChild(moduleName, defaultLevel); err == nil {
		return child
	}

	// should be unreachable
	return nil
}
