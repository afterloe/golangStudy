package fs

import (
	"io/ioutil"
	"path/filepath"
)

const (
	PARENT = "/tmp"
	MODULE = 0755
)

func WriteRealFile(buf []byte, name string) (bool, error) {
	err := ioutil.WriteFile(filepath.Join(PARENT, name), buf, MODULE)
	if nil != err {
		return false, &Error{"can't write file.", 500}
	}

	return true, nil
}
