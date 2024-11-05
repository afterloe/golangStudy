package internal

// PageParam 分页查询
type PageParam struct {
	Page int `default:"1" json:"page"`  // 第几页
	Size int `default:"10" json:"size"` // 每页多少条
}

// CommonResponse 通用返回格式
type CommonResponse[T any] struct {
	BizCode int    `json:"biz_code"` // 业务编码，0 - 正常
	Error   string `json:"error"`    // 异常信息，正确请求时为空白
	Result  T      `json:"result"`   // 返回的数据
}
