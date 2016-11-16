package databases

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type Predicate struct {
	iDisplayStart       int
	iDisplayLength      int
	iColumns            int
	sSearch             string
	bEscapeRegex        bool
	sEcho               int
	bSortable           []bool
	bSearchable         []bool
	sSearchColumns      []string
	iSortCol            []string
	sSortDir            []string
	bEscapeRegexColumns []bool
	mDataProp           []string
	extPara             string
	Params              []orm.ParamsList
}
