package util

import (
	"io/ioutil"
	"path/filepath"
	"../../exceptions"
	"encoding/json"
)

func ReadRealFile(path,name string) (string, error) {
	// 读取文件
	data, err := ioutil.ReadFile(filepath.Join(path, name))
	if nil != err {
		return "", &exceptions.Error{Msg: "no such this file", Code: 500}
	}

	// 将 byte 转换为 string
	return string(data), nil
}

func FormatToStruct(chunk *string) (map[string]interface{}, error){
	rep := make(map[string]interface{})
	err := json.Unmarshal([]byte(*chunk), &rep)
	if nil != err {
		return nil, &exceptions.Error{Msg: "json format error", Code: 500}
	}
	return rep, nil
}

func FormatToString(vol interface{}) (string, error){
	buf, err := json.Marshal(vol)
	if nil != err {
		return "", &exceptions.Error{Msg: "format object error", Code: 500}
	}
	return string(buf), nil
}