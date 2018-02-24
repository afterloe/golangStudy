package routers

import (
	"fmt"
	"net/http"
)

type responseDTO struct {
	data interface{}
	code int
	msg *string
}

func (res *responseDTO) String() string {
	return fmt.Sprintf("Response {data: %s, code: %d, msg: %s}", res.data, res.code, res.msg)
}

func Success(data interface{}) *responseDTO {
	return &responseDTO{data, http.StatusOK, nil}
}

func Fail(code int, msg string) *responseDTO {
	return &responseDTO{nil, code, &msg}
}

func Build(data interface{}, code int, msg string) *responseDTO {
	return &responseDTO{data, code, &msg}
}