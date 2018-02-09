package DBUtil

func QueryInfo(sql string, args ...interface{})([]map[string]interface{}, error){
	db, error := getConnection()
	if nil != error {
		return nil, &Error{error.Error()}
	}
	rows, err := db.Query(sql, args...)
	if nil != err {
		return nil, &Error{err.Error()}
	}
	cols, _ := rows.Columns()
	result := make([]map[string]interface{}, 0)
	for rows.Next() {
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}

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

	db.Close()
	return result, nil
}
