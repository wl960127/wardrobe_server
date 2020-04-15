package picservice

import (
	"fmt"
	"wardrobe_server/models"
)

// Pic picservice.
type Pic struct {
	UserID       int
	Md5          string
	URL          string
	AbsolutePath string
	Brand        string // 品牌
	Color        string //颜色
	Lable        string // 备注
	Type         string // 上衣之类
	Season       int    // 季节  0 默认
	Count        int    // 调用次数 每次使用就 +1
	Size         int64  // 图片大小
}

// CheckImageSize 获取文件大小
// func CheckImageSize(f multipare.File)  {

// }

// CheckImageMD5 检查文件 md5 值
func CheckImageMD5(md5Value string) (isExit bool,url, absolutePath string) {
	isExist,url,absolutePath := models.CheckImageMD5(md5Value)
	fmt.Printf("\n %s 是否存在 %t  ", md5Value, isExist)
	return isExist,url,absolutePath
}

// AddPic .
func (pic *Pic) AddPic() error {

	data := map[string]interface{}{
		"userid":       pic.UserID,
		"md5":          pic.Md5,
		"url":          pic.URL,
		"absolutePath": pic.AbsolutePath,
		"brand":        pic.Brand,
		"color":        pic.Color,
		"lable":        pic.Lable,
		"type":         pic.Type,
		"season":       pic.Season,
		"size":         pic.Size,
	}
	return models.AddPic(data)
}

// QueryPic 数据库查看对应的所有图片
func (pic *Pic) QueryPic() map[string]interface{} {
	var picList []models.Picture
	var err error
	if picList, err = models.QueryPic(pic.UserID); err != nil {
		return map[string]interface{}{
			"msg": nil,
			"err": err.Error(),
		}
	}
	return map[string]interface{}{
		"mag": picList,
		"err": nil,
	}
}
