package picservice

import "wardrobe_server/models"

// Pic picservice.
type Pic struct {
	Md5         string
	URL          string
	AbsolutePath string
	Brand        string // 品牌
	Color      string //颜色
	Lable        string // 备注
	Type        string // 上衣之类
	Season       int    // 季节  0 默认
	Count        int    // 调用次数 每次使用就 +1
	Size         int    // 图片大小
}

// CheckImageSize 获取文件大小
// func CheckImageSize(f multipare.File)  {

// }

// CheckImageMD5 检查文件 md5 值
func CheckImageMD5(md5Value string) bool {
	return models.CheckImageMD5(md5Value)
}

// AddPic .
func (pic *Pic) AddPic() (bool,error)  {
	data:=map[string]interface{}{

	}
	return models.AddPic(pic)
}
