package main

// IUserModel 用户管理存储接口
type IUserModel interface {
	// 查询用户数据
	Query(params UserQueryParam) ([]*User, error)
	// 查询指定用户数据
	Get(recordID string) (*User, error)
	// 创建用户数据
	Create(item User) error
	// 更新用户数据
	Update(recordID string, item User) error
	// 删除用户数据
	Delete(recordID string) error
}
