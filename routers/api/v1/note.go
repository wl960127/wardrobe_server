package v1

import (
	"net/http"

	"wardrobe_server/pkg/app"
	"wardrobe_server/pkg/msg"
	noteservice "wardrobe_server/service/noteService"


	"github.com/gin-gonic/gin"
)

type note struct {
	Experience string `form:"experience"` // 心得 备注
	PicWhole   string `form:"whole"`      //整体
	PicCoat    string `form:"coat"`       //上衣
	PicSkirt   string `form:"skirt"`      //裙子
	PicPants   string `form:"pants"`      //裤子
	PicShoes   string `form:"shoes"`      //鞋子
}

// AddNote 添加note的api 接口
func AddNote(c *gin.Context)  {
	appG := app.Gin{C:c}

	var noteInfo note

	if err := c.ShouldBind(&noteInfo);err!=nil {
		appG.Response(http.StatusBadRequest, msg.INVALID_PARAMS, nil)
		return
	}


	userID := c.GetInt("claimsID")

	noteService := noteservice.Note{
		UserID: userID,
		Experience: noteInfo.Experience,
		PicWhole: noteInfo.PicWhole,
		PicCoat: noteInfo.PicCoat,
		PicSkirt: noteInfo.PicSkirt,
		PicPants: noteInfo.PicPants,
		PicShoes: noteInfo.PicShoes
	}


	if err :=noteservice.AddNote();err !=nil{
		appG.Response(http.StatusBadRequest, msg.ERROR_ADD_FAIL, err.Error())
		return
	}
	appG.Response(http.StatusOK, msg.SUCCESS, nil)
	return


}