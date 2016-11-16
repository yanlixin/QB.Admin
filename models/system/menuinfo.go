package models

import "fmt"

type MenuInfo struct {
	ID        int    // Unique identifier
	Title     string // Description
	IconClass string
	Action    string
	Parent    []*interface{}
	Children  []*MenuInfo
}

// NewTask creates a new task given a title, that can't be empty.
func newMenu(id int, menuName string, iconClass string, action string, children []interface{}, parent interface{}) (interface{}, error) {
	if menuName == "" {
		return nil, fmt.Errorf("empty MenuName")
	}
	var result = make(map[string]interface{})
	result["ID"] = id
	result["Name"] = menuName
	result["IconClass"] = iconClass
	result["Action"] = action
	result["Children"] = children
	result["Parent"] = parent
	return result, nil
}
func BuildMenu(pid int, menus []*Menu) ([]interface{}, error) {
	var result []interface{}
	if nil == menus {
		menus = DefaultMenuList.menus
	}
	for _, t := range menus {
		if t.PID == pid {
			childen, err := BuildMenu(t.ID, menus)
			if nil != err {
				return nil, err
			}
			parent := getMenuParent(pid, menus)
			menu, err := newMenu(t.ID, t.Name, t.IconUrl, t.NavUrl, childen, parent)
			if nil != err {
				return nil, err
			}
			result = append(result, menu)
		}
	}
	return result, nil
}
func getMenuParent(pid int, menus []*Menu) interface{} {
	var result = make(map[string]interface{})
	for _, t := range menus {
		if t.ID == pid {
			result["ID"] = t.ID
			result["Name"] = t.Name
			result["Action"] = t.NavUrl
			if 0 != t.PID {
				parent := getMenuParent(t.PID, menus)
				result["Parent"] = parent

			}
		}
	}
	return result
}
