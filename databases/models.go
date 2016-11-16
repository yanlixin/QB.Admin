package databases

import (
	"time"

	"github.com/astaxie/beego/orm"
)

func init() {
	// Need to register model in init
	orm.RegisterModelWithPrefix("pa_", new(Users))
	orm.RegisterModelWithPrefix("pa_", new(Roles))
	orm.RegisterModelWithPrefix("pa_", new(Menus))

	orm.RegisterModelWithPrefix("pa_", new(R_Users_Roles))
	orm.RegisterModelWithPrefix("pa_", new(Permissions))
	orm.RegisterModelWithPrefix("pa_", new(R_Roles_Permissions))
	orm.RegisterModelWithPrefix("pa_", new(Modules))
	//orm.RunSyncdb("default", false, true)
}

type ModelDef struct {
	FieldName    string `orm:"column(Field)"`
	FieldType    string `orm:"column(Type)"`
	FieldNull    bool   `orm:"column(Null)"`
	FieldKey     string `orm:"column(Key)"`
	FieldDefault string `orm:"column(Default)"`
}
type BaseModel struct {
}
type Users struct {
	UserID                    int       `orm:"pk;column(userid);auto"`
	UserLoginName             string    `orm:"column(userloginname);unique"`
	UserPassword              string    `orm:"column(userpassword)"`
	UserDisplayName           string    `orm:"column(userdisplayname)"`
	UserPhoto                 string    `orm:"column(userphoto)"`
	UserSignature             string    `orm:"column(usersignature)"`
	UserMustChangePwd         bool      `orm:"column(usermustchangepwd);null"`
	UserCannotChangePwd       bool      `orm:"column(usercannotchangepwd);null"`
	UserAccountExpirationDate time.Time `orm:"column(useraccountexpirationdate);null"`
	RecordStatus              int       `orm:"column(recordstatus)"`
	CreatedDate               time.Time `orm:"column(createddate)"`
	CreatedByUserID           int       `orm:"column(createdbyuserid)"`
	LastUpdated               time.Time `orm:"column(lastupdated)"`
	LastUpdatedByUserID       int       `orm:"column(lastupdatedbyuserid)"`
}

func (u *Users) TableName() string {
	return "users"
}

type Roles struct {
	RoleID              int       `orm:"pk;column(roleid);auto"`
	RoleName            string    `orm:"column(rolename)"`
	RoleDesc            string    `orm:"column(roledesc);null"`
	RecordStatus        int       `orm:"column(recordstatus)"`
	CreatedDate         time.Time `orm:"column(createddate)"`
	CreatedByUserID     int       `orm:"column(createdbyuserid)"`
	LastUpdated         time.Time `orm:"column(lastupdated)"`
	LastUpdatedByUserID int       `orm:"column(lastupdatedbyuserid)"`
}

func (t *Roles) TableName() string {
	return "roles"
}

type Permissions struct {
	PermissionID        int       `orm:"pk;column(permissionid);auto"`
	MenuID              int       `orm:"column(menuid)"`
	ActionID            string    `orm:"column(actionid);null"`
	PermissionName      string    `orm:"column(permissionname)"`
	PermissionDesc      string    `orm:"column(permissiondesc);null"`
	PermissionGroup     string    `orm:"column(permissiongroup);null"`
	PermissionMemo      string    `orm:"column(permissionmemo);null"`
	RecordStatus        int       `orm:"column(recordstatus)"`
	CreatedDate         time.Time `orm:"column(createddate)"`
	CreatedByUserID     int       `orm:"column(createdbyuserid)"`
	LastUpdated         time.Time `orm:"column(lastupdated)"`
	LastUpdatedByUserID int       `orm:"column(lastupdatedbyuserid)"`
}

func (t *Permissions) TableName() string {
	return "permissions"
}

type R_Users_Roles struct {
	R_User_RoleID int `orm:"pk;column(r_user_roleid);auto"`
	RoleID        int `orm:"column(roleid)"`
	UserID        int `orm:"column(userid)"`
}

func (t *R_Users_Roles) TableName() string {
	return "r_users_roles"
}

type Modules struct {
	ModuleID            int       `orm:"pk;column(moduleid);auto"`
	ModulePID           int       `orm:"column(modulepid)"`
	ModuleName          string    `orm:"column(modulename)"`
	ModuleDesc          string    `orm:"column(moduledesc);null"`
	SortIndex           int       `orm:"column(sortindex)"`
	IsTreeLeaf          bool      `orm:"column(istreeleaf)"`
	TreeLevel           int       `orm:"column(treelevel)"`
	RecordStatus        int       `orm:"column(recordstatus)"`
	CreatedDate         time.Time `orm:"column(createddate)"`
	CreatedByUserID     int       `orm:"column(createdbyuserid)"`
	LastUpdated         time.Time `orm:"column(lastupdated)"`
	LastUpdatedByUserID int       `orm:"column(lastupdatedbyuserid)"`
}

func (t *Modules) TableName() string {
	return "modules"
}

type Menus struct {
	MenuID              int       `orm:"pk;column(menuid);auto"`
	ModuleID            int       `orm:"column(moduleid)"`
	MenuPID             int       `orm:"column(menupid)"`
	MenuName            string    `orm:"column(menuname)"`
	MenuIconUrl         string    `orm:"column(menuiconurl);null"`
	MenuNavUrl          string    `orm:"column(menunavurl);null"`
	MenuTarget          string    `orm:"column(menutarget);null"`
	MenuDesc            string    `orm:"column(menudesc);null"`
	SortIndex           int       `orm:"column(sortindex)"`
	IsTreeLeaf          bool      `orm:"column(istreeleaf)"`
	ActionID            string    `orm:"column(actionid);null"`
	TreeLevel           int       `orm:"column(treelevel)"`
	RecordStatus        int       `orm:"column(recordstatus)"`
	CreatedDate         time.Time `orm:"column(createddate)"`
	CreatedByUserID     int       `orm:"column(createdbyuserid)"`
	LastUpdated         time.Time `orm:"column(lastupdated)"`
	LastUpdatedByUserID int       `orm:"column(lastupdatedbyuserid)"`
}

func (t *Menus) TableName() string {
	return "menus"
}

type R_Roles_Permissions struct {
	R_Role_PermissionID int `orm:"pk;column(r_role_permissionid);auto"`
	RoleID              int `orm:"column(roleid)"`
	PermissionID        int `orm:"column(permissionid)"`
}

func (t *R_Roles_Permissions) TableName() string {
	return "r_roles_permissions"
}
