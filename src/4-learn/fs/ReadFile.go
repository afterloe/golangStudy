package fs

import (
	"path/filepath"
	"io/ioutil"
)

const (
	BUFFER_SIZE = 64 * 1024
)

func ReadRealFile(path,name string) (string, error) {
	// 读取文件
	data, err := ioutil.ReadFile(filepath.Join(path, name))
	if nil != err {
		return "", &Error{"no such this file", 500}
	}

	// 将 byte 转换为 string
	return string(data), nil
}