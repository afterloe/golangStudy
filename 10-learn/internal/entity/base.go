package entity

import "time"

// _base
// 基础数据类
type _base struct {
	Id         string    `json:"_id"`
	CreatedAt  time.Time `json:"created_at"`  // 创建时间
	CreatedBy  string    `json:"created_by"`  // 创建人
	ModifiedAt time.Time `json:"modified_at"` // 修改时间
	ModifiedBy string    `json:"modified_by"` // 修改人
	Deleted    bool      `json:"deleted"`     // 是否逻辑删除
}
