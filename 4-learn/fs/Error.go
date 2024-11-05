package fs

import "fmt"

// 自定义 Error
type Error struct {
	msg string
	code int
}

// 自定义异常要实现这个方法
func (e *Error) Error() string {
	return fmt.Sprintf("%d - %s", e.code, e.msg)
}
