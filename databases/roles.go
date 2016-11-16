package databases

import (
	"fmt"
	"time"

	"QB.Admin/utils"
	"github.com/astaxie/beego/orm"
)

func (u *DbContent) GetAllRoles(pageSize int, pageStart int) ([]*Roles, int64) {

	o := orm.NewOrm()
	//var user Users
	var roles []*Roles
	count, _ := o.QueryTable("pa_roles").Count()
	_, err := o.QueryTable("pa_roles").Limit(pageSize, pageStart).All(&roles)
	if nil != err {

		utils.Logger.Error(fmt.Sprintf("%v", err))
	}

	return roles, count
}
func (c *DbContent) GetRole(roleID int) (*Roles, bool) {
	status := true
	o := orm.NewOrm()
	var role Roles
	err := o.QueryTable("pa_roles").Filter("RoleID", roleID).One(&role)
	if nil != err {
		status = false
		utils.Logger.Error(fmt.Sprintf("%v", err))
	}
	return &role, status
}
func (db *DbContent) InsertRole(r *Roles) (int64, bool) {
	status := true
	r.CreatedDate = time.Now()
	o := orm.NewOrm()
	num, err := o.Insert(r)
	if nil != err {
		status = false
		utils.Logger.Error(fmt.Sprintf("%v", err))
	}
	return num, status
}
func (db *DbContent) UpdateRole(r *Roles) (int64, bool) {
	o := orm.NewOrm()
	num, err := o.Update(r, "RecordStatus", "RoleName", "RoleDesc", "LastUpdatedByUserID")
	if err != nil {
		utils.Logger.Error(fmt.Sprintf("%v", err))
		return num, false
	}
	return num, true
}
func (db *DbContent) DeleteRole(roleID int) (int64, bool) {
	result := true
	o := orm.NewOrm()
	var r = new(Roles)
	r.RoleID = roleID
	r.RecordStatus = -1
	num, err := o.Update(r, "RecordStatus")
	if err != nil {
		utils.Logger.Error(fmt.Sprintf("%v", err))
		result = false
	}
	return num, result
}
func (db *DbContent) GetRolesByUser(userID int) ([]*Roles, bool) {
	o := orm.NewOrm()
	//var user Users

	var roles []*Roles
	_, err := o.Raw("select a.* from pa_roles a left join pa_r_users_roles b on a.roleid=b.roleid where b.userid=?", userID).QueryRows(&roles)
	if nil != err {
		utils.Logger.Error(fmt.Sprintf("%v", err))
		return nil, false
	}

	return roles, true
}

func (db *DbContent) AddUsersToRole(roleID int, userIDs []int) bool {
	o := orm.NewOrm()
	err := o.Begin()
	if nil != err {
		utils.Logger.Error(fmt.Sprintf("%v", err))
		return false
	}
	var list []*R_Users_Roles
	for _, id := range userIDs {
		var ob = new(R_Users_Roles)
		ob.UserID = id
		ob.RoleID = roleID
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
func (db *DbContent) DeleteUsersFromRole(roleID int, userIDs []int) bool {
	o := orm.NewOrm()
	err := o.Begin()
	if nil != err {
		utils.Logger.Error(fmt.Sprintf("%v", err))
		return false
	}
	//_, err = o.Raw("DELETE pa_r_users_roles WHERE UserID in (?,?,?) and RoleID=?", userIDs, roleID).Exec()
	_, err = o.QueryTable("pa_r_users_roles").Filter("RoleID", roleID).Filter("UserID__in", userIDs).Delete()
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
func (db *DbContent) AddPermissionsToRole(roleID int, permissionIDs []int) bool {
	o := orm.NewOrm()
	err := o.Begin()
	if nil != err {
		utils.Logger.Error(fmt.Sprintf("%v", err))
		return false
	}
	_, err = o.QueryTable("pa_r_roles_permissions").Filter("RoleID", roleID).Delete()
	if nil != err {
		utils.Logger.Error(fmt.Sprintf("%v", err))
		o.Rollback()
		return false
	}
	var list []*R_Roles_Permissions
	for _, id := range permissionIDs {
		var ob = new(R_Roles_Permissions)
		ob.PermissionID = id
		ob.RoleID = roleID
		list = append(list, ob)
	}
	_, err = o.InsertMulti(10000, list)
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
