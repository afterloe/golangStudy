package errors

type BizError struct {
	ErrorCode int    `json:"error_code"`
	Msg       string `json:"msg"`
}

func (e *BizError) Error() string {
	return e.Msg
}
