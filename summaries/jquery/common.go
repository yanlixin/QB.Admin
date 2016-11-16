package jquery

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type DataTablesParam struct {
	DisplayStart       int      `form:"iDisplayStart"`
	DisplayLength      int      `form:"iDisplayLength"`
	Columns            int      `form:"iColumns"`
	Search             string   `form:"sSearch"`
	EscapeRegex        bool     `form:"bEscapeRegex"`
	Echo               int      `form:"sEcho"`
	Sortable           []bool   `form:"bSortable"`
	Searchable         []bool   `form:"bSearchable"`
	SearchColumns      []string `form:"sSearchColumns"`
	SortCol            []string `form:"iSortCol"`
	SortDir            []string `form:"sSortDir"`
	EscapeRegexColumns []bool   `form:"bEscapeRegexColumns"`
	DataProp           []string `form:"mDataProp"`
	ExtPara            string   `form:"extPara"`
}

func (s *DataTablesParam) Fill(m map[string]string) error {
	start, _ := strconv.Atoi(m["iDisplayStart"])
	//	length, _ := strconv.Atoi(m["iDisplayLength"])
	s.DisplayStart = start
	//s.iDisplayLength = length
	return nil
}
func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("Provided value type didn't match obj field type")
	}

	structFieldValue.Set(val)
	return nil
}

func (s *DataTablesParam) FillStruct(m map[string]interface{}) error {
	for k, v := range m {
		err := SetField(s, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
func (s DataTablesParam) GetSortCol() string {
	if len(s.SortCol) > 0 {
		return s.SortCol[0]
	} else {
		return "id"
	}
}
func (s DataTablesParam) GetSortDir() string {
	if len(s.SortDir) > 0 {
		return s.SortDir[0]
	} else {
		return "asc"
	}

}
