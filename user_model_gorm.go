package main

import "github.com/jinzhu/gorm"

// SchemaUser 定义用户对象别名
type SchemaUser User

// ToUserGormEntity 转换为用户实体
func (a SchemaUser) ToUserGormEntity() *UserGormEntity {
	return &UserGormEntity{
		RecordID: &a.RecordID,
		UserName: &a.UserName,
		RealName: &a.RealName,
		Age:      &a.Age,
		Gender:   &a.Gender,
		Memo:     &a.Memo,
	}
}

// UserGormEntity 定义用户实体
type UserGormEntity struct {
	gorm.Model
	RecordID *string `gorm:"column:record_id;size:36;"` // 记录ID
	UserName *string `gorm:"column:user_name;size:50;"` // 用户名
	RealName *string `gorm:"column:real_name;size:50;"` // 真实姓名
	Age      *int    `gorm:"column:age"`                // 年龄
	Gender   *int    `gorm:"column:gender"`             // 性别（1：男 2：女）
	Memo     *string `gorm:"column:memo;size:1024;"`    // 备注
}

// TableName 表名
func (a UserGormEntity) TableName() string {
	return "t_user"
}

// ToSchemaUser 转换为用户对象
func (a UserGormEntity) ToSchemaUser() *User {
	return &User{
		RecordID: *a.RecordID,
		UserName: *a.UserName,
		RealName: *a.RealName,
		Age:      *a.Age,
		Gender:   *a.Gender,
		Memo:     *a.Memo,
	}
}

// UserGormEntityList 定义用户实体列表别名
type UserGormEntityList []UserGormEntity

// ToSchemaUserList 转换为用户对象列表
func (a UserGormEntityList) ToSchemaUserList() []*User {
	items := make([]*User, len(a))

	for i, item := range a {
		items[i] = item.ToSchemaUser()
	}

	return items
}

// NewUserGormModel 创建用户管理gorm存储实例
func NewUserGormModel(db *gorm.DB) *UserGormModel {
	// 自动映射数据表结构到数据库
	db.AutoMigrate(new(UserGormEntity))

	return &UserGormModel{
		db: db,
	}
}

// UserGormModel 用户管理gorm存储
type UserGormModel struct {
	db *gorm.DB
}

// Query 查询用户数据
func (a *UserGormModel) Query(params UserQueryParam) ([]*User, error) {

	db := a.db.Model(UserGormEntity{})
	if v := params.UserName; v != "" {
		db = db.Where("user_name=?", v)
	}

	var items UserGormEntityList
	result := db.Find(&items)
	if err := result.Error; err != nil {
		return nil, err
	}

	return items.ToSchemaUserList(), nil
}

// Get 查询指定用户数据
func (a *UserGormModel) Get(recordID string) (*User, error) {
	var item UserGormEntity
	result := a.db.Model(UserGormEntity{}).Where("record_id=?", recordID).First(&item)
	if err := result.Error; err != nil {
		return nil, err
	}
	return item.ToSchemaUser(), nil
}

// Create 创建用户数据
func (a *UserGormModel) Create(item User) error {
	eitem := SchemaUser(item).ToUserGormEntity()

	result := a.db.Model(UserGormEntity{}).Create(eitem)
	return result.Error
}

// Update 更新用户数据
func (a *UserGormModel) Update(recordID string, item User) error {
	eitem := SchemaUser(item).ToUserGormEntity()

	result := a.db.Model(UserGormEntity{}).Where("record_id=?", recordID).Omit("record_id").Updates(eitem)
	return result.Error
}

// Delete 删除用户数据
func (a *UserGormModel) Delete(recordID string) error {
	result := a.db.Model(UserGormEntity{}).Where("record_id=?", recordID).Delete(UserGormEntity{})

	return result.Error
}
