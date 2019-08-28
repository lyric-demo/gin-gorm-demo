package main

// User 用户对象
type User struct {
	RecordID string `json:"record_id"` // 记录ID
	UserName string `json:"user_name"` // 用户名
	RealName string `json:"real_name"` // 真实姓名
	Age      int    `json:"age"`       // 年龄
	Gender   int    `json:"gender"`    // 性别（1：男 2：女）
	Memo     string `json:"memo"`      // 备注
}

// UserQueryParam 用户查询参数
type UserQueryParam struct {
	UserName string // 用户名
}
