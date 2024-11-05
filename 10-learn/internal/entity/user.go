package entity

// User
// 用户
type User struct {
	ID         string `json:"id"`
	Username   string `json:"username"`
	Ciphertext string `json:"ciphertext"`
	_base
}
