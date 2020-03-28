package router

import (
	"wardrobe_server/middleware/jwt"
	"wardrobe_server/pkg/setting"
	api "wardrobe_server/routers/api"
	v1 "wardrobe_server/routers/api/v1"

	"github.com/gin-gonic/gin"
)

// InitRouter .
func InitRouter() *gin.Engine {

	gin.ForceConsoleColor() // 强制日志颜色化

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)
	// 数据库查看信息 如果有对应账号密码 就生成token
	r.GET("/auth", api.Auth)

	r.POST("/users", api.AddUser)

	apiV1 := r.Group("/api/v1")
	apiV1.Use(jwt.JWT())
	{
		apiV1.GET("/users", api.QureyUser)

		apiV1.POST("/pics", v1.AddPic)

		// apiV1.GET("/userInfo", api.GetUserInfo)

	}

	return r
}
