package controllers

import (
	"QB.Admin/models/system"
)

//PermissionController is permission controller
type PermissionController struct {
	BaseAdminRouter
}

//GetAllPermissions of PermissionController
func (ctl *PermissionController) GetAllPermissions() {
	ctl.Data["HasAdd"] = true
	ctl.Data["HasView"] = true
	ctl.TplName = "system/permissionmanager/permissionlist.html"

}

//GetAllPermissionsJSON of PermissionController
func (ctl *PermissionController) GetAllPermissionsJSON() {

	permissions := models.DefaultPermissionList.GetPermissions()
	cnnJSON := make(map[string]interface{})
	cnnJSON["sEcho"] = 1
	cnnJSON["iTotalRecords"] = len(permissions)
	cnnJSON["iTotalDisplayRecords"] = len(permissions)
	cnnJSON["aaData"] = permissions

	//b, _ := json.Marshal(cnnJson)
	ctl.Data["json"] = cnnJSON
	ctl.ServeJSON()
}
