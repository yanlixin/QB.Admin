package models

import (
	"../utils"
	"QB.Admin/summaries/jquery"
	"fmt"
	"testing"
	"time"
)

func TestGetAllRoles(t *testing.T) {
	param := jquery.DataTablesParam{}
	param.DisplayLength = 10
	param.DisplayStart = 0
	actual, _ := DefaultRoleList.GetAllRoles(&param)

	if len(actual) == 0 {

		t.Errorf("GetAllRoles is excepted more then 0 actualed %v", len(actual))
	}
}
func TestGetRolesByUser(t *testing.T) {
	actual, err := DefaultRoleList.GetRolesByUser(1)
	if nil != err {
		t.Errorf("%v", err)
	}
	if len(actual) == 0 {

		t.Errorf("GetRolesByUser is excepted more then 0 actualed %v", len(actual))
	}
}

func TestInsertRole(t *testing.T) {
	m := DefaultRoleList
	r := new(Role)
	r.RoleDesc = "Model Test master"

	r.RoleName = fmt.Sprintf("%s%s", "master", time.Now())
	_, err := m.Save(r)
	if nil != err {
		t.Errorf("Insert role has error %v", err)
	}

}
func TestUpdateRole(t *testing.T) {
	param := new(jquery.DataTablesParam)
	param.DisplayLength = 15
	param.DisplayStart = 0
	m := DefaultRoleList
	roles, _ := m.GetAllRoles(param)
	r := roles[len(roles)-1]
	if r.RecordStatus == 0 {
		r.RecordStatus = -1
	} else {
		r.RecordStatus = 0
	}
	num, err := m.Save(r)
	if nil != err {
		t.Errorf("Update role has error %v", err)
	}
	if num != 1 {
		t.Errorf(utils.TestMessageLength, 1, num)
	}
}
func TestDeleteRole(t *testing.T) {
	param := new(jquery.DataTablesParam)
	param.DisplayLength = 15
	param.DisplayStart = 0
	m := DefaultRoleList
	roles, _ := m.GetAllRoles(param)
	actual := m.Delete(roles[len(roles)-1].RoleID)
	if !actual {
		t.Errorf(utils.TestMessageBool, true, actual)
	}
}
