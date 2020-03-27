package v1

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"wardrobe_server/pkg/app"
	"wardrobe_server/pkg/e"
	"wardrobe_server/pkg/logging"
	"wardrobe_server/pkg/utils"
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

	// 获取文件
	picFile, err := c.FormFile("upload")
	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	// 获取参数

	if err := c.ShouldBind(&picInfo); err != nil {
		log.Printf("品牌 %s 颜色 %s 备注 %s 类别 %s 季节%d", picInfo.BRAND, picInfo.COLOR, picInfo.LABLE, picInfo.TYPE, picInfo.SEASON)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	valid := validation.Validation{}
	valid.Max(picInfo.BRAND, 10, "brand").Message("最长为10字符")
	valid.Max(picInfo.COLOR, 10, "color").Message("最长为10字符")
	valid.Max(picInfo.LABLE, 10, "lable").Message("最长为150字符")
	valid.Max(picInfo.TYPE, 10, "type").Message("最长为10字符")
	valid.Max(picInfo.SEASON, 3, "season").Message("最长为3字符")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_FAIL, valid.Errors)
		return
	}


	// 创建文件夹存储图片
	faceDir := "upload/pic"
	datePath := time.Now().Format("200601")
	folderPath := "./" + faceDir + "/" + datePath

	utils.IsNotExistMkDir(folderPath)

	log.Printf("上传的文件名字 %s", picFile.Filename)

	// src, err := picFile.Open()
	// if err != nil {
	// 	logging.Error(err)
	// 	appG.Response(http.StatusBadRequest, e.ERROR_ADD_FAIL, nil)
	// 	return
	// }
	// defer src.Close()

	if buf, err := ioutil.ReadFile(picFile.Filename); err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_ADD_FAIL, nil)
		return
	} else {
		hash := md5.New()
		hash.Write(buf)
		// 计算文件 md5值
		md5Hex := fmt.Sprintf("%x", hash.Sum(nil))


		// 有的话就不处理了吧
		if isExist := picservice.CheckImageMD5(md5Hex); isExist {
			appG.Response(http.StatusOK, e.SUCCESS, nil)
			return
		}
		// 获取后缀
		extName := utils.GetExt(picFile.Filename)
		if extName == "" {
			extName = ".jpg"
		}

		fileFullPath := filepath.Join(folderPath, md5Hex+extName)

		if f, err := utils.Open(fileFullPath, os.O_RDWR|os.O_CREATE, 0644); err != nil {
			appG.Response(http.StatusBadRequest, e.ERROR_ADD_FAIL, nil)
			return
		} else {

			absPath, _ := filepath.Abs(f.Name())

			if _,err:= f.Write(buf);err != nil {
				appG.Response(http.StatusBadRequest, e.ERROR_ADD_FAIL, nil)
				return
			}
			f.Close()
			log.Printf("Write Image File Success, path=%s", filefullPath)
			fileName := fmt.Sprintf("%s%s", md5Hex, extName)

		
			picService := picservice.Pic{
				Md5: md5Hex,
				URL: fileFullPath,
				AbsolutePath: absPath,
				Brand: picInfo.BRAND,
				Color: picInfo.COLOR,
				Lable: picInfo.LABLE,
				Type: picInfo.TYPE,
				Size: utils.GetSize(picFile),
			}

			if isSuccess ,err :=picService.AddPic(); err == nil {
			appG.Response(http.StatusOK, e.SUCCESS, nil)
			return
			}	
			
			appG.Response(http.StatusBadRequest, e.ERROR_ADD_FAIL, nil)
			return
	}


	// picService :=picservice.Pic{

	// }
}

// func UpdatePic(c *gin.Context)  {

// }

// func DelPic(c *gin.Context)  {

// }
