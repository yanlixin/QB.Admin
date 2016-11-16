package routers

import (
	"QB.Admin/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.DefaultController{})

	beego.Router("/index", &controllers.DefaultController{}, "GET:Get")
	beego.Router("/home", &controllers.DefaultController{}, "GET:GetHome")
	beego.Router("/getuserinfo", &controllers.UserController{}, "*:GetUserInfo")

	beego.Router("/login", &controllers.LoginController{}, "GET:Index")
	beego.Router("/logout", &controllers.LoginController{}, "GET:Logout")
	beego.Router("/login", &controllers.LoginController{}, "POST:Login")
	beego.Router("/user/profile", &controllers.UserController{}, "POST:GetProfile")
	beego.Router("/user/profile/save", &controllers.UserController{}, "POST:SaveProfile")
	beego.Router("/user/settings", &controllers.UserController{}, "GET:GetSettings")

	beego.Router("/public/orgunit/list/json", &controllers.CommonController{}, "GET:GetOrgUnitJSON")
	beego.Router("/public/employee/list/json/?:orgunitid", &controllers.CommonController{}, "GET:GetEmployeeJSON")

	beego.Router("/system/users/list", &controllers.UserController{}, "*:GetAllUsers")
	beego.Router("/system/users/list/json", &controllers.UserController{}, "POST:GetAllUsersJSON")
	beego.Router("/system/user/edit/?:id", &controllers.UserController{}, "GET:EditUser")
	beego.Router("/system/user/view/:id", &controllers.UserController{}, "GET:GetUser")
	beego.Router("/system/user/save", &controllers.UserController{}, "POST:SaveUser")
	beego.Router("/system/user/delete/:id", &controllers.UserController{}, "PUT:RemoveUser")
	beego.Router("/system/user/changepwd/?:id", &controllers.UserController{}, "GET:ChangeUserPwd")
	beego.Router("/system/user/savepwd", &controllers.UserController{}, "POST:SaveUserPwd")

	beego.Router("/system/permission/list", &controllers.PermissionController{}, "POST:GetAllPermissions")
	beego.Router("/system/permission/list/json", &controllers.PermissionController{}, "POST:GetAllPermissionsJSON")

	beego.Router("/system/roles/list", &controllers.RoleController{}, "*:GetAllRoles")
	beego.Router("/system/roles/list/json", &controllers.RoleController{}, "POST:GetAllRolesJSON")
	beego.Router("/system/role/edit/?:id", &controllers.RoleController{}, "GET:EditRole")
	beego.Router("/system/role/view/:id", &controllers.RoleController{}, "GET:GetRole")
	beego.Router("/system/role/user/list/?:id", &controllers.RoleController{}, "*:GetRoleUsers")
	beego.Router("/system/role/user/list/json/?:id", &controllers.RoleController{}, "*:GetRoleUsersJSON")
	beego.Router("/system/role/user/add/:id", &controllers.RoleController{}, "GET:AddRoleUsers")
	beego.Router("/system/role/user/save", &controllers.RoleController{}, "POST:SaveRoleUsers")
	beego.Router("/system/role/user/delete", &controllers.RoleController{}, "PUT:RemoveRoleUser")

	beego.Router("/system/role/users/json", &controllers.RoleController{}, "GET:GetRoleUsersJSON")
	beego.Router("/system/role/save", &controllers.RoleController{}, "POST:SaveRole")
	beego.Router("/system/role/delete/:id", &controllers.RoleController{}, "PUT:RemoveRole")
	beego.Router("/system/role/permission/list/?:id", &controllers.RoleController{}, "GET:GetRolePermissions")
	beego.Router("/system/role/permission/list/json/?:id", &controllers.RoleController{}, "POST:GetRolePermissionsJSON")
	beego.Router("/system/role/permission/save", &controllers.RoleController{}, "POST:SaveRolePermissions")

	beego.Router("/system/menu/list", &controllers.MenuController{}, "POST:GetAllMenus")
	beego.Router("/system/menu/flat/json", &controllers.MenuController{}, "*:GetFlatMenusJSON")
	beego.Router("/system/menu/edit/?:id", &controllers.MenuController{}, "GET:EditMenu")
	beego.Router("/system/menu/view/:id", &controllers.MenuController{}, "GET:GetMenu")
	beego.Router("/system/menu/save", &controllers.MenuController{}, "POST:SaveMenu")
	beego.Router("/system/menu/delete/:id", &controllers.MenuController{}, "PUT:RemoveMenu")

	beego.Router("/system/model/view", &controllers.UserController{}, "POST:GetModel")

	beego.Router("/system/model/json/:tableName", &controllers.UserController{}, "Get:GetModelJSON")

	beego.Router("/weixin/mp/index", &controllers.WeixinController{}, "*:Index")
	beego.Router("/weixin/mp/message/list", &controllers.WeixinController{}, "POST:GetAllMessages")
	beego.Router("/weixin/mp/message/edit/?:id", &controllers.WeixinController{}, "GET:EditMessage")
	beego.Router("/weixin/mp/message/send", &controllers.WeixinController{}, "POST:SendMessage")

}
