package picservice

import (
	"fmt"
	"wardrobe_server/models"
)

// Pic picservice.
type Pic struct {
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
func CheckImageMD5(md5Value string) bool {
	isExist := models.CheckImageMD5(md5Value)
	fmt.Printf("\n %s 是否存在 %t  ", md5Value, isExist)
	return isExist
}

// AddPic .
func (pic *Pic) AddPic() error {
	data := map[string]interface{}{
		"md5":         pic.Md5,
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

	// fmt.Printf("图片信息 %s %s ",pic.Md5,pic.URL)
	// return nil
}
