package controllers

import (
	"fmt"

	"QB.Admin/models/system"
	"QB.Admin/summaries/jquery"
)

//DefaultController 继承 BaseAdminRouter
type DefaultController struct {
	BaseAdminRouter
}

//Get of DefaultController
func (ctr *DefaultController) Get() {
	ctr.Layout = "shared/layout.html"
	ctr.LayoutSections = make(map[string]string)
	ctr.LayoutSections["Head"] = "shared/head.html"
	ctr.LayoutSections["Scripts"] = "shared/scripts.html"
	ctr.LayoutSections["Sidebar"] = "shared/leftnav.html"
	ctr.LayoutSections["Topbar"] = "shared/topbar.html"
	ctr.TplName = "index.html"

}

//GetHome of DefaultController
func (ctr *DefaultController) GetHome() {
	ctr.TplName = "index.html"
}

//GetModel of UserController
func (ctr *UserController) GetModel() {
	ctr.TplName = "tools/modelview.html"

}

//GetModelJSON of UserController
func (ctr *UserController) GetModelJSON() {
	tableName := ctr.Ctx.Input.Param(":tableName")
	m := models.NewUserManager()
	param := &jquery.DataTablesParam{}
	ctr.ParseForm(&param)
	//dd := this.Ctx.Request.Form
	//result := &jquery.DataTablesParam{}
	//	fmt.Println(err)
	//}
	//dd := this.Ctx.Request
	//fmt.Println(param)
	fmt.Println(param)
	models, count := m.GenModel(tableName, param)
	result := make(map[string]interface{})
	result["sEcho"] = 3
	result["iTotalRecords"] = count
	result["iTotalDisplayRecords"] = count
	result["aaData"] = models

	//b, _ := json.Marshal(cnnJson)
	ctr.Data["json"] = result
	ctr.ServeJSON()
}

//GetUserInfo of DefaultController

type CommonController struct {
	BaseAdminRouter
}

func (ctr *CommonController) GetOrgUnitJSON() {
	var list []interface{}

	top := make(map[string]interface{})
	top["Name"] = "top"
	top["ID"] = "1"
	list = append(list, top)
	one := make(map[string]interface{})
	one["Name"] = "one"
	one["ID"] = "2"
	list = append(list, one)
	result := make(map[string]interface{})
	result["sEcho"] = 0
	result["iTotalRecords"] = len(list)
	result["iTotalDisplayRecords"] = len(list)
	result["aaData"] = list

	//b, _ := json.Marshal(cnnJson)
	ctr.Data["json"] = result
	ctr.ServeJSON()
}
func (ctr *CommonController) GetEmployeeJSON() {
	var list []interface{}
	orgUnitId := ctr.Ctx.Input.Param(":orgunitid")
	top := make(map[string]interface{})
	top["Name"] = fmt.Sprintf("%s", orgUnitId)
	top["ID"] = orgUnitId
	list = append(list, top)

	result := make(map[string]interface{})
	result["sEcho"] = 0
	result["iTotalRecords"] = len(list)
	result["iTotalDisplayRecords"] = len(list)
	result["aaData"] = list

	//b, _ := json.Marshal(cnnJson)
	ctr.Data["json"] = result
	ctr.ServeJSON()
}
