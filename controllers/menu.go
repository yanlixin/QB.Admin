package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"

	"QB.Admin/models/system"
	"QB.Admin/summaries"
	"QB.Admin/summaries/htmlhelper"
	"QB.Admin/utils"
)

//MenuController is menu controller
type MenuController struct {
	BaseAdminRouter
}

//GetAllMenus of MenuController
func (ctl *MenuController) GetAllMenus() {
	ctl.Data["HasAdd"] = true
	ctl.Data["HasView"] = true
	ctl.TplName = "system/menumanager/menulist.html"

}

//GetMenu of MenuController
func (ctl *MenuController) GetMenu() {
	menuID, _ := strconv.Atoi(ctl.Ctx.Input.Param(":id"))
	m, _ := models.DefaultMenuList.Find(menuID)
	flash := beego.ReadFromRequest(&ctl.Controller)
	if nil == m {
		flash.Error("获取表单信息错误!")
		flash.Store(&ctl.Controller)
		ctl.Redirect("/setting", 302)
		return
	}
	ctl.Data["Model"] = m
	ctl.TplName = "system/menumanager/menuview.html"

}
func (ctl *MenuController) RemoveMenu() {
	menuID, _ := strconv.Atoi(ctl.Ctx.Input.Param(":id"))
	result := summaries.JsonResultSuccess("0", nil)
	if status := models.DefaultMenuList.Delete(menuID); status {
		result = summaries.JsonResultSuccess(fmt.Sprintf("%d", menuID), nil)
	} else {
		result = summaries.JsonResultLogicFailed("Delete menu failed ", nil)
	}
	ctl.Data["json"] = result
	ctl.ServeJSON()
}

//EditMenu of MenuController
func (ctl *MenuController) EditMenu() {
	flash := beego.ReadFromRequest(&ctl.Controller)
	menuID, _ := strconv.Atoi(ctl.Ctx.Input.Param(":id"))
	m := new(models.Menu)
	if menuID > 0 {
		var status = false
		m, status = models.DefaultMenuList.Find(menuID)
		if !status {
			flash.Error("获取表单信息错误!")
			flash.Store(&ctl.Controller)
			ctl.Redirect("/setting", 302)
			return
		}
	}

	//Build parent select item
	modules := htmlhelper.DefultSelectItem()
	for _, item := range models.DefaultModuleList.GetAllModules() {
		isSelect := item.ID == m.ModuleID
		selectItem := htmlhelper.SelectItem{strconv.Itoa(item.ID), item.Name, isSelect}
		modules = append(modules, selectItem)
	}

	//Build parent select item
	parenMenus := htmlhelper.DefultSelectItem()
	for _, item := range models.DefaultMenuList.GetMenus(models.DefaultMenuList.IsTreeLeaf) {
		isSelect := item.ID == m.PID
		selectItem := htmlhelper.SelectItem{strconv.Itoa(item.ID), item.Name, isSelect}
		parenMenus = append(parenMenus, selectItem)
	}

	ctl.Data["Model"] = m

	ctl.Data["Modules"] = modules
	ctl.Data["ParenMenus"] = parenMenus
	ctl.TplName = "system/menumanager/menuedit.html"

}

//SaveMenu of MenuController
func (ctl *MenuController) SaveMenu() {
	var ob *models.Menu
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	result := summaries.JsonResultSuccess("0", nil)
	if nil != err {
		utils.Logger.Error(fmt.Sprintf("%v", err))
		result = summaries.JsonResultUnmarshalFailed(fmt.Sprintf("%v", err), nil)
	} else {
		utils.Logger.Debug(fmt.Sprintf("%+v", ob))
		num, err := models.DefaultMenuList.Save(ob)
		if nil == err {
			result = summaries.JsonResultSuccess(fmt.Sprintf("%d", num), nil)
		} else {
			utils.Logger.Error(fmt.Sprintf("%v", err))
			result = summaries.JsonResultLogicFailed(fmt.Sprintf("%v", err), nil)
		}
	}
	ctl.Data["json"] = result
	ctl.ServeJSON()
}

//GetAllMenusJSON of MenuController
func (ctl *MenuController) GetAllMenusJSON() {

	menus := models.DefaultMenuList.GetMenus(nil)
	cnnJSON := make(map[string]interface{})
	cnnJSON["sEcho"] = 1
	cnnJSON["iTotalRecords"] = len(menus)
	cnnJSON["iTotalDisplayRecords"] = len(menus)
	cnnJSON["aaData"] = menus

	//b, _ := json.Marshal(cnnJson)
	ctl.Data["json"] = cnnJSON
	ctl.ServeJSON()
}

//GetFlatMenusJSON of MenuController
func (ctl *MenuController) GetFlatMenusJSON() {

	menus := models.DefaultMenuList.GetMenus(nil)

	result := make(map[string]interface{})
	result["sEcho"] = 0
	result["iTotalRecords"] = len(menus)
	result["iTotalDisplayRecords"] = len(menus)
	result["aaData"] = menus

	//b, _ := json.Marshal(cnnJson)
	ctl.Data["json"] = result
	ctl.ServeJSON()
}
