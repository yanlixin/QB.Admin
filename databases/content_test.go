package databases

import (
	"testing"

	"QB.Admin/utils"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

func TestGenModel(t *testing.T) {
	db := Instance()
	_, actual := db.GenModel("pa_users")
	if !actual {
		t.Errorf(utils.TestMessageBool, true, actual)
	}
}
