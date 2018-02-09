package DBUtil

func QueryInfo() (map[string]string, error){
	db, error := getConnection()
	if nil != error {
		return nil, &Error{error.Error()}
	}
	db.Close()

	return nil, nil
}
