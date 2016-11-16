package databases

import (
	"fmt"
	"testing"

	"QB.Admin/utils"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

func TestGetAllPermissions(t *testing.T) {
	db := Instance()
	excepted := int(5)
	actual, b := db.GetAllPermissions(2)
	if !b {
		t.Errorf(utils.TestMessageBool, true, b)
	}
	if excepted >= len(actual) {
		t.Errorf(utils.TestMessageLength, excepted, len(actual))
	}
}
func TestGetPermissionsByRole(t *testing.T) {
	db := Instance()
	excepted := int(6)
	actual, b := db.GetRolePermissions()
	if !b {
		t.Errorf(utils.TestMessageBool, true, b)
	}
	if excepted >= len(actual) {
		t.Errorf(utils.TestMessageLength, excepted, len(actual))
	}
}
func TestGetRolePermissions(t *testing.T) {
	db := Instance()
	excepted := int(6)
	actual, b := db.GetRolePermissions()
	fmt.Printf("%+v", actual)
	if !b {
		t.Errorf(utils.TestMessageBool, true, b)
	}
	if excepted >= len(actual) {
		t.Errorf(utils.TestMessageLength, excepted, len(actual))
	}
}
