package token

import (
	"errors"
	"time"
)

const (
	ACCESS  = "access"
	REFRESH = "refresh"
)

// ITokenHelper tokenHelper接口
type iTokenHelper interface {
	// CreateToken
	// 为登陆用户创建token
	// @Param id string 用户id
	// @Param username string 用户名
	// @Param role []string 用户角色列表
	// @Param tokenType string token类型 access, refresh
	// @Return token， token.Payload， 异常
	CreateToken(id, username string, role []string, tokenType string) (string, *Payload, error)

	// VerifyToken
	// 验证Token字符串是否有效
	// @Param token string token字符串
	// @Return token.Payload，异常
	VerifyToken(token string) (*Payload, error)

	// SetDurationTime
	// 设置token存在时间
	// @Param accessTokenDuration 访问token存在时间
	// @Param refreshTokenDuration 刷新token存在时间
	SetDurationTime(accessTokenDuration, refreshTokenDuration time.Duration)
}

// ValidExpired
// 验证token是否超时
// Param expired 设置的时间
// Return error 未超时返回nil
func ValidExpired(expired time.Time) error {
	if time.Now().After(expired) {
		return errors.New("token 已过期")
	}
	return nil
}

// generatorPayload
// 生成token封装信息
// Param id string 用户id
// Param username string 用户名
// Param role []string 用户角色列表
// Param duration time.Duration token存在时间
// Return token.Payload
func generatorPayload(id, username string, role []string, duration time.Duration) *Payload {
	return &Payload{
		ID:        id,
		Username:  username,
		Role:      role,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
}
