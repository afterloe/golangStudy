package DBUtil

import (
	"database/sql"
)

type queryExecute struct {
	sql string
	args []interface{}
}

func (query *queryExecute) rowsToMap(rows *sql.Rows) ([]map[string]interface{}, error) {
	cols, _ := rows.Columns()
	result := make([]map[string]interface{}, 0)
	columns := make([]interface{}, len(cols))
	columnPointers := make([]interface{}, len(cols))
	for i := range columns {
		columnPointers[i] = &columns[i]
	}
	for rows.Next() {
		if err := rows.Scan(columnPointers...); err != nil {
			return nil, &Error{err.Error()}
		}

		m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}
		result = append(result, m)
	}

	return result, nil
}

func (query *queryExecute) execute(db *sql.DB) (interface{}, error) {
	rows, err := db.Query(query.sql, query.args...)
	if nil != err {
		return nil, &Error{err.Error()}
	}
	return query.rowsToMap(rows)
}

// 传入查询的sql 返回 结果集组成的map结果
func QueryInfo(sql string, args ...interface{})([]map[string]interface{}, error){
	var point breakthroughPoint
	point = &queryExecute{sql, args}
	result, err := UseConnection(point)
	if nil != err {
		return nil, &Error{err.Error()}
	}
	return result.([]map[string]interface{}), nil
}
