package DBUtil

import (
	"database/sql"
	"reflect"
)

type updateExecute struct {
	sql string
	domainList []interface{}
}

func checkRes(updateHistory []interface{}) (interface{}, error) {
	for _, update := range updateHistory {
		val := reflect.ValueOf(update)
		val.MethodByName("RowsAffected").Call(nil)
	}

	return true, nil
}

func (update *updateExecute) execute(db *sql.DB) (interface{}, error) {
	stmt, error := db.Prepare(update.sql)
	if nil != error {
		return nil, &Error{error.Error()}
	}
	updateHistory := make([]interface{}, 0)
	for _, domain := range update.domainList {
		if res, err := stmt.Exec(mapToValueList(domain)...); nil != err {
			return nil, &Error{err.Error()}
		} else {
			updateHistory = append(updateHistory, res)
		}
	}

	return checkRes(updateHistory)
}

func UpdateMap(sql string, domainList []interface{}) ([]interface{}, error) {
	var point breakthroughPoint
	point = &updateExecute{sql, domainList}
	result, err := UseConnection(point)
	if nil != err {
		return nil, &Error{err.Error()}
	}

	return result.([]interface{}), nil
}
