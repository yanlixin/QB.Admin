package databases

import (
	"fmt"
	"strconv"

	"QB.Admin/utils"
	"github.com/astaxie/beego/orm"
)

func (u *DbContent) GetAllPermissions(menuID int) ([]*Permissions, bool) {
	status := true
	o := orm.NewOrm()
	var result []*Permissions
	if menuID > 0 {
		_, err := o.QueryTable("pa_permissions").OrderBy("PermissionID").Filter("MenuID", menuID).All(&result)
		if nil != err {
			status = false
			utils.Logger.Error(fmt.Sprintf("%v", err))
		}
	} else {
		_, err := o.QueryTable("pa_permissions").OrderBy("PermissionID").All(&result)
		if nil != err {
			status = false
			utils.Logger.Error(fmt.Sprintf("%v", err))
		}
	}
	return result, status
}

func (db *DbContent) GetRolePermissions() (map[int]map[string]int, bool) {
	o := orm.NewOrm()
	var lists []orm.ParamsList
	result := make(map[int]map[string]int)
	queryStr := `SELECT a.RoleID,a.PermissionID,b.ActionID
			FROM pa_r_roles_permissions a left join pa_permissions b on a.permissionid=b.permissionid`
	num, err := o.Raw(queryStr).ValuesList(&lists)
	if err == nil && num > 0 {
		for _, item := range lists {
			roleID, _ := strconv.Atoi(fmt.Sprintf("%v", item[0]))
			permissionID, _ := strconv.Atoi(item[1].(string))
			actionID := item[2].(string)

			if nil == result[roleID] {
				result[roleID] = make(map[string]int)
			}
			result[roleID][actionID] = permissionID
		}
	}
	return result, true
}
