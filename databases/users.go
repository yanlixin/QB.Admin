package databases

import (
	"fmt"
	"time"

	"QB.Admin/utils"
	"github.com/astaxie/beego/orm"
)

func (c *DbContent) LoginUser(loginName string, loginPwd string) (bool, *Users) {
	dk := utils.GenPwdStr(loginPwd)

	o := orm.NewOrm()
	var user []*Users
	qs := o.QueryTable("pa_users")

	//_, err := o.QueryTable("pa_users").All(&user)
	qs.Filter("RecordStatus", 0).Filter("UserLoginName", loginName).Filter("UserPassword", dk).All(&user)
	if 1 == len(user) {
		utils.Logger.Debug(fmt.Sprintf("%+v", user))
		return true, user[0]
	}

	return false, nil
}
func (c *DbContent) GetUser(userID int) (*Users, bool) {
	o := orm.NewOrm()
	var user Users
	o.QueryTable("pa_users").Filter("UserID", userID).One(&user)
	return &user, true
}
func (u *DbContent) GetAllUsers(pageSize int, pageStart int, sortCol string, sortDir string) ([]*Users, int64) {
	o := orm.NewOrm()
	//var user User
	var users []*Users
	count, _ := o.QueryTable("pa_users").Count()
	if "id" == sortCol || "" == sortCol {
		sortCol = "userid"
	}
	if "asc" != sortDir {
		sortCol = "-" + sortCol
	}
	o.QueryTable("pa_users").Limit(pageSize, pageStart).OrderBy(sortCol).All(&users)
	return users, count
}
func (db *DbContent) InsertUser(u *Users) (int64, bool) {
	status := true
	prpareSave(u)
	u.UserPassword = utils.GenPwdStr(u.UserPassword)
	u.CreatedDate = time.Now()
	o := orm.NewOrm()
	num, err := o.Insert(u)
	if nil != err {
		status = false
		utils.Logger.Error(fmt.Sprintf("%v", err))
	}
	return num, status
}
func (db *DbContent) ChangeUserPassword(u *Users) (int64, bool) {
	o := orm.NewOrm()
	u.UserPassword = utils.GenPwdStr(u.UserPassword)
	_, err := o.Update(u, "UserPassword", "LastUpdatedByUserID")
	if err != nil {
		utils.Logger.Error(fmt.Sprintf("%v", err))
		return -1, false
	}
	return int64(u.UserID), true
}
func (db *DbContent) UpdateUser(u *Users) (int64, bool) {
	o := orm.NewOrm()
	prpareSave(u)
	o.Begin()
	_, err := o.Update(u, "RecordStatus", "UserDisplayName", "UserAccountExpirationDate", "LastUpdated", "LastUpdatedByUserID", "UserMustChangePwd", "UserCannotChangePwd")
	if err != nil {
		o.Rollback()
		utils.Logger.Error(fmt.Sprintf("%v", err))
		return -1, false
	}
	o.Commit()
	return int64(u.UserID), true
}
func (db *DbContent) DeleteUser(userID int) (int64, bool) {
	result := true
	o := orm.NewOrm()
	var u = new(Users)
	u.UserID = userID
	prpareSave(u)
	u.RecordStatus = -1
	_, err := o.Update(u, "RecordStatus", "LastUpdated", "LastUpdatedByUserID")
	if err != nil {
		utils.Logger.Error(fmt.Sprintf("%v", err))
		result = false
	}
	return int64(userID), result
}
func (db *DbContent) GetUsersByRole(roleID int) ([]*Users, bool) {
	o := orm.NewOrm()
	//var user Users

	var users []*Users
	_, err := o.Raw("select a.* from pa_users a left join pa_r_users_roles b on a.userid=b.userid where b.roleid=?", roleID).QueryRows(&users)
	if nil != err {
		utils.Logger.Error(fmt.Sprintf("%v", err))
		return nil, false
	}

	return users, true
}
func (db *DbContent) GetRolesUsers() ([]*R_Users_Roles, bool) {
	o := orm.NewOrm()
	//var user Users
	var result []*R_Users_Roles

	_, err := o.QueryTable("pa_r_users_roles").All(&result)
	if nil != err {
		utils.Logger.Error(fmt.Sprintf("%v", err))
		return nil, false
	}
	return result, true
}
func (db *DbContent) AddRolesToUser(userID int, roleIDs []int) bool {
	o := orm.NewOrm()
	err := o.Begin()
	if nil != err {
		utils.Logger.Error(fmt.Sprintf("%v", err))
		return false
	}
	var list []*R_Users_Roles
	for _, id := range roleIDs {
		var ob = new(R_Users_Roles)
		ob.UserID = userID
		ob.RoleID = id
		list = append(list, ob)
	}
	_, err = o.InsertMulti(1000, list)
	if nil != err {
		utils.Logger.Error(fmt.Sprintf("%v", err))
		o.Rollback()
		return false
	}
	err = o.Commit()
	if nil != err {
		utils.Logger.Error(fmt.Sprintf("%v", err))
		return false
	}
	return true
}
func (db *DbContent) DeleteRolesFromUser(userID int, roleIDs []int) bool {
	o := orm.NewOrm()
	err := o.Begin()
	if nil != err {
		utils.Logger.Error(fmt.Sprintf("%v", err))
		return false
	}
	//	_, err = o.Raw("DELETE pa_r_users_roes WHERE UserID=? and RoleID in (?,?,?)", userID, roleIDs).Exec()
	_, err = o.QueryTable("pa_r_users_roles").Filter("UserID", userID).Filter("RoleID__in", roleIDs).Delete()
	if nil != err {
		utils.Logger.Error(fmt.Sprintf("%v", err))
		o.Rollback()
		return false
	}
	err = o.Commit()
	if nil != err {
		utils.Logger.Error(fmt.Sprintf("%v", err))
		return false
	}
	return true
}
