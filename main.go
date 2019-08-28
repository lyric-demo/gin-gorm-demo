package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建gin引擎实例
	app := gin.New()

	// 使用日志中间件(全局注入)
	app.Use(gin.Logger())

	// 使用崩溃恢复中间件(全局注入)
	app.Use(gin.Recovery())

	// 创建用户存储
	userModel := NewUserMemoryModel()

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
