package DBUtil

import "fmt"

type Error struct {
	msg string
}

func (e *Error) Error() string {
	return fmt.Sprintf("has error -> %s", e.msg)
}