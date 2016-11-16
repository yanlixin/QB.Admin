package databases

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"QB.Admin/utils"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq" // import your used driver
	//_ "github.com/go-sql-driver/mysql" // import your used driver
)

func init() {
	//	orm.RegisterDriver("mysql", orm.DRMySQL)
	//	orm.RegisterDataBase("default", "mysql", "root:ibcc.c0m@/qubei?charset=utf8")
	DefaultContent = &DbContent{"qubei"}
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=pgsql.user password=Aa123456 dbname=qubei-admin host=127.0.0.1 port=5432 sslmode=disable")
	DefaultContent = &DbContent{"test"}
	orm.Debug = true

}

type DbContent struct {
	DataBaseName string
}

var DefaultContent *DbContent

func Instance() *DbContent {
	result := &DbContent{}
	result.DataBaseName = "test"
	return result
}
func currentUserId() int64 {
	return 1
}
func getPKFile(val reflect.Value) string {
	ind := reflect.Indirect(val)
	for i := 0; i < ind.NumField(); i++ {
		//field := ind.Field(i)
		sf := ind.Type().Field(i)
		if sf.Tag != "" && strings.Index(string(sf.Tag), "pk") > 0 {

			return sf.Name
		}

	}
	return ""
}
func prpareSave(o interface{}) {

	ov := reflect.ValueOf(o).Elem()
	now := time.Now()
	var userId = currentUserId()
	ov.FieldByName("LastUpdatedByUserID").SetInt(userId)
	ov.FieldByName("LastUpdated").Set(reflect.ValueOf(now))
	pkFile := getPKFile(ov)
	if pkFile != "" && ov.FieldByName(pkFile).Int() <= 0 {
		ov.FieldByName("CreatedByUserID").SetInt(userId)
		ov.FieldByName("CreatedDate").Set(reflect.ValueOf(now))

	}

}
func (h *DbContent) SayHi(userName string) string {
	return "Hi" + userName
}

func (u *DbContent) GenModel(tableName string) ([]*ModelDef, bool) {

	status := true
	o := orm.NewOrm()
	var models []*ModelDef
	_, err := o.Raw("show columns from " + tableName).QueryRows(&models)
	if nil != err {
		status = false
		utils.Logger.Error(fmt.Sprintf("%v", err))
	}
	return models, status
}
