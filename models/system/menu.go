package models

import (
	"errors"
	"fmt"

	"QB.Admin/databases"
	"QB.Admin/utils"
)

func init() {
	DefaultMenuList = newMenuManager()
}

type Menu struct {
	ID           int
	ModuleID     int
	PID          int
	Name         string
	IconUrl      string
	NavUrl       string
	Target       string
	Desc         string
	SortIndex    int
	IsTreeLeaf   bool
	ActionID     string
	TreeLevel    int
	RecordStatus int
}
type filterMain func(*Menu) bool

func (m *MenuManager) IsTreeLeaf(menu *Menu) bool {

	return !menu.IsTreeLeaf
}

var DefaultMenuList *MenuManager

// TaskManager manages a list of tasks in memory.
type MenuManager struct {
	menus  []*Menu
	lastID int
}

// NewTaskManager returns an empty TaskManager.
func newMenuManager() *MenuManager {
	result := &MenuManager{}
	//dbMenus, _ := databases.DefaultContent.GetAllMenus()
	//	menus, err := result.BuildMenu(0, dbMenus)
	//	if nil != err {
	//		utils.Logger.Error(fmt.Sprintf("%v", err))
	//	}
	//result.menus = menus
	menus, _ := result.fillAllMenus()
	result.menus = menus
	return result
}

// Find returns the Task with the given id in the TaskManager and a boolean
// indicating if the id was found.

func (m *MenuManager) fillMenu(t *databases.Menus) (*Menu, error) {

	if t.MenuName == "" {
		return nil, fmt.Errorf("empty MenuName")
	}
	mo := new(Menu)
	mo.ID = t.MenuID
	mo.ModuleID = t.ModuleID
	mo.PID = t.MenuPID
	mo.Name = t.MenuName
	mo.IconUrl = t.MenuIconUrl
	mo.NavUrl = t.MenuNavUrl
	mo.Target = t.MenuTarget
	mo.Desc = t.MenuDesc
	mo.SortIndex = t.SortIndex
	mo.IsTreeLeaf = t.IsTreeLeaf
	mo.ActionID = t.ActionID
	mo.TreeLevel = t.TreeLevel
	mo.RecordStatus = t.RecordStatus
	return mo, nil
}
func (m *MenuManager) fillAllMenus() ([]*Menu, error) {
	var result []*Menu
	dbMenus, _ := databases.DefaultContent.GetAllMenus()
	for _, t := range dbMenus {
		mo, _ := m.fillMenu(t)
		result = append(result, mo)
	}
	return result, nil
}

func (m *MenuManager) sync(menuID int) {

	isFound := false
	menu, b := databases.DefaultContent.GetMenu(menuID)
	if !b || nil == menu {
		errors.New("获取角色信息失败!")
		return
	}
	newMenu, _ := m.fillMenu(menu)
	for i, r := range m.menus {
		if r.ID == newMenu.ID {

			m.menus[i] = newMenu
			isFound = true
			break
		}
	}
	if !isFound {
		m.menus = append(m.menus, newMenu)
	}

}

func (m *MenuManager) GetMenus(f filterMain) []*Menu {
	if nil != f {
		var result []*Menu
		for _, item := range m.menus {
			if f(item) {
				result = append(result, item)
			}

		}
		return result
	}
	return m.menus
}

// Find returns the Task with the given id in the TaskManager and a boolean
// indicating if the id was found.
func (m *MenuManager) Find(menuID int) (*Menu, bool) {

	for _, t := range m.menus {
		if t.ID == menuID {
			return t, true
		}

	}

	return nil, false
}
func (m *MenuManager) Save(menu *Menu) (int64, error) {
	isNew := true
	mo := new(databases.Menus)
	if 0 < menu.ID {
		isNew = false
		//角色编辑
		mo.MenuID = menu.ID
	}
	mo.ModuleID = menu.ModuleID
	mo.MenuPID = menu.PID
	mo.MenuName = menu.Name
	mo.MenuIconUrl = menu.IconUrl
	mo.MenuNavUrl = menu.NavUrl
	mo.MenuTarget = menu.Target
	mo.MenuDesc = menu.Desc
	mo.SortIndex = menu.SortIndex
	mo.IsTreeLeaf = menu.IsTreeLeaf
	mo.ActionID = menu.ActionID
	mo.TreeLevel = menu.TreeLevel
	mo.RecordStatus = menu.RecordStatus
	var menuID int64
	if isNew {
		menuID, _ = databases.DefaultContent.InsertMenu(mo)
	} else {
		menuID, _ = databases.DefaultContent.UpdateMenu(mo)
	}
	m.sync(int(menuID))
	return menuID, nil
}
func (m *MenuManager) GetMenuPermission(menuID int) ([]*databases.Permissions, error) {
	result, b := databases.DefaultContent.GetAllPermissions(menuID)
	if !b {
		utils.Logger.Debug(fmt.Sprintf("GetAllPermission return %t", b))
		return nil, nil
	}
	return result, nil
}
func (m *MenuManager) Delete(menuID int) bool {
	_, b := databases.DefaultContent.DeleteMenu(menuID)
	m.sync(menuID)
	return b
}
