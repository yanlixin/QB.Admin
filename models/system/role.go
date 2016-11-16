package models

import (
	"fmt"

	"errors"

	"QB.Admin/databases"
	"QB.Admin/summaries/jquery"
	"QB.Admin/utils"
)

func init() {
	DefaultRoleList = NewRoleManager()
}

type Role struct {
	ID           int    // Unique identifier
	Name         string // Description
	Desc         string
	RecordStatus int
	Permissions  map[string]int
}

type RoleManager struct {
	roles  []*Role
	lastID int
}

var DefaultRoleList *RoleManager

func fillRole(o *databases.Roles) (*Role, error) {

	if o.RoleName == "" {
		return nil, fmt.Errorf("empty RoleName")
	}
	role := new(Role)
	role.ID = o.RoleID
	role.Name = o.RoleName
	role.Desc = o.RoleDesc
	role.RecordStatus = o.RecordStatus
	role.Permissions = DefaultPermissionList.GetPermissionsByRole(o.RoleID)
	return role, nil
}

func NewRoleManager() *RoleManager {
	result := &RoleManager{}
	result.fillAll()
	return result
}
func (m *RoleManager) fillAll() error {
	var result []*Role
	roles, _ := databases.DefaultContent.GetAllRoles(10000, 0)
	for _, o := range roles {
		role, _ := fillRole(o)
		result = append(result, role)
	}
	m.roles = result
	return nil
}

// Save saves the given Task in the TaskManager.
func (m *RoleManager) GetAllRoles(param *jquery.DataTablesParam) ([]*Role, int64) {
	//var v *jquery.DataTablesParam
	var result []*Role
	//v.iDisplayStart = 1
	roles, count := databases.DefaultContent.GetAllRoles(param.DisplayLength, param.DisplayStart)

	for _, item := range roles {
		o, err := fillRole(item)
		if nil != err {
			utils.Logger.Debug(fmt.Sprintf("%v", err))
			return nil, 0
		}
		result = append(result, o)
	}

	return result, count
}
func (m *RoleManager) sync(roleID int) {

	isFound := false
	role, b := databases.DefaultContent.GetRole(roleID)
	if !b || nil == role {
		errors.New("获取角色信息失败!")
		return
	}
	newRole, _ := fillRole(role)
	for i, r := range m.roles {
		if r.ID == newRole.ID {

			m.roles[i] = newRole
			isFound = true
			break
		}
	}
	if !isFound {
		m.roles = append(m.roles, newRole)
	}

}

// Save saves the given Task in the TaskManager.
func (m *RoleManager) Save(role *Role) (int64, error) {
	isNew := true
	r := new(databases.Roles)
	if 0 < role.ID {
		isNew = false
		//角色编辑
		r.RoleID = role.ID
	}
	r.RoleName = role.Name
	r.RoleDesc = role.Desc
	r.RecordStatus = role.RecordStatus //var err error
	var roleID int64
	if isNew {
		roleID, _ = databases.DefaultContent.InsertRole(r)
	} else {
		roleID, _ = databases.DefaultContent.UpdateRole(r)
	}
	m.sync(int(roleID))
	return roleID, nil
}
func (m *RoleManager) Delete(roleID int) bool {
	_, b := databases.DefaultContent.DeleteRole(roleID)
	m.sync(roleID)
	return b
}
func (m *RoleManager) GetRolesByUser(userID int) ([]*Role, error) {

	var result []*Role
	//v.iDisplayStart = 1
	roles, b := databases.DefaultContent.GetRolesByUser(userID)
	if !b {

		utils.Logger.Debug(fmt.Sprintf("GetRolesByUser return %t", b))
		return nil, nil

	}
	for _, item := range roles {
		o, err := fillRole(item)
		if nil != err {
			utils.Logger.Debug(fmt.Sprintf("%v", err))
			return nil, err
		}
		result = append(result, o)
	}

	return result, nil

}

func (m *RoleManager) Find(roleID int) (*Role, bool) {
	for _, t := range m.roles {
		if t.ID == roleID {
			return t, true
		}
	}
	return nil, false
}
func (m *RoleManager) AddPermissions(roleID int, permissionIDs []int) bool {
	b := databases.DefaultContent.AddPermissionsToRole(roleID, permissionIDs)
	DefaultPermissionList.Reload()
	m.sync(roleID)
	return b
}
func (m *RoleManager) AddUsers(roleID int, userIDs []int) bool {
	b := databases.DefaultContent.AddUsersToRole(roleID, userIDs)
	//DefaultPermissionList.Reload()
	m.sync(roleID)
	return b
}
func (m *RoleManager) RemoveUser(roleID int, userID int) bool {
	var userIDs = []int{userID}
	b := databases.DefaultContent.DeleteUsersFromRole(roleID, userIDs)
	//DefaultPermissionList.Reload()
	m.sync(roleID)
	return b
}
