package v1

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"wardrobe_server/pkg/app"
	"wardrobe_server/pkg/e"
	"wardrobe_server/pkg/logging"
	"wardrobe_server/pkg/utils"
	picservice "wardrobe_server/service/pic_service"

	// "github.com/astaxie/beego/validation"
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

	

	userID := c.GetInt("claimsID")

	// 获取文件
	picFile, err := c.FormFile("file")
	if err != nil {
		// log.Println("文件上传失败")
		logging.Info("文件上传失败 " + err.Error())
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	fmt.Println(picFile.Filename)

	// 获取参数

	if err := c.ShouldBind(&picInfo); err != nil {
		logging.Info("品牌 %s 颜色 %s 备注 %s 类别 %s 季节%d \n %s", picInfo.BRAND, picInfo.COLOR, picInfo.LABLE, picInfo.TYPE, picInfo.SEASON, err.Error)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	fmt.Printf("品牌 %s 颜色 %s 备注 %s 类别 %s 季节%d", picInfo.BRAND, picInfo.COLOR, picInfo.LABLE, picInfo.TYPE, picInfo.SEASON)

	// valid := validation.Validation{}
	// valid.Max(picInfo.BRAND, 100, "brand").Message("最长为100字符")
	// valid.Max(picInfo.COLOR, 100, "color").Message("最长为100字符")
	// valid.Max(picInfo.LABLE, 150, "lable").Message("最长为150字符")
	// valid.Max(picInfo.TYPE, 100, "type").Message("最长为100字符")
	// valid.Max(picInfo.SEASON, 30, "season").Message("最长为30字符")

	// if valid.HasErrors() {
	// 	app.MarkErrors(valid.Errors)
	// 	appG.Response(http.StatusInternalServerError, e.ERROR_ADD_FAIL, valid.Errors)
	// 	return
	// }

	// 创建文件夹存储图片
	faceDir := "upload/pic"
	datePath := time.Now().Format("200601")
	folderPath := "./" + faceDir + "/" + datePath

	utils.IsNotExistMkDir(folderPath)
	fmt.Printf("创建 文件夹%s  ", folderPath)

	// make(*models.Picture{},0,2)

	if err := utils.IsNotExistMkDir(folderPath); err != nil {
		fmt.Printf("/n 创建文件夹失败 不继续了")
		appG.Response(http.StatusBadRequest, e.ERROR_ADD_FAIL, nil)
		return
	}
	// createDirIfNotExists(folderPath)

	src, err := picFile.Open()
	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_ADD_FAIL, nil)
		return
	}

	// log.Printf("上传的文件名字 %s", picFile.Filename)

	buf, err := ioutil.ReadAll(src)
	if err != nil {
		logging.Info(" 读取文件错误 %s\n", err.Error)
		appG.Response(http.StatusBadRequest, e.ERROR_ADD_FAIL, nil)
		return
	}
	hash := md5.New()
	hash.Write(buf)
	// 计算文件 md5值
	md5Hex := fmt.Sprintf("%x", hash.Sum(nil))

	// 有的话就不处理了吧
	if isExist := picservice.CheckImageMD5(md5Hex); isExist == true {
		appG.Response(http.StatusOK, e.SUCCESS, nil)
		return
	}
	// 获取后缀
	extName := utils.GetExt(picFile.Filename)
	if extName == "" {
		extName = ".jpg"
	}

	fileFullPath := filepath.Join(folderPath, md5Hex+extName)

	// f, err := utils.OpenFile(fileFullPath, os.O_RDWR|os.O_CREATE, 0644)
	f, err := os.OpenFile(fileFullPath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("\n打开文件失败 %s  ", err.Error())
		appG.Response(http.StatusBadRequest, e.ERROR_ADD_FAIL, nil)
		return
	}

	absPath, _ := filepath.Abs(f.Name())

	if _, err := f.Write(buf); err != nil {
		fmt.Printf("写入文件 %s  ", err.Error())
		appG.Response(http.StatusBadRequest, e.ERROR_ADD_FAIL, nil)
		return
	}
	f.Close()
	// log.Printf("Write Image File Success, path=%s", fileFullPath)
	// fileName := fmt.Sprintf("%s%s", md5Hex, extName)

	picService := picservice.Pic{
		UserID: userID,
		Md5:          md5Hex,
		URL:          fileFullPath,
		AbsolutePath: absPath,
		Brand:        picInfo.BRAND,
		Color:        picInfo.COLOR,
		Lable:        picInfo.LABLE,
		Type:         picInfo.TYPE,
		Size:         picFile.Size,
	}

	if err := picService.AddPic(); err == nil {
		appG.Response(http.StatusOK, e.SUCCESS, nil)
		return
	}

	appG.Response(http.StatusBadRequest, e.ERROR_ADD_FAIL, nil)
	return

}

// QueryPic 根据类型加载图片
func QueryPic(c *gin.Context) {
	appG := app.Gin{C: c}

	userID := c.GetInt("claimsID")

	picService := picservice.Pic{UserID: userID}

	data := picService.QueryPic()
	appG.Response(http.StatusOK, e.SUCCESS, data)

}

//创建目录, 如果目录不存在的话
func createDirIfNotExists(dir string) {
	// 判断目录是否存在
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		os.Mkdir(dir, 0777) //0777也可以os.ModePerm
		os.Chmod(dir, 0777)
	}
}
