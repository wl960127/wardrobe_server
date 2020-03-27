package models

import "github.com/jinzhu/gorm"

// Pic 图片 MD5 计算唯一值  查看是否存在 不存在就保存.
// type Pic struct {
// 	MD5          string
// 	URL          string
// 	AbsolutePath string
// 	BRAND        string // 品牌
// 	COLOR        string //颜色
// 	LABLE        string // 备注
// 	TYPE         string // 上衣之类
// 	SEASON       int    // 季节  0 默认
// 	Count        int    // 调用次数 每次使用就 +1
// 	Size         int    // 图片大小
// }

//AddPic

// CheckImageMD5 检查是否文件存在 md5 可以验证唯一.
func CheckImageMD5(md5Value string) bool {
	// var picFile Picture
	if err := db.Where("MD5 = ?", md5Value).First(&Picture{}).Error; err != gorm.ErrRecordNotFound {
		return true
	}
	return false
}

// AddPic 添加图片
func AddPic() {

}
