package route

import (
	"hoge-api/controller"

	"github.com/gin-gonic/gin"
)

func GetRouter() (*gin.Engine, error) {
	engine := gin.Default() // エンジンを作成

	// endpoints
	// root page
	engine.GET("/", controller.ShowRootPage)
	// json test

	// endpoints group
	// ver1グループ
	v1 := engine.Group("/v1")
	{
		// usersグループ
		users := v1.Group("/users")
		{
			users.POST("/user", controller.RegisterUserHandler)
		}
	}

	return engine, nil // router設定されたengineを返す。
}
