package DBUtil

import (
	"database/sql"
	"reflect"
)

type insertExecute struct {
	sql string
	domainList []interface{}
}

func mapToValueList(m interface{}) []interface{} {
	val := reflect.ValueOf(m)
	valueList := make([]interface{}, 0)
	for i := 0; i < val.Len(); i++ {
		valueList = append(valueList, val.Index(i).Interface())
	}
	return valueList
}

func (insert *insertExecute) execute(db *sql.DB) (interface{}, error) {
	stmt, error := db.Prepare(insert.sql)
	if nil != error {
		return nil, error
	}
	insertHistory := make([]interface{}, 0)
	for _, domain := range insert.domainList {
		if res, err := stmt.Exec(mapToValueList(domain)...); nil != err {
			return false, err
		} else {
			insertHistory = append(insertHistory, res)
		}
	}

	return insertHistory, nil
}

func InsertMap(sql string, domainList []interface{}) ([]interface{}, error) {
	var point breakthroughPoint
	point = &insertExecute{sql, domainList}
	result, err := UseConnection(point)
	if nil != err {
		return nil, &Error{err.Error()}
	}

	return result.([]interface{}), nil
}
