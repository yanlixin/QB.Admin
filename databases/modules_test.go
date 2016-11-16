package databases

import (
	"testing"

	"QB.Admin/utils"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

func TestGetAllModules(t *testing.T) {
	db := Instance()
	excepted := int(6)
	actual, b := db.GetAllModules()
	if !b {
		t.Errorf(utils.TestMessageBool, b)
	}
	if excepted >= len(actual) {
		t.Errorf(utils.TestMessageLength, excepted, len(actual))
	}
}
