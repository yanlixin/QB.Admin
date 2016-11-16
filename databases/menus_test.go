package databases

import (
	"fmt"
	"testing"
	"time"

	"QB.Admin/utils"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

func TestInsertMenut(t *testing.T) {
	db := DefaultContent
	m := new(Menus)
	m.MenuName = fmt.Sprintf("%s%s", "TestMenu", time.Now())
	m.MenuDesc = "Unit Test Menu"
	m.MenuPID = 0
	m.ModuleID = 1
	m.CreatedByUserID = 1
	m.LastUpdatedByUserID = 1
	actual, b := db.InsertMenu(m)
	excepted := int64(0)
	if !b {
		t.Errorf(utils.TestMessageBool, true, b)
	}

	if actual <= excepted {
		t.Errorf(utils.TestMessageLength, excepted, actual)
	}
}
func TestUpdateMenu(t *testing.T) {
	db := DefaultContent
	menus, _ := db.GetAllMenus()
	m := menus[len(menus)-1]
	m.LastUpdatedByUserID += m.LastUpdatedByUserID

	actual, b := DefaultContent.UpdateMenu(m)
	if !b {
		t.Errorf(utils.TestMessageBool, true, b)
	}
	except := int64(1)
	if actual != except {
		t.Errorf(utils.TestMessageLength, actual, except)
	}
}
func TestDeleteMenu(t *testing.T) {
	db := DefaultContent
	menus, _ := DefaultContent.GetAllMenus()
	_, b := db.DeleteMenu(menus[len(menus)-1].MenuID)
	if !b {
		t.Errorf(utils.TestMessageBool, true, b)
	}
}

func TestGetAllMenus(t *testing.T) {

	except := 0
	actual, b := DefaultContent.GetAllMenus()
	if !b {
		t.Errorf(utils.TestMessageBool, b)
	}
	if except >= len(actual) {
		t.Errorf(utils.TestMessageLength, len(actual))
	}
}
func TestGetMenu(t *testing.T) {
	db := Instance()

	actual, b := db.GetMenu(1)
	if !b {
		t.Errorf(utils.TestMessageBool, b)
	}
	if actual == nil {
		t.Errorf(utils.TestMessageLength, nil, actual)
	}
}
