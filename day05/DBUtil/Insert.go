package DBUtil

import (
	"database/sql"
	"reflect"
	"fmt"
	"strconv"
)

type insertExecute struct {
	sql string
	domainList []map[string]interface{}
}

func mapToValueList(m map[string]interface{}) []interface{} {
	val := reflect.ValueOf(m)
	valueList := make([]interface{}, 0)
	for _, key := range val.MapKeys() {
		valueList = append(valueList, val.MapIndex(key).Interface())
	}
	return valueList
}

func (insert *insertExecute) execute(db *sql.DB) (interface{}, error) {
	stmt, error := db.Prepare(insert.sql)
	if nil != error {
		return nil, error
	}
	resultList := make([]interface{}, 0)
	for i, domain := range insert.domainList {
		if res, err := stmt.Exec(mapToValueList(domain)...); nil != err {
			resultList = append(resultList, res)
		} else {
			fmt.Println(err.Error())
			fmt.Println("insert error! => index is " + strconv.Itoa(i))
			return resultList, err
		}
	}

	return resultList, nil
}

func InsertMap(sql string, domainList []map[string]interface{}) (bool, error) {
	var point breakthroughPoint
	point = &insertExecute{sql, domainList}
	_, err := UseConnection(point)
	if nil != err {
		return false, &Error{err.Error()}
	}

	return true, nil
}
