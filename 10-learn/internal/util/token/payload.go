package token

import (
	"time"
)

// Payload token中存储的信息
type Payload struct {
	ID        string    `json:"id"`         // 用户id
	Username  string    `json:"username"`   // 用户名
	Role      []string  `json:"role"`       // 用户角色列表
	IssuedAt  time.Time `json:"issued_at"`  // token创建时间
	ExpiredAt time.Time `json:"expired_at"` // token过期时间
}
