package models

import (
	"fmt"
	"testing"
	"time"

	"QB.Admin/summaries/jquery"
	"QB.Admin/utils"
)

func TestGetAllUsers(t *testing.T) {
	param := new(jquery.DataTablesParam)
	param.DisplayLength = 15
	param.DisplayStart = 15
	m := DefaultUserList
	users, _ := m.GetAllUsers(param)
	actual := len(users)
	if actual == 0 {

		t.Errorf(utils.TestMessageLength, 0, actual)
	}
	//utils.Logger.Debug(fmt.Sprintf("%+v", users[1]))
	if users[1] == nil {

		t.Errorf("Target user is nil")
	}
}
func TestFindUser(t *testing.T) {
	m := DefaultUserList
	user, _ := m.Find(1)
	if user == nil {
		t.Errorf(utils.TestMessageLength, 1, 0)
	}
	//if 0 == len(user.Roles) {
	//	t.Errorf(utils.TestMessageLength, 2, len(user.Roles))
	//}
}
func TestFindUser4NoneUser(t *testing.T) {
	m := DefaultUserList
	user, _ := m.Find(0)
	if user != nil {
		t.Errorf("Except is nil actual is %v", user)
	}
}
func TestFindUserByLoginName(t *testing.T) {
	m := DefaultUserList
	_, b := m.FindByLoginName("master")
	if !b {
		t.Errorf(utils.TestMessageBool, true, b)
	}
}

func TestInsertUserDupLoginName(t *testing.T) {
	m := DefaultUserList
	u := new(User)
	u.UserLoginName = "master"
	u.UserDisplayName = "Model Test master"

	_, err := m.Save(u)
	if nil == err {
		t.Errorf("Insert User check user loginname error %v", err)
	}

}

func TestInsertUser(t *testing.T) {
	m := DefaultUserList
	u := new(User)
	u.UserDisplayName = "Model Test master"

	u.UserLoginName = fmt.Sprintf("%s_%s", "master", time.Now())
	_, err := m.Save(u)
	if nil != err {
		t.Errorf("Insert Users has error %v", err)
	}

}
func TestUpdateUser(t *testing.T) {
	m := DefaultUserList
	u, b := m.Find(12)
	if !b {
		t.Errorf(utils.TestMessageBool, true, b)
	}
	if u.RecordStatus == 0 {
		u.RecordStatus = -1
	} else {
		u.RecordStatus = 0
	}
	num, err := m.Save(u)

	if nil != err {
		t.Errorf("Update Users has error %v", err)
	}
	if num != 12 {
		t.Errorf(utils.TestMessageLength, 1, num)
	}
}
func TestDeleteUser(t *testing.T) {
	actual := DefaultUserList.Delete(12)
	if !actual {
		t.Errorf(utils.TestMessageBool, true, actual)
	}
}
func TestGetUsersByRole(t *testing.T) {
	actual, b := DefaultUserList.GetUsersByRole(1)
	if !b {
		t.Errorf(utils.TestMessageBool, true, b)
	}
	if nil == actual || 0 == len(actual) {
		t.Errorf(utils.TestMessageLength, 1, len(actual))
	}
}

/*
func TestNewUser(t *testing.T) {
/*
func TestNewUser(t *testing.T) {
	title := "learn Go"
	user := newUserOrFatal(t, title)
	if user.UserName != title {
		t.Errorf("expected title %q, got %q", title, user.UserName)
	}
	if user.OnLine {
		t.Errorf("new user is online")
	}
	if 3 != len(user.Menus) {
		t.Errorf("new user menu miss")
	}
}

func TestNewUserEmptyUserName(t *testing.T) {
	_, err := NewUser("")
	if err == nil {
		t.Errorf("expected 'empty UserName' error, got nil")
	}
}

func TestSaveUserAndRetrieve(t *testing.T) {
	user := newUserOrFatal(t, "learn Go")

	m := NewUserManager()
	m.Save(user)

	all := m.All()
	if len(all) != 1 {
		t.Errorf("expected 1 User, got %v", len(all))
	}
}
*/
