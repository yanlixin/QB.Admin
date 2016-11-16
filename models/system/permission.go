package models

import (
	"errors"
	"fmt"

	"QB.Admin/databases"
	"QB.Admin/utils"
)

type Permissions struct {
	ID           int
	MenuID       int
	ActionID     string
	Name         string
	Desc         string
	Group        string
	Memo         string
	RecordStatus int
}

func init() {
	DefaultPermissionList = NewPermissionManager()
}
func fillPermission(o *databases.Permissions) (*Permissions, error) {

	if o.PermissionName == "" {
		return nil, fmt.Errorf("empty Permission Name")
	}
	per := new(Permissions)
	per.ID = o.PermissionID
	per.MenuID = o.MenuID
	per.ActionID = o.ActionID
	per.Name = o.PermissionName
	per.Desc = o.PermissionDesc
	per.Group = o.PermissionGroup
	per.Memo = o.PermissionMemo
	per.RecordStatus = o.RecordStatus
	return per, nil
}

// NewTask creates a new task given a title, that can't be empty.
func newPermission(id int, action string) (*Menu, error) {
	if action == "" {
		return nil, fmt.Errorf("empty MenuName")
	}
	return nil, nil
}

var DefaultPermissionList *PermissionManager

// TaskManager manages a list of tasks in memory.
type PermissionManager struct {
	rolePermissions map[int]map[string]int
}

// NewTaskManager returns an empty TaskManager.
func NewPermissionManager() *PermissionManager {
	result := &PermissionManager{}
	result.buildPermissionsRole()
	return result
}

// Find returns the Task with the given id in the TaskManager and a boolean
// indicating if the id was found.
func (m *PermissionManager) Find(menuID int) (*Menu, bool) {
	return nil, false
}

func (m *PermissionManager) GetPermissionsByMenu(menuID int) ([]*databases.Permissions, error) {
	result, b := databases.DefaultContent.GetAllPermissions(menuID)
	if !b {
		utils.Logger.Debug("The GetAllPermissions return false")
		err := errors.New("The GetAllPermissionsByMenu return false")
		return nil, err
	}
	return result, nil
}
func (m *PermissionManager) buildPermissionsRole() error {

	result, b := databases.DefaultContent.GetRolePermissions()
	if !b {
		err := errors.New("The GetAllGetPermissionsByRolePermissionsByMenu return false")
		return err
	}
	m.rolePermissions = result
	return nil
}
func (m *PermissionManager) GetPermissionsByRole(roleID int) map[string]int {
	//utils.Logger.Debug("%v", m.rolePermissions)
	return m.rolePermissions[roleID]
}
func (m *PermissionManager) Reload() {
	m.buildPermissionsRole()
}
func (m *PermissionManager) GetPermissions() []*Permissions {
	var result []*Permissions
	list, _ := databases.DefaultContent.GetAllPermissions(0)
	for _, o := range list {
		per, _ := fillPermission(o)
		result = append(result, per)
	}
	return result
}
