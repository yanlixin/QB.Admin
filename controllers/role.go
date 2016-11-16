package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"QB.Admin/models/system"
	"QB.Admin/summaries"
	"QB.Admin/summaries/htmlhelper"
	"QB.Admin/summaries/jquery"
	"QB.Admin/utils"
)

//RoleController is role controller
type RoleController struct {
	BaseAdminRouter
}

//GetAllRoles of RoleController
func (ctl *RoleController) GetAllRoles() {
	//ctl.Data["xsrf_token"] = ctl.XSRFToken()
	ctl.Data["HasAdd"] = true
	ctl.TplName = "system/rolemanager/rolelist.html"

}

//GetRole of RoleController
func (ctl *RoleController) GetRole() {
	roleID, _ := strconv.Atoi(ctl.Ctx.Input.Param(":id"))
	r, _ := models.DefaultRoleList.Find(roleID)
	ctl.Data["Model"] = r
	ctl.TplName = "system/rolemanager/roleview.html"

}

//EditRole of RoleController
func (ctl *RoleController) EditRole() {

	roleID, _ := strconv.Atoi(ctl.Ctx.Input.Param(":id"))
	r, _ := models.DefaultRoleList.Find(roleID)
	if nil == r {
		r = new(models.Role)
	}
	ctl.Data["Model"] = r
	ctl.TplName = "system/rolemanager/roleedit.html"

}

//SaveRole of RoleController
func (ctl *RoleController) SaveRole() {
	var ob *models.Role
	err := json.Unmarshal(ctl.Ctx.Input.RequestBody, &ob)
	result := summaries.JsonResultSuccess("0", nil)
	if nil != err {
		utils.Logger.Error(fmt.Sprintf("%v", err))
		result = summaries.JsonResultUnmarshalFailed(fmt.Sprintf("%v", err), nil)
	} else {
		utils.Logger.Debug(fmt.Sprintf("%+v", ob))
		num, err := models.DefaultRoleList.Save(ob)
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

func (ctl *RoleController) RemoveRole() {
	roleID, _ := strconv.Atoi(ctl.Ctx.Input.Param(":id"))
	result := summaries.JsonResultSuccess("0", nil)
	if status := models.DefaultRoleList.Delete(roleID); status {
		result = summaries.JsonResultSuccess(fmt.Sprintf("%d", roleID), nil)
	} else {
		result = summaries.JsonResultLogicFailed("Delete role failed ", nil)
	}
	ctl.Data["json"] = result
	ctl.ServeJSON()
}
func (ctl *RoleController) GetRoleUsers() {

	roleID, _ := strconv.Atoi(ctl.Ctx.Input.Param(":id"))
	r, _ := models.DefaultRoleList.Find(roleID)
	if nil == r {
		r = new(models.Role)
	}
	ctl.Data["Model"] = r
	ctl.TplName = "system/rolemanager/roleusers.html"

}

//GetRoleUsersJSON of RoleController
func (ctl *RoleController) GetRoleUsersJSON() {
	roleID, _ := strconv.Atoi(ctl.Ctx.Input.Param(":id"))
	list, _ := models.DefaultUserList.GetUsersByRole(roleID)

	json := make(map[string]interface{})
	json["sEcho"] = 0
	json["iTotalRecords"] = len(list)
	json["iTotalDisplayRecords"] = len(list)
	json["aaData"] = list
	ctl.Data["json"] = json
	ctl.ServeJSON()
}

//GetAllRolesJSON of RoleController
func (ctl *RoleController) GetAllRolesJSON() {
	param := jquery.DataTablesParam{}
	if err := ctl.ParseForm(&param); err != nil {
		fmt.Println(err)
	}
	list, count := models.DefaultRoleList.GetAllRoles(&param)

	result := make(map[string]interface{})
	result["sEcho"] = param.Echo
	result["iTotalRecords"] = count
	result["iTotalDisplayRecords"] = count
	result["aaData"] = list

	ctl.Data["json"] = result

	ctl.ServeJSON()
}

//GetRolePermissions of RoleController
func (ctl *RoleController) GetRolePermissions() {

	roleID, _ := strconv.Atoi(ctl.Ctx.Input.Param(":id"))
	r, _ := models.DefaultRoleList.Find(roleID)
	if nil == r {
		r = new(models.Role)
	}
	ctl.Data["Model"] = r
	ctl.TplName = "system/rolemanager/rolepermissions.html"

}

//GetRolePermissionsJSON of RoleController
func (ctl *RoleController) GetRolePermissionsJSON() {

	roleID, _ := strconv.Atoi(ctl.Ctx.Input.Param(":id"))
	menus := models.DefaultMenuList.GetMenus(nil)
	var list []map[string]interface{}

	//rolePermissions := models.DefaultPermissionList.GetPermissionsByRole(roleID)
	rolePermissions := models.DefaultPermissionList.GetPermissionsByRole(roleID)
	for _, item := range menus {
		permissions, err := models.DefaultPermissionList.GetPermissionsByMenu(item.ID)
		if nil != err {
			fmt.Printf("%v", err)
		}
		if nil == permissions || 0 == len(permissions) {
			continue
		}
		var permissionList []map[string]interface{}
		for _, pitem := range permissions {
			p := make(map[string]interface{})
			p["Name"] = pitem.PermissionName
			p["Desc"] = pitem.PermissionDesc
			p["ID"] = pitem.PermissionID
			b := 0 != rolePermissions[pitem.ActionID]
			p["Checked"] = b

			permissionList = append(permissionList, p)
		}
		m := make(map[string]interface{})
		m["GroupID"] = item.ID
		m["GroupName"] = item.Name
		m["Permissions"] = permissionList
		list = append(list, m)
	}

	json := make(map[string]interface{})
	json["sEcho"] = 1
	json["iTotalRecords"] = len(list)
	json["iTotalDisplayRecords"] = len(list)
	json["aaData"] = list
	ctl.Data["json"] = json
	ctl.ServeJSON()
	//this.Data["Model"] = data
	//this.TplName = "system/rolemanager/roleview.ghtml"
}

//SaveRole of RoleController
func (ctl *RoleController) SaveRolePermissions() {
	form := ctl.Ctx.Input.Context.Request.Form
	roleID, _ := strconv.Atoi(form["RoleId"][0])
	permissionIds := strings.Split(form["PermissionIds"][0], ",")
	var perIds []int
	for _, str := range permissionIds {
		id, _ := strconv.Atoi(str)
		perIds = append(perIds, id)
	}
	result := summaries.JsonResultSuccess("0", nil)
	if status := models.DefaultRoleList.AddPermissions(roleID, perIds); status {
		result = summaries.JsonResultSuccess(fmt.Sprintf("%d", roleID), nil)
	} else {
		result = summaries.JsonResultLogicFailed("Delete role failed ", nil)
	}
	ctl.Data["json"] = result
	ctl.ServeJSON()
}
func (ctl *RoleController) AddRoleUsers() {

	roleID, _ := strconv.Atoi(ctl.Ctx.Input.Param(":id"))
	r, _ := models.DefaultRoleList.Find(roleID)
	if nil == r {
		r = new(models.Role)
	}
	var users []htmlhelper.SelectItem
	for _, item := range models.DefaultUserList.All() {
		selectItem := htmlhelper.SelectItem{strconv.Itoa(item.ID), item.Name, false}
		users = append(users, selectItem)
	}

	ctl.Data["Model"] = r
	ctl.Data["Users"] = users
	ctl.TplName = "system/rolemanager/roleuseradd.html"

}
func (ctl *RoleController) SaveRoleUsers() {
	form := ctl.Ctx.Input.Context.Request.Form
	roleID, _ := strconv.Atoi(form["RoleId"][0])
	userIds := strings.Split(form["UserIds"][0], ",")
	var uIds []int
	for _, str := range userIds {
		id, _ := strconv.Atoi(str)
		uIds = append(uIds, id)
	}
	result := summaries.JsonResultSuccess("0", nil)
	if status := models.DefaultRoleList.AddUsers(roleID, uIds); status {
		result = summaries.JsonResultSuccess(fmt.Sprintf("%d", roleID), nil)
	} else {
		result = summaries.JsonResultLogicFailed("Delete role failed ", nil)
	}
	ctl.Data["json"] = result
	ctl.ServeJSON()
}
func (ctl *RoleController) RemoveRoleUser() {
	form := ctl.Ctx.Input.Context.Request.Form
	roleID, _ := strconv.Atoi(form["RoleId"][0])
	userId, _ := strconv.Atoi(form["UserId"][0])

	result := summaries.JsonResultSuccess("0", nil)
	if status := models.DefaultRoleList.RemoveUser(roleID, userId); status {
		result = summaries.JsonResultSuccess(fmt.Sprintf("%d", roleID), nil)
	} else {
		result = summaries.JsonResultLogicFailed("Delete role failed ", nil)
	}
	ctl.Data["json"] = result
	ctl.ServeJSON()
}
