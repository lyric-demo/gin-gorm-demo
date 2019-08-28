package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// NewUserCtl 创建用户管理控制器实例
func NewUserCtl(mUser IUserModel) *UserCtl {
	return &UserCtl{
		UserModel: mUser,
	}
}

// UserCtl 用户控制器
type UserCtl struct {
	UserModel IUserModel
}

// Query 查询用户数据
func (a *UserCtl) Query(c *gin.Context) {
	params := UserQueryParam{
		UserName: c.Query("user_name"),
	}

	users, err := a.UserModel.Query(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"list": users})
}

// Get 查询指定用户数据
func (a *UserCtl) Get(c *gin.Context) {
	recordID := c.Param("id")

	user, err := a.UserModel.Get(recordID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Create 创建用户
func (a *UserCtl) Create(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if user.UserName == "" || user.RealName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数"})
		return
	}

	user.RecordID = uuid.Must(uuid.NewRandom()).String()
	err := a.UserModel.Create(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

// Update 更新用户
func (a *UserCtl) Update(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if user.UserName == "" || user.RealName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数"})
		return
	}

	recordID := c.Param("id")
	err := a.UserModel.Update(recordID, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

// Delete 删除用户
func (a *UserCtl) Delete(c *gin.Context) {
	recordID := c.Param("id")
	err := a.UserModel.Delete(recordID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
