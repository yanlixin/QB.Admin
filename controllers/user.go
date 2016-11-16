package controllers

import (
	"QB.Admin/models/system"
	"QB.Admin/summaries"
	"QB.Admin/summaries/jquery"
	"QB.Admin/utils"
	//"encoding/json"
	"fmt"
	"strconv"
	"time"
)

//UserController is user controller
type UserController struct {
	BaseAdminRouter
}

//Get of UserController
func (ctl *UserController) Get() {
	//	m := NewUserManager()
	//	c.Data["userList"] = m.All()
	//	c.TplName = "userlist.ghtml"
	//	c.Layout = "layout.ghtml"

}

//GetUserInfo of UserController
func (ctl *UserController) GetUserInfo() {
	allMenus, _ := models.BuildMenu(0, ctl.CurrentUser.Menus)
	var result = make(map[string]interface{})
	result["User"] = ctl.CurrentUser
	result["Menus"] = allMenus

	ctl.Data["json"] = result
	ctl.ServeJSON()
}

//GetAllUsers of UserController
func (ctl *UserController) GetAllUsers() {
	ctl.TplName = "system/usermanager/userlist.html"

}

//GetUser of UserController
func (ctl *UserController) GetUser() {
	userID, _ := strconv.Atoi(ctl.Ctx.Input.Param(":id"))
	u, _ := models.DefaultUserList.Find(userID)
	if roles, err := models.DefaultRoleList.GetRolesByUser(userID); err == nil {
		u.Roles = roles
	}

	ctl.Data["Model"] = u

	ctl.TplName = "system/usermanager/userview.html"

}
func (ctl *UserController) GetProfile() {
	//userID, _ := strconv.Atoi(ctl.Ctx.Input.Param(":id"))
	//u, _ := models.DefaultUserList.Find(userID)
	//if roles, err := models.DefaultRoleList.GetRolesByUser(userID); err == nil {
	//	u.Roles = roles
	//}

	ctl.Data["Model"] = ctl.CurrentUser

	ctl.TplName = "user/profile.html"

}
func (ctl *UserController) SaveProfile() {
	result := summaries.JsonResultSuccess("0", nil)
	fmt.Printf("%+v", ctl.Ctx.Request.Form)
	oldPassword := ctl.Ctx.Request.Form["OldPassword"][0]
	newPassword := ctl.Ctx.Request.Form["UserPassword"][0]
	if len(newPassword) > 0 && utils.GenPwdStr(oldPassword) != ctl.CurrentUser.Password {
		result = summaries.JsonResultLogicFailed("密码验证失败，请重新输入旧密码！", nil)
	} else {
		if len(newPassword) > 0 {
			_, err := models.DefaultUserList.ChangePassword(ctl.CurrentUser.ID, newPassword)
			if nil == err {
				result = summaries.JsonResultSuccess("保存成功!", nil)
			} else {
				utils.Logger.Error(fmt.Sprintf("%v", err))
				result = summaries.JsonResultLogicFailed(fmt.Sprintf("%v", err), nil)
			}
		}
		//fmt.Printf("%+v", utils.GenPwdStr(oldPassword))
		/*
			if err := ctl.ParseForm(&ob); err != nil {
				fmt.Println(err)
			}


			utils.Logger.Debug(fmt.Sprintf("%+v", ob))
			num, err := models.DefaultUserList.Save(&ob)
			if nil == err {
				result = summaries.JsonResultSuccess(fmt.Sprintf("%d", num), nil)
			} else {
				utils.Logger.Error(fmt.Sprintf("%v", err))
				result = summaries.JsonResultLogicFailed(fmt.Sprintf("%v", err), nil)
			}
		*/
	}
	ctl.Data["json"] = result
	ctl.ServeJSON()

}
func (ctl *UserController) GetSettings() {

	//ctl.Data["Model"] = u

	ctl.TplName = "user/settings.html"

}

//EditUser of UserController
func (ctl *UserController) EditUser() {

	userID, _ := strconv.Atoi(ctl.Ctx.Input.Param(":id"))
	u, _ := models.DefaultUserList.Find(userID)
	if nil == u {
		u = new(models.User)
		u.AccountExpirationDate = time.Now()
	}
	fmt.Printf("%+v", u)
	ctl.Data["Model"] = u
	ctl.Data["IsEdit"] = userID > 0

	ctl.TplName = "system/usermanager/useredit.html"

}
func (ctl *UserController) ChangeUserPwd() {

	userID, _ := strconv.Atoi(ctl.Ctx.Input.Param(":id"))
	u, _ := models.DefaultUserList.Find(userID)
	ctl.Data["Model"] = u
	ctl.TplName = "system/usermanager/userpwd.html"

}
func (ctl *UserController) SaveUserPwd() {
	//var ob *models.User
	ob := models.User{}

	if err := ctl.ParseForm(&ob); err != nil {
		fmt.Println(err)
	}
	result := summaries.JsonResultSuccess("0", nil)

	num, err := models.DefaultUserList.ChangePassword(ob.ID, ob.Password)
	if nil == err {
		result = summaries.JsonResultSuccess(fmt.Sprintf("%d", num), nil)
	} else {
		utils.Logger.Error(fmt.Sprintf("%v", err))
		result = summaries.JsonResultLogicFailed(fmt.Sprintf("%v", err), nil)
	}
	//}
	ctl.Data["json"] = result
	ctl.ServeJSON()
}

//SaveUser of UserController
func (ctl *UserController) SaveUser() {
	//var ob *models.User
	ob := models.User{}

	if err := ctl.ParseForm(&ob); err != nil {
		fmt.Println(err)
	}
	result := summaries.JsonResultSuccess("0", nil)

	utils.Logger.Debug(fmt.Sprintf("%+v", ob))
	num, err := models.DefaultUserList.Save(&ob)
	if nil == err {
		result = summaries.JsonResultSuccess(fmt.Sprintf("%d", num), nil)
	} else {
		utils.Logger.Error(fmt.Sprintf("%v", err))
		result = summaries.JsonResultLogicFailed(fmt.Sprintf("%v", err), nil)
	}
	//}
	ctl.Data["json"] = result
	ctl.ServeJSON()
}
func (ctl *UserController) RemoveUser() {
	userID, _ := strconv.Atoi(ctl.Ctx.Input.Param(":id"))
	result := summaries.JsonResultSuccess("0", nil)
	if status := models.DefaultUserList.Delete(userID); status {
		result = summaries.JsonResultSuccess(fmt.Sprintf("%d", userID), nil)
	} else {
		utils.Logger.Error(fmt.Sprintf("Delete User fail by UserID %v", userID))
		result = summaries.JsonResultLogicFailed(fmt.Sprintf("删除用户失败"), nil)
	}
	ctl.Data["json"] = result
	ctl.ServeJSON()
}

//GetAllUsersJSON of UserController
func (ctl *UserController) GetAllUsersJSON() {

	param := jquery.DataTablesParam{}
	if err := ctl.ParseForm(&param); err != nil {
		fmt.Println(err)
	}

	users, count := models.DefaultUserList.GetAllUsers(&param)

	result := make(map[string]interface{})
	result["sEcho"] = param.Echo
	result["iTotalRecords"] = count
	result["iTotalDisplayRecords"] = count
	result["aaData"] = users

	ctl.Data["json"] = result
	ctl.ServeJSON()
}
