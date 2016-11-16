package models

import (
	"fmt"

	"QB.Admin/databases"
)

func init() {
	DefaultModuleList = newModuleManager()
}

type Module struct {
	ID     int    // Unique identifier
	Name   string // Description
	Desc   string
	SortIndex    int
	RecordStatus int
}

type ModuleManager struct {
	modules []*Module
	lastID  int
}

var DefaultModuleList *ModuleManager

func (m *ModuleManager) fill(o *databases.Modules) (*Module, error) {

	if o.ModuleName == "" {
		return nil, fmt.Errorf("empty ModuleName")
	}
	module := new(Module)
	module.ID = o.ModuleID
	module.Name = o.ModuleName
	module.Desc = o.ModuleDesc
	module.SortIndex = o.SortIndex
	module.RecordStatus = o.RecordStatus
	return module, nil
}

func newModuleManager() *ModuleManager {
	result := &ModuleManager{}
	result.fillAll()
	return result
}
func (m *ModuleManager) fillAll() error {
	var result []*Module
	modules, _ := databases.DefaultContent.GetAllModules()
	for _, o := range modules {
		module, _ := m.fill(o)
		result = append(result, module)
	}
	m.modules = result
	return nil
}

// Save saves the given Task in the TaskManager.
func (m *ModuleManager) GetAllModules() []*Module {
	return m.modules
}

func (m *ModuleManager) Find(moduleID int) *Module {
	for _, t := range m.modules {
		if t.ID == moduleID {
			return t
		}
	}
	return nil
}
