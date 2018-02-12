package DBUtil

import "fmt"

// 自定义异常输出
type Error struct {
	msg string
}

func (e *Error) Error() string {
	return fmt.Sprintf("has error -> %s", e.msg)
}