package models

import (
	"errors"
	"fmt"
	"time"

	"QB.Admin/databases"
	"QB.Admin/summaries/jquery"
	"QB.Admin/utils"
)

func init() {
	DefaultUserList = NewUserManager()
}

var DefaultUserList *UserManager

type User struct {
	ID          int    // Unique identifier
	Name        string // Description
	LoginName   string
	DisplayName string
	Password    string
	Email       string
	//Roles                 string
	AccountExpirationDate time.Time
	Photo                 string
	Signature             string
	MustChangePwd         string
	CannotChangePwd       bool
	RecordStatus          int
	OnLine                bool // Is this task done?
	Menus                 []*Menu
	Roles                 []*Role
	Permissions           map[string]int
}

func fillUser(u *databases.Users) (*User, error) {

	if u.UserLoginName == "" {
		return nil, errors.New("empty UserLoginName")
	}
	//var menus []*Menu
	result := new(User)
	result.ID = u.UserID
	result.LoginName = u.UserLoginName
	result.DisplayName = u.UserDisplayName
	result.Name = fmt.Sprintf("[%s] %s", u.UserLoginName, u.UserDisplayName)
	result.AccountExpirationDate = u.UserAccountExpirationDate
	if u.UserMustChangePwd {
		result.MustChangePwd = "checked"
	} else {
		result.MustChangePwd = ""
	}
	result.CannotChangePwd = u.UserCannotChangePwd
	result.RecordStatus = u.RecordStatus
	result.Password = u.UserPassword
	return result, nil
}
func (this *User) IsLogin() bool {
	return true
}
func Login(loginName string, loginPwd string) (bool, *User) {
	dbContent := databases.Instance()
	b, u := dbContent.LoginUser(loginName, loginPwd)
	var user *User
	if b == true {
		user, _ = fillUser(u)
		roles, _ := DefaultRoleList.GetRolesByUser(user.ID)
		user.Roles = roles
		permissions := make(map[string]int)
		for _, role := range roles {
			utils.MapMerge(permissions, role.Permissions)
		}
		user.Permissions = permissions
		menus := DefaultMenuList.GetMenus(nil)
		user.Menus = []*Menu{&Menu{ID: -1}}
		if user.ID == 1 {
			user.Menus = menus
		} else {

			for _, menu := range menus {
				if len(menu.ActionID) == 0 {
					user.Menus = append(user.Menus, menu)
					continue
				}
				if _, ok := permissions[menu.ActionID]; ok {

					user.Menus = append(user.Menus, menu)
				}
			}
		}
	}

	//user.Menus=

	return b, user
}

// cloneTask creates and returns a deep copy of the given Task.
func cloneUser(t *User) *User {
	c := *t
	return &c
}

// TaskManager manages a list of tasks in memory.
type UserManager struct {
	users  []*User
	lastID int
}

// NewUserManager returns an empty UserManager.
func NewUserManager() *UserManager {
	result := &UserManager{}
	result.fillAll()
	return result
}
func (m *UserManager) fillAll() error {
	var result []*User
	users, _ := databases.DefaultContent.GetAllUsers(-1, 0, "", "")
	for _, u := range users {
		user, _ := fillUser(u)
		result = append(result, user)
	}
	m.users = result
	return nil
}
func (m *UserManager) CheckUserLoginName(userLoginName string) (string, bool) {

	var msg string
	var status bool = true
	if _, b := m.FindByLoginName(userLoginName); b {
		msg = "登录名称重复"
		status = false
	}
	return msg, status
}
func (m *UserManager) CheckUserPassword(password string) (string, bool) {
	return "", true
}
func (m *UserManager) ChangePassword(userID int, userPassword string) (int, error) {

	u, b := databases.DefaultContent.GetUser(userID)
	if !b {

		err := errors.New("Find User Fail")
		return 0, err
	}
	if msg, b := m.CheckUserPassword(userPassword); !b {
		//密码不符合要求
		err := errors.New(msg)
		return 0, err
	}
	u.UserPassword = userPassword

	databases.DefaultContent.ChangeUserPassword(u)

	m.sync(int(userID))

	return userID, nil
}

// Save saves the given Task in the TaskManager.
func (m *UserManager) Save(user *User) (int64, error) {
	isNew := true
	var u *databases.Users
	if 0 == user.ID {
		isNew = true
		u = new(databases.Users)
		u.UserID = 0
		u.UserLoginName = user.LoginName
		if msg, b := m.CheckUserLoginName(user.LoginName); !b {
			err := errors.New(msg)
			return 0, err
		}
	} else {
		isNew = false

		var b bool
		u, b = databases.DefaultContent.GetUser(user.ID)
		if !b {

			err := errors.New("Find User Fail")
			return 0, err
		}

		if 0 >= len(user.Password) {
			u.UserPassword = user.Password
			if msg, b := m.CheckUserPassword(user.Password); !b {
				//密码不符合要求
				err := errors.New(msg)
				return 0, err
			}
		}

	}

	u.UserDisplayName = user.DisplayName
	u.UserAccountExpirationDate = user.AccountExpirationDate
	if "on" == user.MustChangePwd {
		u.UserMustChangePwd = true
	} else {
		u.UserMustChangePwd = false
	}
	u.UserCannotChangePwd = user.CannotChangePwd
	u.RecordStatus = user.RecordStatus
	//var err error
	var userID int64
	if isNew {
		userID, _ = databases.DefaultContent.InsertUser(u)
	} else {
		userID, _ = databases.DefaultContent.UpdateUser(u)

	}
	m.sync(int(userID))
	return userID, nil
}
func (m *UserManager) sync(userID int) {

	isFound := false
	user, b := databases.DefaultContent.GetUser(userID)
	if !b || nil == user {
		errors.New("获取用户信息失败!")
		return
	}
	newUser, _ := fillUser(user)
	for i, u := range m.users {
		if u.ID == userID {

			m.users[i] = newUser
			isFound = true
			break
		}
	}
	if !isFound {
		m.users = append(m.users, newUser)
	}

}
func (m *UserManager) Delete(userID int) bool {
	_, b := databases.DefaultContent.DeleteUser(userID)
	m.sync(userID)
	return b
}

// All returns the list of all the Tasks in the TaskManager.
func (m *UserManager) All() []*User {
	return m.users
}
func (m *UserManager) GetAllUsers(param *jquery.DataTablesParam) ([]*User, int64) {
	//var v *jquery.DataTablesParam
	var result []*User
	//v.iDisplayStart = 1
	users, count := databases.DefaultContent.GetAllUsers(param.DisplayLength, param.DisplayStart, param.GetSortCol(), param.GetSortDir())
	for _, u := range users {
		user, _ := fillUser(u)
		result = append(result, user)
	}
	return result, count
}

// Find returns the Task with the given id in the TaskManager and a boolean
// indicating if the id was found.
func (m *UserManager) Find(userID int) (*User, bool) {
	for _, t := range m.users {
		if t.ID == userID {
			return t, true
		}
	}
	return nil, false
}
func (m *UserManager) FindByLoginName(userLoginName string) (*User, bool) {
	//utils.Logger.Debug(fmt.Sprintf("%+v", m.users))
	for _, t := range m.users {
		//	utils.Logger.Debug(fmt.Sprintf("User:%+v", t))
		if nil != t && t.LoginName == userLoginName {
			return t, true
		}
	}
	return nil, false
}
func (m *UserManager) GetUsersByRole(roleID int) ([]*User, bool) {

	var result []*User
	//v.iDisplayStart = 1
	users, b := databases.DefaultContent.GetUsersByRole(roleID)
	if !b {

		utils.Logger.Debug(fmt.Sprintf("GetUsersByRole return %t", b))
		return nil, b

	}
	for _, item := range users {
		o, err := fillUser(item)
		if nil != err {
			utils.Logger.Debug(fmt.Sprintf("%v", err))
			return nil, false
		}
		result = append(result, o)
	}

	return result, true

}

func (m *UserManager) GenModel(tableName string, param *jquery.DataTablesParam) ([]*databases.ModelDef, int64) {
	//var v *jquery.DataTablesParam
	//v.iDisplayStart = 1
	fmt.Printf("%+v", param)
	models, _ := databases.DefaultContent.GenModel(tableName)
	return models, int64(len(models))
}
