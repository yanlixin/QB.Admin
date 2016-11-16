package databases

import (
	"fmt"

	"QB.Admin/utils"
	"github.com/astaxie/beego/orm"
)

func (u *DbContent) GetAllModules() ([]*Modules, bool) {
	status := true
	o := orm.NewOrm()
	//var user Users
	var result []*Modules

	_, err := o.QueryTable("pa_modules").All(&result)
	if nil != err {
		status = false
		utils.Logger.Error(fmt.Sprintf("%v", err))
	}
	return result, status
}
