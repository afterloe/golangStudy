package exceptions

import "fmt"

// 自定义 Error
type Error struct {
	Msg string
	Code int
}

// 自定义异常要实现这个方法
func (e *Error) Error() string {
	return fmt.Sprintf("%d - %s", e.Code, e.Msg)
}