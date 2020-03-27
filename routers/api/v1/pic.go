package v1

import (
	"log"
	"net/http"
	"wardrobe_server/pkg/app"
	"wardrobe_server/pkg/e"
	picservice "wardrobe_server/service/pic_service"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

//
type pic struct {
	BRAND  string `form:"brand" `  // 品牌
	COLOR  string `form:"color" `  //颜色
	LABLE  string `form:"lable" `  // 备注
	TYPE   string `form:"type" `   // 上衣之类
	SEASON int    `form:"season" ` // 季节  0 默认
}

// AddPic 添加图片 .
func AddPic(c *gin.Context) {
	appG := app.Gin{C: c}
	var picInfo pic
	if file, err := c.FormFile("upload"); err != nil {

		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	} 

	// 先获取文件

	// if err := c.ShouldBind(&picInfo); err != nil {
	// 	log.Printf("品牌 %s 颜色 %s 备注 %s 类别 %s 季节%d", picInfo.BRAND, picInfo.COLOR, picInfo.LABLE, picInfo.TYPE, picInfo.SEASON)
	// 	appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	// 	return
	// }

	// valid := validation.Validation{}
	// valid.Max(picInfo.BRAND, 10, "brand").Message("最长为10字符")
	// valid.Max(picInfo.COLOR, 10, "color").Message("最长为10字符")
	// valid.Max(picInfo.LABLE, 10, "lable").Message("最长为150字符")
	// valid.Max(picInfo.TYPE, 10, "type").Message("最长为10字符")
	// valid.Max(picInfo.SEASON, 3, "season").Message("最长为3字符")

	// if valid.HasErrors() {
	// 	app.MarkErrors(valid.Errors)
	// 	appG.Response(http.StatusInternalServerError, e.ERROR_ADD_FAIL, valid.Errors)
	// 	return
	// }

	// picService :=picservice.Pic{

	// }
}

// func UpdatePic(c *gin.Context)  {

// }

// func DelPic(c *gin.Context)  {

// }
