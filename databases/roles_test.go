package databases

import (
	"fmt"
	"testing"
	"time"

	"QB.Admin/utils"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

func TestInsertRole(t *testing.T) {
	db := DefaultContent
	r := new(Roles)
	r.RoleName = fmt.Sprintf("%s%s", "TestRole", time.Now())
	r.RoleDesc = "Unit Test Role"
	r.CreatedByUserID = 1
	r.LastUpdatedByUserID = 1
	actual, b := db.InsertRole(r)
	excepted := int64(0)
	if !b {
		t.Errorf(utils.TestMessageBool, true, b)
	}

	if actual <= excepted {
		t.Errorf(utils.TestMessageLength, excepted, actual)
	}
}
func TestUpdateRole(t *testing.T) {
	db := DefaultContent
	roles, _ := db.GetAllRoles(0, 20)
	r := roles[len(roles)-1]
	r.LastUpdatedByUserID += r.LastUpdatedByUserID

	actual, b := DefaultContent.UpdateRole(r)
	if !b {
		t.Errorf(utils.TestMessageBool, true, b)
	}
	excepted := int64(1)
	if actual != excepted {
		t.Errorf(utils.TestMessageLength, actual, excepted)
	}
}
func TestDeleteRole(t *testing.T) {
	db := DefaultContent
	roles, _ := db.GetAllRoles(0, 20)
	_, b := db.DeleteRole(roles[len(roles)-1].RoleID)
	if !b {
		t.Errorf(utils.TestMessageBool, true, b)
	}
}
func TestGetAllRoles(t *testing.T) {
	db := Instance()
	excepted := int(6)
	actual, _ := db.GetAllRoles(0, 10)

	if excepted >= len(actual) {
		t.Errorf(utils.TestMessageLength, excepted, len(actual))
	}
}
func TestGetRole(t *testing.T) {
	db := Instance()

	actual, b := db.GetRole(1)
	if !b {
		t.Errorf(utils.TestMessageBool, b)
	}
	if actual == nil {
		t.Errorf(utils.TestMessageLength, nil, actual)
	}
}
func TestAddUsersToRole(t *testing.T) {
	db := DefaultContent
	userIDs := []int{2, 3, 4, 5}
	roleID := 2
	b := db.AddUsersToRole(roleID, userIDs)
	if !b {
		t.Errorf(utils.TestMessageBool, true, b)
	}
}
func TestDeleteUsersFromRole(t *testing.T) {
	db := DefaultContent
	userIDs := []int{2, 3, 4, 5}
	roleID := 2
	b := db.DeleteUsersFromRole(roleID, userIDs)
	if !b {
		t.Errorf(utils.TestMessageBool, true, b)
	}
}

func TestAddPermissionsToRole(t *testing.T) {
	db := DefaultContent
	perIDs := []int{2, 3, 4, 5}
	roleID := 2
	b := db.AddPermissionsToRole(roleID, perIDs)
	if !b {
		t.Errorf(utils.TestMessageBool, true, b)
	}
}
