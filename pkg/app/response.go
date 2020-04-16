package app

import (
	"wardrobe_server/pkg/msg"

	"github.com/gin-gonic/gin"
)

// Gin .
type Gin struct {
	C *gin.Context
}
// Response .
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": httpCode,
		"msg":  msg.GetMsg(errCode),
		"data": data,
	})

	return
}
