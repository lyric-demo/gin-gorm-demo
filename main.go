package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	// 引入mysql驱动
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func newGormDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local&allowNativePasswords=true")
	if err != nil {
		panic(err)
	}
	return db
}

func main() {
	gdb := newGormDB()

	// 创建gin引擎实例
	app := gin.New()

	// 使用日志中间件(全局注入)
	app.Use(gin.Logger())

	// 使用崩溃恢复中间件(全局注入)
	app.Use(gin.Recovery())

	// 使用验证用户中间件(全局注入)
	app.Use(verifyUserMiddleware)

	// 创建用户存储
	userModel := NewUserGormModel(gdb)
	// userModel := NewUserMemoryModel()

	// 创建用户控制器
	userCtl := NewUserCtl(userModel)

	// 增加路由组/api
	api := app.Group("/api")
	{ // 约束作用域为api组
		// 增加嵌套路由组/v1
		v1 := api.Group("/v1")
		{ // 约束作用域为v1组

			// 增加GET请求的/users路由(完整请求路由为：GET /api/v1/users)
			v1.GET("/users", userCtl.Query)

			// 增加GET请求的/users/:id路由(完整请求路由为：GET /api/v1/users/id)
			v1.GET("/users/:id", userCtl.Get)

			// 增加POST请求的/users路由(完整请求路由为：POST /api/v1/users)
			v1.POST("/users", userCtl.Create)

			// 增加PUT请求的/users/:id路由(完整请求路由为：PUT /api/v1/users/id)
			v1.PUT("/users/:id", userCtl.Update)

			// 增加DELETE请求的/users/:id路由(完整请求路由为：DELETE /api/v1/users/id)
			v1.DELETE("/users/:id", userCtl.Delete)

		}
	}

	log.Printf("HTTP服务监听在8080端口...")
	err := app.Run(":8080")
	if err != nil {
		log.Fatalf("监听服务发生错误：%s", err.Error())
	}
}

// 验证用户中间件
// 中间件的作用是什么：
// 拦截所有请求数据，统一处理公共业务（比如：记录所有请求日志、验证请求用户），同时把处理的业务数据放入上下文，在后续业务中可以从上下文中获取并使用
// 如何实现一个中间件：
// 中间件本质上是一个回调函数，特殊之处在于上下文处理及调用Next函数执行下一个中间件，以此类推
func verifyUserMiddleware(c *gin.Context) {
	token := c.GetHeader("Access-Token")
	if token != "000000" {
		c.JSON(401, gin.H{"message": "无效的访问令牌", "code": 0})
		c.Abort()
		return
	}

	c.Set("user_id", "root")
	c.Next()
}
