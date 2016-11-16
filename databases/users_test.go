package databases

import (
	"fmt"
	"testing"
	"time"

	"QB.Admin/utils"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

func TestLogin(t *testing.T) {
	db := Instance()
	except := true
	actual, _ := db.LoginUser("master", "123")
	if except != actual {
		t.Errorf(utils.TestMessageBool, except, actual)
	}
}
func TestGetUser(t *testing.T) {
	db := Instance()
	_, b := db.GetUser(12)
	if !b {
		t.Errorf(utils.TestMessageBool, true, b)
	}
}
func TestGetAllUsers(t *testing.T) {
	db := Instance()
	except := int64(17)
	users, actual := db.GetAllUsers(15, 15)
	if except >= actual {
		t.Errorf(utils.TestMessageLength, except, actual)
	}
	utils.Logger.Debug(fmt.Sprintf("%+v", users[1]))
	if users[1] == nil {

		t.Errorf("Traget user is nil")
	}
}
func TestInsertUser(t *testing.T) {
	db := DefaultContent
	u := new(Users)
	u.UserLoginName = fmt.Sprintf("%s%s", "TestUser", time.Now())
	u.UserDisplayName = "Unit Test User"
	u.CreatedByUserID = 1
	u.LastUpdatedByUserID = 1
	actual, b := db.InsertUser(u)
	excepted := int64(0)
	if !b {
		t.Errorf(utils.TestMessageBool, true, b)
	}

	if actual <= excepted {
		t.Errorf(utils.TestMessageLength, excepted, actual)
	}
}
func TestUpdateUser(t *testing.T) {
	db := DefaultContent
	users, _ := db.GetAllUsers(20, 0)
	u := users[12]
	u.LastUpdatedByUserID += u.LastUpdatedByUserID
	actual, b := DefaultContent.UpdateUser(u)
	if !b {
		t.Errorf(utils.TestMessageBool, true, b)
	}
	except := int64(13)
	if actual != except {
		t.Errorf(utils.TestMessageLength, actual, except)
	}
}
func TestDeleteUser(t *testing.T) {
	db := DefaultContent
	_, b := db.DeleteUser(12)
	if !b {
		t.Errorf(utils.TestMessageBool, true, b)
	}
}
func TestGetRolesByUser(t *testing.T) {
	db := DefaultContent
	_, b := db.GetRolesByUser(1)
	if !b {
		t.Error(utils.TestMessageBool, true, b)
	}
}
func TestAddRolesToUser(t *testing.T) {
	db := DefaultContent
	roleIDs := []int{1, 2, 3}
	userID := 2
	b := db.AddRolesToUser(userID, roleIDs)
	if !b {
		t.Errorf(utils.TestMessageBool, true, b)
	}
}
func TestDeleteRolesFromUser(t *testing.T) {
	db := DefaultContent
	roleIDs := []int{1, 2, 3, 4}
	userID := 2
	b := db.DeleteRolesFromUser(userID, roleIDs)
	if !b {
		t.Errorf(utils.TestMessageBool, true, b)
	}
}
