package fs

import (
	"encoding/json"
)

func FormatToStruct(chunk string) (map[string]interface{}, error){
	rep := make(map[string]interface{})
	err := json.Unmarshal([]byte(chunk), &rep) // 使用这种方式来转换复杂json
	// slcB, _ := json.Marshal(slcD) // 这种方式可以直接转换 boolean、int、float、string、slice、map
	if nil != err {
		return nil, &Error{"json format error", 500}
	}
	return rep, nil
}

func FormatToString(vol interface{}) (string, error){
	buf, err := json.Marshal(vol)
	if nil != err {
		return "", &Error{"format object error", 500}
	}
	return string(buf),nil
}