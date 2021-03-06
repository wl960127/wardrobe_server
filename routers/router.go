package router

import (
	// "wardrobe_server/middleware/jwt"
	parseConfig "wardrobe_server/pkg/app/parseConfig"
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

	r.Static("/upload", "./upload")

	gin.SetMode(parseConfig.ServerSetting.RunMode)
	// 数据库查看信息 如果有对应账号密码 就生成token

	r.GET("/auth", api.Auth)

	r.POST("/users", api.AddUser)

	r.GET("/ws", api.WsPage)


	apiV1 := r.Group("/api/v1")
	// apiV1.Use(jwt.JWT())
	{
		apiV1.GET("/users", api.QureyUser)

		apiV1.POST("/pics", v1.AddPic)
		apiV1.GET("/pics", v1.QueryPic)




		// apiV1.GET("/userInfo", api.GetUserInfo)

	}

	return r
}
