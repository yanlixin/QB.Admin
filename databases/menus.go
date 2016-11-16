package databases

import (
	"fmt"

	"QB.Admin/utils"
	"github.com/astaxie/beego/orm"
)

func (u *DbContent) GetAllMenus() ([]*Menus, bool) {
	status := true
	o := orm.NewOrm()
	var menus []*Menus
	_, err := o.QueryTable("pa_menus").OrderBy("SortIndex").All(&menus)
	if nil != err {
		status = false
		utils.Logger.Error(fmt.Sprintf("%v", err))
	}
	return menus, status
}
func (c *DbContent) GetMenu(menuID int) (*Menus, bool) {
	status := true
	o := orm.NewOrm()
	var menu Menus
	err := o.QueryTable("pa_menus").Filter("MenuID", menuID).One(&menu)
	if nil != err {
		status = false
		utils.Logger.Error(fmt.Sprintf("%v", err))
	}
	return &menu, status
}

func (db *DbContent) InsertMenu(m *Menus) (int64, bool) {
	status := true
	prpareSave(m)

	o := orm.NewOrm()
	num, err := o.Insert(m)
	if err != nil {
		utils.Logger.Error(fmt.Sprintf("%v", err))
		return num, false
	}
	return num, status
}
func (db *DbContent) UpdateMenu(m *Menus) (int64, bool) {
	prpareSave(m)
	o := orm.NewOrm()
	_, err := o.Update(m)
	if err != nil {
		utils.Logger.Error(fmt.Sprintf("%v", err))
		return int64(m.MenuID), false
	}
	return int64(m.MenuID), true
}
func (db *DbContent) DeleteMenu(menuID int) (int64, bool) {
	result := true
	o := orm.NewOrm()
	var m = new(Menus)
	m.MenuID = menuID
	m.RecordStatus = -1
	num, err := o.Update(m, "RecordStatus")
	if err != nil {
		utils.Logger.Error(fmt.Sprintf("%v", err))
		result = false
	}
	return num, result
}
